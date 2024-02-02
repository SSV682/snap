package service

import (
	"fmt"
	"sync"

	"worker/internal/dto"
	"worker/internal/entity"
)

type taxFn func(price float64) float64
type createTradingStrategy func(inCh chan entity.Candle, outCh chan Event) TradingStrategy

type BackTestConfig struct {
	TradingInfoProvider TradingInfoProvider
}

type BackTestService struct {
	tradingInfoProvider TradingInfoProvider
}

func NewBackTestService(cfg *BackTestConfig) *TradingService {
	return &TradingService{
		tradingInfoProvider: cfg.TradingInfoProvider,
	}
}

func (c *TradingService) BackTest(filter dto.Filter) (entity.BackTestResult, error) {
	candles, err := c.tradingInfoProvider.HistoricCandles(filter.Ticker, filter.StartTime, filter.EndTime)
	if err != nil {
		return entity.BackTestResult{}, fmt.Errorf("historic candles: %v", err)
	}

	inCh := make(chan entity.Candle)
	//outCh := make(chan Event)

	//TODO: get from client
	tinkoffTax := func(price float64) float64 { return price * 0.05 / 100 }

	backtest := NewBackTest(inCh, NewVWAPStrategy, tinkoffTax)

	var (
		wg     sync.WaitGroup
		result entity.BackTestResult
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		result = backtest.Do()
	}()

	go func() {
		for _, candle := range candles {
			inCh <- candle
		}

		close(inCh)
	}()

	wg.Wait()

	return result, nil
}

type BackTest struct {
	strategy TradingStrategy

	pnl           float64
	calculateTax  taxFn
	numberDeal    int
	strategyInCh  chan entity.Candle
	strategyOutCh chan Event

	inCh chan entity.Candle
}

func NewBackTest(inCh chan entity.Candle, createStrategyFn createTradingStrategy, calculateTaxFn taxFn) BackTest {
	strategyInCh := make(chan entity.Candle)
	strategyOutCh := make(chan Event)

	s := createStrategyFn(strategyInCh, strategyOutCh)

	return BackTest{
		strategy:      s,
		calculateTax:  calculateTaxFn,
		strategyInCh:  strategyInCh,
		strategyOutCh: strategyOutCh,

		inCh: inCh,
	}
}

func (s *BackTest) Do() entity.BackTestResult {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		s.strategy.Do()
	}()

	go func() {
		wg.Wait()

		close(s.strategyOutCh)
	}()

	go func() {
		for candle := range s.inCh {
			s.strategyInCh <- candle
		}

		close(s.strategyInCh)
	}()

	var (
		dials       []entity.Dial
		numberDials int
		currentDial entity.Dial
		pnl         float64
	)

	for event := range s.strategyOutCh {
		numberDials++

		if event.Typ == Buy {
			currentDial = entity.Dial{
				Buy: event.Price,
			}

			continue
		}

		currentDial.Sell = event.Price
		currentDial.CalculatePNL()
		pnl += currentDial.PNL

		dials = append(dials, currentDial)
	}

	return entity.BackTestResult{
		NumberDial: numberDials,
		PNL:        pnl,
		Dials:      dials,
	}
}

package service

import (
	"fmt"
	"sync"

	"worker/internal/dto"
	"worker/internal/entity"
)

type createTradingStrategy func(inCh chan entity.Candle, outCh chan Event) TradingStrategy

type BackTestConfig struct {
	TradingInfoProvider TradingInfoProvider
	BrokerProvider      BrokerProvider
}

type BackTestService struct {
	tradingInfoProvider TradingInfoProvider
	brokerProvider      BrokerProvider
}

func NewBackTestService(cfg *BackTestConfig) *BackTestService {
	return &BackTestService{
		tradingInfoProvider: cfg.TradingInfoProvider,
		brokerProvider:      cfg.BrokerProvider,
	}
}

func (s *BackTestService) BackTest(filter dto.Filter) (entity.BackTestResult, error) {
	candles, err := s.tradingInfoProvider.HistoricCandles(filter.Ticker, filter.StartTime, filter.EndTime)
	if err != nil {
		return entity.BackTestResult{}, fmt.Errorf("historic candles: %v", err)
	}

	inCh := make(chan entity.Candle)

	//TODO: get from client

	backTest := NewBackTest(inCh, NewVWAPStrategy, s.brokerProvider.GetTaxFn())

	var (
		wg     sync.WaitGroup
		result entity.BackTestResult
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		result = backTest.Do()
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
	calculateTax  entity.TaxFn
	numberDeal    int
	strategyInCh  chan entity.Candle
	strategyOutCh chan Event

	inCh chan entity.Candle
}

func NewBackTest(inCh chan entity.Candle, createStrategyFn createTradingStrategy, calculateTaxFn entity.TaxFn) BackTest {
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
		currentDial.Period = event.Period
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

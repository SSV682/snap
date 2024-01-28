package service

import (
	"fmt"
	"sync"
	"time"
	"worker/internal/dto"
	"worker/internal/entity"
)

type taxFn func(price float64) float64
type createTradingStrategy func(inCh chan entity.Candle, outCh chan Event) TradingStrategy

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
}

type TradingStrategy interface {
	Do()
}

type Config struct {
	TradingInfoProvider TradingInfoProvider
	SignalCh            chan Event
}

type Calculator struct {
	tradingInfoProvider TradingInfoProvider
	signalCh            chan Event
}

func NewCalculator(cfg *Config) *Calculator {
	return &Calculator{
		tradingInfoProvider: cfg.TradingInfoProvider,
		signalCh:            cfg.SignalCh,
	}
}

func (c *Calculator) Backtest(filter dto.Filter) (entity.BackTestResult, error) {
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

type BackTestStrategy struct {
	strategy      TradingStrategy
	pnl           float64
	calculateTax  taxFn
	numberDeal    int
	strategyInCh  chan entity.Candle
	strategyOutCh chan Event

	inCh  chan entity.Candle
	outCh chan Event
}

func NewBackTest(inCh chan entity.Candle, createStrategyFn createTradingStrategy, calculateTaxFn taxFn) BackTestStrategy {
	strategyInCh := make(chan entity.Candle)
	strategyOutCh := make(chan Event)

	s := createStrategyFn(strategyInCh, strategyOutCh)

	return BackTestStrategy{
		strategy:      s,
		calculateTax:  calculateTaxFn,
		strategyInCh:  strategyInCh,
		strategyOutCh: strategyOutCh,

		inCh: inCh,
	}
}

func (s *BackTestStrategy) Do() entity.BackTestResult {
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

type VWAPStrategy struct {
	bestPriceForThisPeriod float64
	dealPrice              float64
	numbersPeriod          int
	candles                []entity.Candle

	inCh  chan entity.Candle
	outCh chan Event

	positiveFlag     bool
	startSellingFlag bool
}

func NewVWAPStrategy(inCh chan entity.Candle, outCh chan Event) TradingStrategy {
	return &VWAPStrategy{
		inCh:  inCh,
		outCh: outCh,
	}
}

func (c *VWAPStrategy) Do() {
	for candle := range c.inCh {

		c.numbersPeriod++
		if len(c.candles) < 5 {
			c.candles = append(c.candles, candle)

			continue
		}

		c.candles = append(c.candles[1:], candle)

		metricVWAP := c.calculateVWAPMetric()

		if c.positiveFlag && c.bestPriceForThisPeriod < candle.Close {
			c.bestPriceForThisPeriod = candle.Close
		}

		if !c.startSellingFlag && metricVWAP <= candle.Close && c.positiveFlag && c.dealPrice != 0 {
			c.startSellingFlag = true
		}

		if c.startSellingFlag && c.dealPrice != 0 && candle.Close < ((c.bestPriceForThisPeriod-c.dealPrice)/2)+c.dealPrice {
			c.generateSellEvent(candle)

			continue
		}

		if metricVWAP > candle.Close && !c.positiveFlag {
			c.generateBuyEvent(candle)

			continue
		}
	}
}

func (c *VWAPStrategy) generateSellEvent(candle entity.Candle) {
	c.positiveFlag = false
	c.dealPrice = 0
	c.startSellingFlag = false

	c.outCh <- Event{
		Typ:       Sell,
		Price:     candle.Close,
		Period:    c.numbersPeriod,
		BestPrice: c.bestPriceForThisPeriod,
	}

	c.numbersPeriod = 0
	c.bestPriceForThisPeriod = 0
}

func (c *VWAPStrategy) generateBuyEvent(candle entity.Candle) {
	c.positiveFlag = true
	c.dealPrice = candle.Close

	c.outCh <- Event{
		Typ:       Buy,
		Price:     candle.Close,
		Period:    c.numbersPeriod,
		BestPrice: c.bestPriceForThisPeriod,
	}

	c.numbersPeriod = 0
}

func (c *VWAPStrategy) calculateVWAPMetric() float64 {
	var (
		tradingVolume   int64
		financialVolume float64
	)

	for _, cv := range c.candles {
		financialVolume += cv.Close * float64(cv.Volume)
		tradingVolume += cv.Volume
	}

	metricVWAP := financialVolume / float64(tradingVolume)
	return metricVWAP
}

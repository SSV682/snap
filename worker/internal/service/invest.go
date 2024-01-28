package service

import (
	"fmt"
	"sync"
	"time"
	"worker/internal/dto"
	"worker/internal/entity"
)

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

func (c *Calculator) Backtest(filter dto.Filter) ([]string, error) {
	candles, err := c.tradingInfoProvider.HistoricCandles(filter.Ticker, filter.StartTime, filter.EndTime)
	if err != nil {
		return nil, fmt.Errorf("historic candles: %v", err)
	}

	inCh := make(chan entity.Candle)
	outCh := make(chan Event)

	strategy := NewVWAPStrategy(inCh, outCh)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		strategy.Do()
	}()

	go func() {
		for _, candle := range candles {
			inCh <- candle
		}

		close(inCh)
	}()

	go func() {
		wg.Wait()

		close(outCh)
	}()

	var result []string

	for event := range outCh {
		result = append(result, fmt.Sprintf("%s by price: %f", event.Typ, event.Price))

		if event.Typ != Sell {
			continue
		}

		result = append(result, fmt.Sprintf("PNL:  %f \n", strategy.pnl))
		result = append(result, fmt.Sprintf("Period: %d", strategy.numbersPeriod))
	}

	result = append(result, fmt.Sprintf("Number dial: %d", strategy.numberDeal))

	return result, nil
}

type VWAPStrategy struct {
	brokerTaxPercent       float64
	bestPriceForThisPeriod float64
	dealPrice              float64
	numbersPeriod          int
	numberDeal             int
	candles                []entity.Candle
	inCh                   chan entity.Candle
	outCh                  chan Event
	pnl                    float64

	positiveFlag     bool
	startSellingFlag bool
}

func NewVWAPStrategy(inCh chan entity.Candle, outCh chan Event) VWAPStrategy {
	return VWAPStrategy{
		brokerTaxPercent: 0.05,
		inCh:             inCh,
		outCh:            outCh,
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
	c.numberDeal++
	c.positiveFlag = false
	c.pnl += candle.Close - c.dealPrice
	c.dealPrice = 0
	c.calculateTax(candle.Close)
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
	c.numberDeal++
	c.calculateTax(candle.Close)

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

func (c *VWAPStrategy) calculateTax(price float64) {
	c.pnl -= (price * c.brokerTaxPercent) / 100
}

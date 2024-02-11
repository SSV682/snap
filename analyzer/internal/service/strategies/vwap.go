package strategies

import (
	"analyzer/internal/entity"
	"context"
	"sync"
)

type VWAPStrategy struct {
	inCh                       chan entity.Candle
	outCh                      chan entity.Event
	thresholdTakeProfitPercent float64
	thresholdStopLostPercent   float64

	cancelFn context.CancelFunc
	wg       sync.WaitGroup

	bestPriceForThisPeriod float64
	dealPrice              float64
	numbersPeriod          int
	candles                []entity.Candle
	positiveFlag           bool
	startSellingFlag       bool
}

type VWAPStrategyConfig struct {
	InCh                       chan entity.Candle
	OutCh                      chan entity.Event
	ThresholdTakeProfitPercent float64
	ThresholdStopLostPercent   float64
}

func NewVWAPStrategy(cfg *VWAPStrategyConfig) *VWAPStrategy {
	return &VWAPStrategy{
		//baseStrategy: baseStrategy{
		inCh:                       cfg.InCh,
		outCh:                      cfg.OutCh,
		thresholdTakeProfitPercent: cfg.ThresholdTakeProfitPercent,
		thresholdStopLostPercent:   cfg.ThresholdStopLostPercent,
		//},
	}
}

func (c *VWAPStrategy) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancelFn = cancel

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		c.run(ctx)
	}()
}

func (c *VWAPStrategy) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(c.outCh)

			return
		case candle := <-c.inCh:
			c.numbersPeriod++
			if len(c.candles) < 5 {
				c.candles = append(c.candles, candle)

				continue
			}

			c.candles = append(c.candles[1:], candle)

			metricVWAP := c.calculateVWAPMetric()

			if c.numbersPeriod < 2 {
				continue
			}

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
}

func (c *VWAPStrategy) Close() {
	c.cancelFn()

	c.wg.Wait()
}

func (c *VWAPStrategy) generateSellEvent(candle entity.Candle) {
	c.positiveFlag = false
	c.dealPrice = 0
	c.startSellingFlag = false

	c.outCh <- entity.Event{
		Typ:       entity.Sell,
		Price:     candle.Close,
		Period:    c.numbersPeriod,
		BestPrice: c.bestPriceForThisPeriod,
	}

	c.numbersPeriod = 0
	c.bestPriceForThisPeriod = 0
}

func (c *VWAPStrategy) generateBuyEvent(candle entity.Candle) {
	c.positiveFlag = true

	finalPrice := candle.Close
	c.dealPrice = finalPrice

	c.outCh <- entity.Event{
		Typ:       entity.Buy,
		Price:     finalPrice,
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

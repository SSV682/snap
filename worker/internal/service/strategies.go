package service

import "worker/internal/entity"

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

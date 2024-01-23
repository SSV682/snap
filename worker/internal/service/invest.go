package service

import (
	"fmt"
	"snap/worker/internal/entity"
	"time"
)

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
}

type Config struct {
	TradingInfoProvider TradingInfoProvider
}

type Calculator struct {
	tradingInfoProvider TradingInfoProvider
}

func NewCalculator(cfg *Config) *Calculator {
	return &Calculator{
		tradingInfoProvider: cfg.TradingInfoProvider,
	}
}

func (c *Calculator) HistoricCandles() {
	candles, err := c.tradingInfoProvider.HistoricCandles("TCSG", time.Now().Add(-1*time.Hour), time.Now())
	if err != nil {
		return
	}

	//var positiveFlag bool

	for i, candle := range candles {
		if i < 5 {
			continue
		}

		var (
			tradingVolume   int64
			financialVolume float64
		)

		for _, cv := range candles[i-4 : i] {
			financialVolume += cv.Close * float64(cv.Volume)
			tradingVolume += cv.Volume
		}

		metricVWAP := financialVolume / float64(tradingVolume)

		fmt.Printf("VWAP: %f", metricVWAP)
		fmt.Printf("Price: %f", candle.Close)
	}
}

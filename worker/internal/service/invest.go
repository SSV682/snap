package service

import (
	"time"
	"worker/internal/entity"
)

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
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

func (c *Calculator) HistoricCandles() {
	candles, err := c.tradingInfoProvider.HistoricCandles("TCSG", time.Now().Add(-48*time.Hour), time.Now().Add(-47*time.Hour))
	if err != nil {
		return
	}

	var positiveFlag bool

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

		//fmt.Printf("Price: %f, VWAP: %f \n", candle.Close, metricVWAP)

		if metricVWAP < candle.Close && !positiveFlag {
			positiveFlag = true

			c.signalCh <- Event{
				Typ:   Buy,
				Price: candle.Close,
			}
		}

		if metricVWAP > candle.Close && positiveFlag {
			positiveFlag = false
			c.signalCh <- Event{
				Typ:   Sell,
				Price: candle.Close,
			}
		}
	}
}

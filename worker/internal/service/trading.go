package service

import (
	"time"
	"worker/internal/entity"
)

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
}

type BrokerProvider interface {
	GetTaxFn() entity.TaxFn
}

type TradingStrategy interface {
	Do()
}

type Config struct {
	TradingInfoProvider TradingInfoProvider
	SignalCh            chan Event
}

type TradingService struct {
	tradingInfoProvider TradingInfoProvider
	signalCh            chan Event
}

func NewTradingService(cfg *Config) *TradingService {
	return &TradingService{
		tradingInfoProvider: cfg.TradingInfoProvider,
		signalCh:            cfg.SignalCh,
	}
}

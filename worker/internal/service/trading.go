package service

import (
	"time"
	"worker/internal/entity"
)

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
	LastCandle(ticker string) (entity.Candle, error)
	GetCurrencies() ([]entity.Instrument, error)
	GetStocks() ([]entity.Instrument, error)
	GetFutures() ([]entity.Instrument, error)
}

type BrokerProvider interface {
	GetTaxFn() entity.TaxFn
}

type TradingConfig struct {
	TradingInfoProvider TradingInfoProvider
	SignalCh            chan entity.Event
}

type TradingService struct {
	tradingInfoProvider TradingInfoProvider
	signalCh            chan entity.Event
}

func NewTradingService(cfg *TradingConfig) *TradingService {
	return &TradingService{
		tradingInfoProvider: cfg.TradingInfoProvider,
		signalCh:            cfg.SignalCh,
	}
}

func (s *TradingService) Currencies() ([]entity.Instrument, error) {
	return s.tradingInfoProvider.GetCurrencies()
}

func (s *TradingService) Stocks() ([]entity.Instrument, error) {
	return s.tradingInfoProvider.GetStocks()
}

func (s *TradingService) Futures() ([]entity.Instrument, error) {
	return s.tradingInfoProvider.GetFutures()
}

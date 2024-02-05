package backtest

import (
	"fmt"
	"worker/internal/dto"
	"worker/internal/entity"
	"worker/internal/service"
	"worker/internal/service/strategies"

	"github.com/pkg/errors"
)

type TradingStrategy interface {
	Run()
	Close()
}

type Config struct {
	TradingInfoProvider service.TradingInfoProvider
	BrokerProvider      service.BrokerProvider
}

type Service struct {
	tradingInfoProvider service.TradingInfoProvider
	brokerProvider      service.BrokerProvider
	strategies          map[string]TradingStrategy
	strategyInCh        chan entity.Candle
	strategyOutCh       chan entity.Event
}

func NewBackTestService(cfg *Config) *Service {
	strategyInCh := make(chan entity.Candle, 1)
	strategyOutCh := make(chan entity.Event, 1)

	return &Service{
		tradingInfoProvider: cfg.TradingInfoProvider,
		brokerProvider:      cfg.BrokerProvider,
		strategyInCh:        strategyInCh,
		strategyOutCh:       strategyOutCh,
		strategies: map[string]TradingStrategy{
			strategies.VWAP: strategies.NewVWAPStrategy(
				&strategies.VWAPStrategyConfig{
					InCh:                       strategyInCh,
					OutCh:                      strategyOutCh,
					ThresholdTakeProfitPercent: 0,
					ThresholdStopLostPercent:   0,
				},
			),
		},
	}
}

func (s *Service) BackTest(filter dto.Filter) (entity.BackTestResult, error) {
	usedStrategy, ok := s.strategies[filter.StrategyName]
	if !ok {
		return entity.BackTestResult{}, errors.New(fmt.Sprintf("unsupported strategy: %s", filter.StrategyName))
	}

	candles, err := s.tradingInfoProvider.HistoricCandles(filter.Ticker, filter.StartTime, filter.EndTime)
	if err != nil {
		return entity.BackTestResult{}, fmt.Errorf("historic candles: %v", err)
	}

	backTest := NewBackTest(&backTestConfig{
		candles:        candles,
		strategy:       usedStrategy,
		calculateTaxFn: s.brokerProvider.GetTaxFn(),
		strategyInCh:   s.strategyInCh,
		strategyOutCh:  s.strategyOutCh,
	})

	result := backTest.Do()

	return result, nil
}

package backtest

import (
	"fmt"
	"sync"
	"worker/internal/service/strategies"

	"worker/internal/dto"
	"worker/internal/entity"
	"worker/internal/service"

	"github.com/pkg/errors"
)

type TradingStrategy interface {
	Do()
	InChannel() chan entity.Candle
	OutChannel() chan entity.Event
}

type Config struct {
	TradingInfoProvider service.TradingInfoProvider
	BrokerProvider      service.BrokerProvider
}

type Service struct {
	tradingInfoProvider service.TradingInfoProvider
	brokerProvider      service.BrokerProvider
	strategies          map[string]TradingStrategy
}

func NewBackTestService(cfg *Config) *Service {
	return &Service{
		tradingInfoProvider: cfg.TradingInfoProvider,
		brokerProvider:      cfg.BrokerProvider,
		strategies: map[string]TradingStrategy{
			strategies.VWAP: strategies.NewVWAPStrategy(make(chan entity.Candle, 1), make(chan entity.Event, 1), cfg.BrokerProvider.GetTaxFn()),
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

	inCh := make(chan entity.Candle)

	backTest := NewBackTest(backTestConfig{
		inCh:           inCh,
		strategy:       usedStrategy,
		calculateTaxFn: s.brokerProvider.GetTaxFn(),
	})

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

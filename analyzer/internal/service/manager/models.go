package manager

import (
	"sync"
	"time"
	"worker/internal/entity"
)

type TradingTime struct {
	Hour   int
	Minute int
}

func (t TradingTime) Before() bool {
	now := time.Now()
	if t.Hour < now.Hour() && t.Minute < now.Minute() {
		return true
	}

	return false
}

func (t TradingTime) After() bool {
	now := time.Now()
	if t.Hour < now.Hour() && t.Minute < now.Minute() {
		return false
	}

	return true
}

type StrategySetting struct {
	strategyCh       chan<- entity.Candle
	strategyTimeFrom time.Time
	strategyTimeTo   time.Time
	tradingTimeFrom  TradingTime
	tradingTimeTo    TradingTime
}

type ConfigStrategySetting struct {
	StrategyCh       chan<- entity.Candle
	StrategyTimeFrom time.Time
	StrategyTimeTo   time.Time
	TradingTimeFrom  TradingTime
	TradingTimeTo    TradingTime
}

//func NewStrategySetting(cfg ConfigStrategySetting) StrategySetting {
//	return StrategySetting{
//		strategyCh:       cfg.StrategyCh,
//		strategyTimeFrom: cfg.StrategyTimeFrom,
//		strategyTimeTo:   cfg.StrategyTimeTo,
//		tradingTimeFrom:  cfg.TradingTimeFrom,
//		tradingTimeTo:    cfg.TradingTimeTo,
//	}
//}

func (s *StrategySetting) IsTradeAvailable() bool {
	if s.strategyTimeTo.Before(time.Now()) {
		return false
	}

	if s.strategyTimeFrom.After(time.Now()) {
		return false
	}

	if s.tradingTimeFrom.Before() {
		return false
	}

	if s.tradingTimeTo.After() {
		return false
	}

	return true
}

type StrategyStorage struct {
	mu      sync.RWMutex
	storage map[string]StrategySetting
}

func NewStrategyStorage() StrategyStorage {
	return StrategyStorage{
		storage: make(map[string]StrategySetting),
	}
}

func (s *StrategyStorage) getTickers() []string {
	result := make([]string, len(s.storage))
	index := 0

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.storage {
		result[index] = k
		index++
	}

	return result
}

func (s *StrategyStorage) getStrategy(ticker string) chan<- entity.Candle {
	s.mu.Lock()
	defer s.mu.Unlock()

	if setting, ok := s.storage[ticker]; ok {
		return setting.strategyCh
	}

	return nil
}

func (s *StrategyStorage) followNew(ticker string, strategyCh chan entity.Candle, setting *entity.StrategySettings) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.storage[ticker] = StrategySetting{
		strategyCh:       strategyCh,
		strategyTimeFrom: setting.StrategyTimeFrom,
		strategyTimeTo:   setting.StrategyTimeTo,
		tradingTimeFrom: TradingTime{
			Hour:   setting.TradingTimeFrom.Hour(),
			Minute: setting.TradingTimeFrom.Minute(),
		},
		tradingTimeTo: TradingTime{
			Hour:   setting.TradingTimeTo.Hour(),
			Minute: setting.TradingTimeTo.Minute(),
		},
	}
}

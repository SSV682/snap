package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"worker/internal/entity"
	"worker/internal/service/strategies"
)

type StrategySettingsRepository interface {
	List(ctx context.Context) ([]*entity.StrategySettings, error)
}

type BrokerProvider interface {
	GetTaxFn() entity.TaxFn
}

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
	LastCandle(ticker string) (entity.Candle, error)
}

type Strategy interface {
	Run()
	Close()
}

type Manager struct {
	settingsRepository StrategySettingsRepository
	brokerProvider     BrokerProvider
	infoProvider       TradingInfoProvider
	settingsStorage    StrategyStorage

	strategies []Strategy

	inCh  chan entity.Candle
	outCh chan entity.Event

	cancelFn context.CancelFunc
	wg       sync.WaitGroup
}

type Config struct {
	SettingsRepository StrategySettingsRepository
	BrokerProvider     BrokerProvider
	InfoProvider       TradingInfoProvider
}

func NewManager(cfg Config) *Manager {
	inCh := make(chan entity.Candle, 1)
	outCh := make(chan entity.Event, 1)

	vwapStrategy := strategies.NewVWAPStrategy(&strategies.VWAPStrategyConfig{
		InCh:                       inCh,
		OutCh:                      outCh,
		ThresholdTakeProfitPercent: 0,
		ThresholdStopLostPercent:   0,
	})

	return &Manager{
		settingsRepository: cfg.SettingsRepository,
		brokerProvider:     cfg.BrokerProvider,
		infoProvider:       cfg.InfoProvider,
		strategies:         []Strategy{vwapStrategy},
		settingsStorage:    NewStrategyStorage(),
		inCh:               inCh,
		outCh:              outCh,
	}
}

// Run runs the manager
func (m *Manager) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancelFn = cancel

	for i := range m.strategies {
		m.strategies[i].Run()
	}

	//run manager
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		m.run(ctx)
	}()
}

// loadStrategySettings loads strategy settings from storage and sets them to the strategies
func (m *Manager) loadStrategySettings(ctx context.Context) error {
	settings, err := m.settingsRepository.List(ctx)
	if err != nil {
		return fmt.Errorf("get strategy settings: %v", err)
	}

	for _, s := range settings {
		m.settingsStorage.followNew(s.Ticker, m.inCh, s)
	}

	return nil
}

func (m *Manager) run(ctx context.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	if err := m.loadStrategySettings(ctx); err != nil {
		//TODO: log error
		return
	}

	for {
		select {
		case <-ticker.C:
			candle, err := m.infoProvider.LastCandle(m.settingsStorage.getTickers()[0])
			if err != nil {
				//TODO: wrap error
				continue
			}

			m.inCh <- candle
		case event := <-m.outCh:
			fmt.Printf("event: %v", event)
			//m.externalCh <- event
		case <-ctx.Done():

			close(m.inCh)

			return
		}
	}
}

// Close closes all strategies and stops manager
func (m *Manager) Close() error {
	m.cancelFn()

	for i := range m.strategies {
		m.strategies[i].Close()
	}

	m.wg.Wait()

	return nil
}

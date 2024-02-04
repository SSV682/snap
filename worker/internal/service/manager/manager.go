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
	GetStrategySettings() ([]*entity.StrategySettings, error)
}

type BrokerProvider interface {
	GetTaxFn() entity.TaxFn
}

type TradingInfoProvider interface {
	HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error)
	LastCandle(ticker string) (entity.Candle, error)
}

type Manager struct {
	settingsRepo    StrategySettingsRepository
	brokerProvider  BrokerProvider
	infoProvider    TradingInfoProvider
	settingsStorage StrategyStorage

	inCh       chan entity.Candle
	outCh      chan entity.Event
	externalCh chan entity.Event

	cancelFn context.CancelFunc
	wg       sync.WaitGroup
}

type Config struct {
	ExternalCh chan entity.Event
}

func NewManager(cfg Config) *Manager {
	return &Manager{
		inCh:       make(chan entity.Candle, 1),
		outCh:      make(chan entity.Event, 1),
		externalCh: cfg.ExternalCh,
	}
}

func (m *Manager) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancelFn = cancel

	vwapStrategy := strategies.NewVWAPStrategy(m.inCh, m.outCh, m.brokerProvider.GetTaxFn())

	//run strategy
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		vwapStrategy.Do()
	}()

	//run manager
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		m.run(ctx)
	}()

	m.wg.Wait()
}

func (m *Manager) loadStrategySettings() error {
	settings, err := m.settingsRepo.GetStrategySettings()
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

	if err := m.loadStrategySettings(); err != nil {
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
			m.externalCh <- event
		case <-ctx.Done():
			close(m.inCh)

			return
		}
	}
}

func (m *Manager) Close() {
	m.cancelFn()

	m.wg.Wait()
}

package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"analyzer/internal/entity"
	"analyzer/internal/service/strategies"

	log "github.com/sirupsen/logrus"
)

type StrategySettingsRepository interface {
	List(ctx context.Context) ([]*entity.StrategySettings, error)
	Create(ctx context.Context, setting *entity.StrategySettings) (*entity.StrategySettings, error)
	Delete(ctx context.Context, id int64) error
}

type BrokerProvider interface {
	GetTaxFn() entity.TaxFn
}

type SolverClient interface {
	MakeDecision(ctx context.Context, event entity.Event) error
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
	solverClient       SolverClient

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
	SolverClient       SolverClient
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
		solverClient:       cfg.SolverClient,

		strategies:      []Strategy{vwapStrategy},
		settingsStorage: NewStrategyStorage(),
		inCh:            inCh,
		outCh:           outCh,
	}
}

// Run runs the manager
func (m *Manager) Run() error {
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

	return nil
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
			log.Infof("event: %v", event)
			err := m.solverClient.MakeDecision(ctx, event)
			if err != nil {
				log.Errorf("make decision: %v", err)
			}
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

func (m *Manager) ListActualStrategiesSettings(ctx context.Context) ([]*entity.StrategySettings, error) {
	settings, err := m.settingsRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("get strategy settings: %v", err)
	}

	return settings, nil
}

func (m *Manager) CreateSetting(ctx context.Context, setting *entity.StrategySettings) (settings *entity.StrategySettings, err error) {
	settings, err = m.settingsRepository.Create(ctx, setting)
	if err != nil {
		return nil, fmt.Errorf("create strategy setting: %v", err)
	}

	return setting, nil
}

func (m *Manager) DeleteSetting(ctx context.Context, id int64) error {
	if err := m.settingsRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete strategy setting: %v", err)
	}

	return nil
}

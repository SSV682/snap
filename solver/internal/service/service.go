package solver

import (
	"context"
	"fmt"
	"solver/internal/entity"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
// orderDirections is a map of order directions.

)

type BrokerProvider interface {
	GetFreeMoney() (int64, error)
	PostOrder(order entity.Order) error
	GetQuantityAvailabilityInstrument(ticker string) (int64, error)
}

// Service is the service for the solver.
type Service struct {
	broker BrokerProvider
	inCh   <-chan entity.Event

	orderFactory   *OrderFactory
	dialTypeSwitch entity.EventType
	cancelFn       context.CancelFunc
	wg             sync.WaitGroup
}

type Config struct {
	Broker BrokerProvider
	InCh   <-chan entity.Event
}

// NewService creates a new solver.
func NewService(cfg Config) *Service {
	return &Service{
		broker:         cfg.Broker,
		inCh:           cfg.InCh,
		dialTypeSwitch: entity.Sell,
	}
}

func (s *Service) Run() error {
	var ctx context.Context

	ctx, s.cancelFn = context.WithCancel(context.Background())

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		s.run(ctx)
	}()

	return nil
}

// run runs the service.
func (s *Service) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-s.inCh:
			if err := s.makeDecision(ctx, event); err != nil {
				log.Errorf("make decision: %v", err)
			}

			s.dialTypeSwitch = event.EventType
		}
	}
}

// makeDecision makes a decision for the given event.
func (s *Service) makeDecision(_ context.Context, event entity.Event) error {
	order, err := s.orderFactory.Create(event)
	if err != nil {
		return err
	}

	if err = s.broker.PostOrder(order); err != nil {
		return fmt.Errorf("post order: %v", err)
	}

	log.Infof("Order placed for %s, direction %s, quantity %d", event.Ticker, order.EventType, order.Quantity)

	return nil
}

// Close stops the service.
func (s *Service) Close() error {
	s.cancelFn()

	s.wg.Wait()

	return nil
}

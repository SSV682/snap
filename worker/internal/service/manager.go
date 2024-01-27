package service

import (
	"context"
	"fmt"
	"sync"
)

type EventType string

const (
	Buy  EventType = "buy"
	Sell EventType = "sell"
)

type Event struct {
	Typ   EventType
	Price float64
}

type Manager struct {
	inputCh  <-chan Event
	cancelFn context.CancelFunc
	wg       sync.WaitGroup
}

func NewManager(inputCh chan Event) *Manager {
	return &Manager{
		inputCh: inputCh,
	}
}

func (m *Manager) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancelFn = cancel

	m.wg.Add(1)

	go func() {
		defer m.wg.Done()

		m.run(ctx)
	}()
}

func (m *Manager) run(ctx context.Context) {
	for {
		select {
		case event := <-m.inputCh:
			fmt.Printf("%s by price: %f \n", event.Typ, event.Price)
		case <-ctx.Done():
			return
		}
	}
}

func (m *Manager) Close() {
	m.cancelFn()

	m.wg.Wait()
}

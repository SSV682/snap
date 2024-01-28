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
	Typ       EventType
	Price     float64
	Period    int
	BestPrice float64
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
	var (
		fullDiff        float64
		currentPurchase float64
		//bestFullDiff    float64
	)
	for {
		select {
		case event := <-m.inputCh:
			switch event.Typ {
			case Buy:
				currentPurchase = event.Price
			case Sell:
				if currentPurchase == 0 {
					continue
				}

				fullDiff += event.Price - currentPurchase
				currentPurchase = 0
			}

			fmt.Printf("%s by price: %f \n", event.Typ, event.Price)
			if currentPurchase == 0 {
				fmt.Printf("Current diff: %f, current period: %d \n", fullDiff, event.Period)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (m *Manager) Close() {
	m.cancelFn()

	m.wg.Wait()
}

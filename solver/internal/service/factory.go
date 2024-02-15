package solver

import (
	"fmt"
	"solver/internal/entity"

	log "github.com/sirupsen/logrus"
)

type OrderFactory struct {
	broker BrokerProvider
}

func NewOrderFactory(b BrokerProvider) *OrderFactory {
	return &OrderFactory{
		broker: b,
	}
}

func (f *OrderFactory) Create(event entity.Event) (entity.Order, error) {
	var (
		err      error
		quantity int64
	)

	switch event.EventType {
	case entity.Buy:
		freeMoney, err := f.broker.GetFreeMoney()
		if err != nil {
			return entity.Order{}, fmt.Errorf("get free money: %v", err)
		}

		//TODO: delete it
		log.Infof("Free money: %d", freeMoney)

		if float64(freeMoney) < event.Price {
			return entity.Order{}, fmt.Errorf("not enough money: %d", freeMoney)
		}

		quantity = int64(float64(freeMoney)/event.Price) - 1

		if quantity <= 0 {
			return entity.Order{}, fmt.Errorf("could not buy")
		}
	case entity.Sell:
		quantity, err = f.broker.GetQuantityAvailabilityInstrument(event.Ticker)
		if err != nil {
			return entity.Order{}, fmt.Errorf("get quantity availability instrument: %v", err)
		}

		if quantity <= 0 {
			return entity.Order{}, fmt.Errorf("could not sell")
		}

		log.Infof("%d elements of %s", quantity, event.Ticker)
	}

	return entity.Order{
		Ticker:    event.Ticker,
		EventType: event.EventType,
		Quantity:  quantity,
	}, err
}

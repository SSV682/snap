package solver

import (
	"errors"
	"solver/internal/entity"
)

// ToEventType converts a string to an EventType.
func ToEventType(s string) (entity.EventType, error) {
	switch s {
	case string(entity.Buy):
		return entity.Buy, nil
	case string(entity.Sell):
		return entity.Sell, nil
	default:
		return "", errors.New("unknown event")
	}
}

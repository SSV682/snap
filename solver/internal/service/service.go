package service

import (
	"context"
	"solver/internal/entity"

	log "github.com/sirupsen/logrus"
)

type BrokerProvider interface {
}

type Solver struct {
}

func (s *Solver) MakeDecision(_ context.Context, event entity.Event) error {
	log.Infof("coming event: %v", event)

	return nil
}

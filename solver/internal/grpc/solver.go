package solver

import (
	"context"

	"solver/internal/entity"

	solverv1 "github.com/SSV682/snap/protos/gen/go/solver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	solverv1.UnimplementedSolverServer

	outCh chan entity.Event
}

func RegisterServerAPI(s *grpc.Server, ch chan entity.Event) {
	solverv1.RegisterSolverServer(s, &ServerAPI{
		outCh: ch,
	})
}

func (s *ServerAPI) MakeDecision(_ context.Context, event *solverv1.EventRequest) (*emptypb.Empty, error) {
	if event.GetTicker() == "" {
		return nil, status.Error(codes.InvalidArgument, "ticker is required")
	}

	var et entity.EventType

	switch event.GetEventType() {
	case solverv1.EventType_EVENT_TYPE_SELL:
		et = entity.Sell
	case solverv1.EventType_EVENT_TYPE_BUY:
		et = entity.Buy
	}

	//TODO: wrap error
	s.outCh <- entity.Event{
		Ticker:    event.GetTicker(),
		EventType: et,
		Price:     float64(event.GetPrice()),
	}

	return &emptypb.Empty{}, nil
}

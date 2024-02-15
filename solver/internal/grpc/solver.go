package solver

import (
	"context"

	"solver/internal/entity"

	"github.com/SSV682/snap/protos/gen/go/analyzer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	analyzer_v1.UnimplementedSolverServer

	outCh chan entity.Event
}

func RegisterServerAPI(s *grpc.Server, ch chan entity.Event) {
	analyzer_v1.RegisterSolverServer(s, &ServerAPI{
		outCh: ch,
	})
}

func (s *ServerAPI) MakeDecision(_ context.Context, event *analyzer_v1.EventRequest) (*emptypb.Empty, error) {
	if event.GetTicker() == "" {
		return nil, status.Error(codes.InvalidArgument, "ticker is required")
	}

	if event.GetEventType() == "" {
		return nil, status.Error(codes.InvalidArgument, "type is required")
	}

	eventType, err := ToEventType(event.GetEventType())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//TODO: wrap error
	s.outCh <- entity.Event{
		Ticker:    event.GetTicker(),
		EventType: eventType,
		Price:     float64(event.GetPrice()),
	}

	return &emptypb.Empty{}, nil
}

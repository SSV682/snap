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

type Manager interface {
	MakeDecision(ctx context.Context, event entity.Event) error
}

type ServerAPI struct {
	analyzer_v1.UnimplementedSolverServer
	manager Manager
}

func RegisterServerAPI(s *grpc.Server, manager Manager) {
	analyzer_v1.RegisterSolverServer(s, &ServerAPI{
		manager: manager,
	})
}

func (s *ServerAPI) MakeDecision(ctx context.Context, event *analyzer_v1.EventRequest) (*emptypb.Empty, error) {
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
	if err := s.manager.MakeDecision(ctx, entity.Event{
		Ticker:    event.GetTicker(),
		EventType: eventType,
		Price:     float64(event.GetPrice()),
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

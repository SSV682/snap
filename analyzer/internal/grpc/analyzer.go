package grpc

import (
	"context"

	"analyzer/internal/entity"

	"github.com/SSV682/snap/protos/gen/go/analyzer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ManagerService interface {
	ListActualStrategiesSettings(ctx context.Context) ([]*entity.StrategySettings, error)
	CreateSetting(ctx context.Context, setting *entity.StrategySettings) (*entity.StrategySettings, error)
	DeleteSetting(ctx context.Context, id int64) error
}

type ServerAPI struct {
	analyzer_v1.UnimplementedAnalyzerServer

	managerService ManagerService
}

type Config struct {
	ManagerService ManagerService
}

func RegisterServerAPI(s *grpc.Server, cfg *Config) {
	analyzer_v1.RegisterAnalyzerServer(s, ServerAPI{managerService: cfg.ManagerService})
}

func (s ServerAPI) ListActualSettings(ctx context.Context, _ *emptypb.Empty) (*analyzer_v1.ActualSettingsResponse, error) {
	settings, err := s.managerService.ListActualStrategiesSettings(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := make([]*analyzer_v1.Setting, len(settings))
	for i := range settings {
		data[i] = &analyzer_v1.Setting{
			Id:             settings[i].ID,
			Ticker:         settings[i].Ticker,
			StrategyName:   settings[i].Strategy,
			Start:          timestamppb.New(settings[i].StrategyTimeFrom),
			End:            timestamppb.New(settings[i].StrategyTimeTo),
			StartInsideDay: timestamppb.New(settings[i].TradingTimeFrom),
			EndInsideDay:   timestamppb.New(settings[i].TradingTimeTo),
		}
	}

	return &analyzer_v1.ActualSettingsResponse{
		Data: data,
	}, err
}

func (s ServerAPI) CreateSetting(ctx context.Context, r *analyzer_v1.CreateSettingRequest) (*analyzer_v1.CreateSettingResponse, error) {
	setting, err := s.managerService.CreateSetting(ctx, &entity.StrategySettings{
		Ticker:           r.Ticker,
		Strategy:         r.StrategyName,
		StrategyTimeFrom: r.Start.AsTime(),
		StrategyTimeTo:   r.End.AsTime(),
		TradingTimeFrom:  r.StartInsideDay.AsTime(),
		TradingTimeTo:    r.EndInsideDay.AsTime(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &analyzer_v1.CreateSettingResponse{
		Id: setting.ID,
	}, nil
}

func (s ServerAPI) DeleteSetting(ctx context.Context, r *analyzer_v1.DeleteSettingRequest) (*emptypb.Empty, error) {
	if err := s.managerService.DeleteSetting(ctx, r.Id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

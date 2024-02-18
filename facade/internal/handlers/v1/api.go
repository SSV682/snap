package v1

import (
	"context"

	analyzer_v1 "github.com/SSV682/snap/protos/gen/go/analyzer"
	"github.com/golang/protobuf/ptypes/empty"
	routing "github.com/qiangxue/fasthttp-routing"
)

const (
	versionAPI = "/v1"
)

type AnalyzerClient interface {
	ListActualSettings(ctx context.Context) (*analyzer_v1.ActualSettingsResponse, error)
	CreateSetting(ctx context.Context, req *analyzer_v1.CreateSettingRequest) (*analyzer_v1.CreateSettingResponse, error)
	DeleteSetting(ctx context.Context, request *analyzer_v1.DeleteSettingRequest) (*empty.Empty, error)
}

//type BackTestService interface {
//	BackTest(filter dto.Filter) (entity.BackTestResult, error)
//}

//type TradingService interface {
//	Currencies() ([]entity.Instrument, error)
//	Stocks() ([]entity.Instrument, error)
//	Futures() ([]entity.Instrument, error)
//}

type Validator interface {
	Struct(object any) error
}

type API struct {
	analyzerClient AnalyzerClient
	validator      Validator
}

type Config struct {
	AnalyzerClient AnalyzerClient

	Validator Validator
}

func NewHandlers(cfg Config) *API {
	return &API{
		analyzerClient: cfg.AnalyzerClient,

		validator: cfg.Validator,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	v1Group := group.Group(versionAPI)

	a.registerAnalyzerHandlers(v1Group)
	//a.registerInfoHandlers(v1Group)
}

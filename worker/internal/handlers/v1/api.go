package v1

import (
	"worker/internal/dto"
	"worker/internal/entity"

	routing "github.com/qiangxue/fasthttp-routing"
)

const (
	versionAPI = "/v1"
)

type BackTestService interface {
	BackTest(filter dto.Filter) (entity.BackTestResult, error)
}

type TradingService interface {
	Currencies() ([]entity.Instrument, error)
}

type Validator interface {
	Struct(object any) error
}

type API struct {
	backTestService BackTestService
	tradingService  TradingService
	validator       Validator
}

type Config struct {
	BackTestService BackTestService
	TradingService  TradingService
	Validator       Validator
}

func NewInvestHandler(cfg Config) *API {
	return &API{
		backTestService: cfg.BackTestService,
		tradingService:  cfg.TradingService,
		validator:       cfg.Validator,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	v1Group := group.Group(versionAPI)

	a.registerBackTestHandlers(v1Group)
	a.registerInfoHandlers(v1Group)
}

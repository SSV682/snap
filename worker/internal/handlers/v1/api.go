package v1

import (
	"worker/internal/dto"
	"worker/internal/entity"

	routing "github.com/qiangxue/fasthttp-routing"
)

const (
	versionAPI = "/v1"
)

type InvestService interface {
	Backtest(filter dto.Filter) (entity.BackTestResult, error)
}

type Validator interface {
	Struct(object any) error
}

type API struct {
	investService InvestService
	validator     Validator
}

func NewInvestHandler(srv InvestService, v Validator) *API {
	return &API{
		investService: srv,
		validator:     v,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	v1Group := group.Group(versionAPI)

	a.registerInvestHandlers(v1Group)
}

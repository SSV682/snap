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

type Validator interface {
	Struct(object any) error
}

type API struct {
	backTestService BackTestService

	validator Validator
}

func NewInvestHandler(srv BackTestService, v Validator) *API {
	return &API{
		backTestService: srv,
		validator:       v,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	v1Group := group.Group(versionAPI)

	a.registerInvestHandlers(v1Group)
}

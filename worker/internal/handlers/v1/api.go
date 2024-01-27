package v1

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

const (
	versionAPI = "/v1"
)

type InvestService interface {
	HistoricCandles()
}

type API struct {
	investService InvestService
}

func NewInvestHandler(srv InvestService) *API {
	return &API{
		investService: srv,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	v1Group := group.Group(versionAPI)

	a.registerInvestHandlers(v1Group)
}

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
	invService InvestService
}

func NewInvestHandler(srv InvestService) *API {
	return &API{
		invService: srv,
	}
}

func (a *API) RegisterHandlers(group *routing.RouteGroup) {
	group = group.Group(versionAPI)

	a.registerInvestHandlers(group)
}

func (a *API) GetHistoricCandles(_ *routing.Context) error {
	a.invService.HistoricCandles()

	return nil
}

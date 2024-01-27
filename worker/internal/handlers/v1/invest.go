package v1

import routing "github.com/qiangxue/fasthttp-routing"

const (
	investURI  = "/invest"
	candlesURI = "/candles"
)

func (a *API) registerInvestHandlers(group *routing.RouteGroup) {
	investGroup := group.Group(investURI)

	investGroup.Get(candlesURI, a.GetHistoricCandles)
}

func (a *API) GetHistoricCandles(_ *routing.Context) error {
	a.investService.HistoricCandles()

	return nil
}

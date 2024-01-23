package v1

import routing "github.com/qiangxue/fasthttp-routing"

const (
	investURI  = "invest"
	candlesURI = "candles"
)

func (a *API) registerInvestHandlers(group *routing.RouteGroup) {
	deploymentsGroup := group.Group(investURI)

	deploymentsGroup.Get(candlesURI)
}

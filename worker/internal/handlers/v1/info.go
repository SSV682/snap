package v1

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

const (
	infoURI       = "/info"
	currenciesURI = "/currencies"
)

func (a *API) registerInfoHandlers(group *routing.RouteGroup) {
	infoGroup := group.Group(infoURI)

	infoGroup.Get(currenciesURI, a.Currencies)
}

func (a *API) Currencies(ctx *routing.Context) error {
	result, err := a.tradingService.Currencies()
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(NewCurrenciesResponseFromEntity(result))
}

package v1

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

const (
	infoURI       = "/info"
	currenciesURI = "/currencies"
	stocksURI     = "/stocks"
	futuresURI    = "/futures"
)

func (a *API) registerInfoHandlers(group *routing.RouteGroup) {
	infoGroup := group.Group(infoURI)

	infoGroup.Get(currenciesURI, a.Currencies)
	infoGroup.Get(stocksURI, a.Stocks)
	infoGroup.Get(futuresURI, a.Futures)
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

func (a *API) Stocks(ctx *routing.Context) error {
	result, err := a.tradingService.Stocks()
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(NewCurrenciesResponseFromEntity(result))
}

func (a *API) Futures(ctx *routing.Context) error {
	result, err := a.tradingService.Futures()
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(NewCurrenciesResponseFromEntity(result))
}

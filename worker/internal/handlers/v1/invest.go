package v1

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

const (
	investURI  = "/invest"
	candlesURI = "/candles"
)

const (
	ContentTypeApplicationJson = "application/json"
)

func (a *API) registerInvestHandlers(group *routing.RouteGroup) {
	investGroup := group.Group(investURI)

	investGroup.Post(candlesURI, a.Backtest)
}

func (a *API) Backtest(ctx *routing.Context) error {
	filter, err := newBacktestRequest(ctx.PostBody(), a.validator)
	if err != nil {
		return err
	}

	backtestResult, err := a.investService.Backtest(filter.ToDTO())
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(backtestResult)
}

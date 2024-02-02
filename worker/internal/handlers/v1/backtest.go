package v1

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

const (
	backTestURI = "/backtest"
	//candlesURI  = "/candles"
)

const (
	ContentTypeApplicationJson = "application/json"
)

func (a *API) registerInvestHandlers(group *routing.RouteGroup) {
	//investGroup := group.Group(backTestURI)

	group.Post(backTestURI, a.BackTest)
}

func (a *API) BackTest(ctx *routing.Context) error {
	filter, err := newBackTestRequest(ctx.PostBody(), a.validator)
	if err != nil {
		return err
	}

	result, err := a.backTestService.BackTest(filter.ToDTO())
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(BackTestResponseFromEntity(result))
}

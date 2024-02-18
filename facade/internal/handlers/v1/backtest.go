package v1

//import (
//	"encoding/json"
//
//	routing "github.com/qiangxue/fasthttp-routing"
//	"github.com/valyala/fasthttp"
//)
//
//const (
//	backTestURI = "/backtest"
//)
//
//func (a *API) registerBackTestHandlers(group *routing.RouteGroup) {
//	group.Post(backTestURI, a.BackTest)
//}
//
//func (a *API) BackTest(ctx *routing.Context) error {
//	filter, err := newBackTestRequest(ctx.PostBody(), a.validator)
//	if err != nil {
//		return err
//	}
//
//	result, err := a.backTestService.BackTest(filter.ToDTO())
//	if err != nil {
//		return err
//	}
//
//	ctx.SetStatusCode(fasthttp.StatusOK)
//	ctx.SetContentType(ContentTypeApplicationJson)
//
//	return json.NewEncoder(ctx).Encode(BackTestResponseFromEntity(result))
//}

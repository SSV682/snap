package v1

import (
	"encoding/json"
	"time"

	analyzer_v1 "github.com/SSV682/snap/protos/gen/go/analyzer"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	settingsURI = "/settings"

	argsID = "id"
)

func (a *API) registerAnalyzerHandlers(group *routing.RouteGroup) {
	group.Get(settingsURI, a.ListSettings)
	group.Post(settingsURI, a.CreateSetting)
	group.Delete(settingsURI, a.DeleteSetting)
}

func (a *API) ListSettings(ctx *routing.Context) error {
	result, err := a.analyzerClient.ListActualSettings(ctx)
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(ListSettingsResponseFromPB(result))
}

func (a *API) CreateSetting(ctx *routing.Context) error {
	req, err := newCreateRequest(ctx.PostBody(), a.validator)
	if err != nil {
		return err
	}

	csr := &analyzer_v1.CreateSettingRequest{
		Ticker:       req.Ticker,
		StrategyName: req.StrategyName,
	}

	if t, err := time.Parse(datetimeStringFormat, req.Start); err != nil {
		csr.Start = timestamppb.New(t)
	}
	if t, err := time.Parse(datetimeStringFormat, req.End); err != nil {
		csr.End = timestamppb.New(t)
	}
	if t, err := time.Parse(datetimeStringFormat, req.StartTimeInsideDay); err != nil {
		csr.StartInsideDay = timestamppb.New(t)
	}
	if t, err := time.Parse(datetimeStringFormat, req.EndTimeInsideDay); err != nil {
		csr.EndInsideDay = timestamppb.New(t)
	}

	result, err := a.analyzerClient.CreateSetting(ctx, csr)
	if err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return json.NewEncoder(ctx).Encode(createSettingResponse{ID: result.Id})
}

func newCreateRequest(data []byte, validatorIns Validator) (settings *createSettingRequest, err error) {
	if err = json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	if err = validate(settings, validatorIns); err != nil {
		return nil, err
	}

	return settings, nil
}

func (a *API) DeleteSetting(ctx *routing.Context) error {
	id, err := ctx.QueryArgs().GetUint(argsID)
	if err != nil {
		return err
	}

	if _, err = a.analyzerClient.DeleteSetting(ctx, &analyzer_v1.DeleteSettingRequest{Id: int64(id)}); err != nil {
		return err
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType(ContentTypeApplicationJson)

	return nil
}

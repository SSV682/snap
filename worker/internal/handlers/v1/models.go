package v1

import (
	"time"
	"worker/internal/dto"
	"worker/internal/entity"
)

type backTestRequest struct {
	StartTime    *int64  `json:"start_time" validate:"required"`
	EndTime      *int64  `json:"end_time" validate:"required"`
	Ticker       *string `json:"ticker" validate:"required"`
	StrategyName *string `json:"strategy_name" validate:"required"`
}

func (r *backTestRequest) ToDTO() dto.Filter {
	return dto.Filter{
		StartTime:    time.Unix(*r.StartTime, 0),
		EndTime:      time.Unix(*r.EndTime, 0),
		Ticker:       *r.Ticker,
		StrategyName: *r.StrategyName,
	}
}

type BackTestResponse struct {
	NumberDial int            `json:"number_dial"`
	PNL        float64        `json:"profit_and_lost"`
	Dials      []DialResponse `json:"dials"`
}

type DialResponse struct {
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	PNL    float64 `json:"profit_and_lost"`
	Period int     `json:"period"`
}

func BackTestResponseFromEntity(result entity.BackTestResult) BackTestResponse {
	dials := make([]DialResponse, len(result.Dials))

	for i := range result.Dials {
		dials[i] = DialResponse{
			Buy:    result.Dials[i].Buy,
			Sell:   result.Dials[i].Sell,
			PNL:    result.Dials[i].PNL,
			Period: result.Dials[i].Period,
		}
	}

	return BackTestResponse{
		NumberDial: result.NumberDial,
		PNL:        result.PNL,
		Dials:      dials,
	}
}

type CurrenciesResponse struct {
	Data []Instrument `json:"data"`
}

type Instrument struct {
	Name   string `json:"name"`
	Figi   string `json:"figi"`
	Ticker string `json:"ticker"`
}

func NewCurrenciesResponseFromEntity(instruments []entity.Instrument) CurrenciesResponse {
	data := make([]Instrument, len(instruments))

	for i := range instruments {
		data[i] = Instrument{
			Name:   instruments[i].Name,
			Figi:   instruments[i].Figi,
			Ticker: instruments[i].Ticker,
		}
	}

	return CurrenciesResponse{
		Data: data,
	}
}

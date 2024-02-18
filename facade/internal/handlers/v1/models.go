package v1

import "time"

//
//import (
//	"analyzer/internal/dto"
//	"analyzer/internal/entity"
//	"time"
//)
//
//type backTestRequest struct {
//	StartTime    *int64  `json:"start_time" validate:"required"`
//	EndTime      *int64  `json:"end_time" validate:"required"`
//	Ticker       *string `json:"ticker" validate:"required"`
//	StrategyName *string `json:"strategy_name" validate:"required"`
//}
//
//func (r *backTestRequest) ToDTO() dto.Filter {
//	return dto.Filter{
//		StartTime:    time.Unix(*r.StartTime, 0),
//		EndTime:      time.Unix(*r.EndTime, 0),
//		Ticker:       *r.Ticker,
//		StrategyName: *r.StrategyName,
//	}
//}
//
//type BackTestResponse struct {
//	NumberDial int            `json:"number_dial"`
//	PNL        float64        `json:"profit_and_lost"`
//	Dials      []DialResponse `json:"dials"`
//}
//
//type DialResponse struct {
//	Buy    float64 `json:"buy"`
//	Sell   float64 `json:"sell"`
//	PNL    float64 `json:"profit_and_lost"`
//	Period int     `json:"period"`
//}
//
//func BackTestResponseFromEntity(result entity.BackTestResult) BackTestResponse {
//	dials := make([]DialResponse, len(result.Dials))
//
//	for i := range result.Dials {
//		dials[i] = DialResponse{
//			Buy:    result.Dials[i].Buy,
//			Sell:   result.Dials[i].Sell,
//			PNL:    result.Dials[i].PNL,
//			Period: result.Dials[i].Period,
//		}
//	}
//
//	return BackTestResponse{
//		NumberDial: result.NumberDial,
//		PNL:        result.PNL,
//		Dials:      dials,
//	}
//}
//
//type CurrenciesResponse struct {
//	Data []Instrument `json:"data"`
//}
//
//type Instrument struct {
//	Name   string `json:"name"`
//	Figi   string `json:"figi"`
//	Ticker string `json:"ticker"`
//}
//
//func NewCurrenciesResponseFromEntity(instruments []entity.Instrument) CurrenciesResponse {
//	data := make([]Instrument, len(instruments))
//
//	for i := range instruments {
//		data[i] = Instrument{
//			Name:   instruments[i].Name,
//			Figi:   instruments[i].FIGI,
//			Ticker: instruments[i].Ticker,
//		}
//	}
//
//	return CurrenciesResponse{
//		Data: data,
//	}
//}

type settingResponse struct {
	ID             int64     `json:"id"`
	StrategyName   string    `json:"strategy_name"`
	Ticker         string    `json:"ticker"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	StartInsideDay time.Time `json:"start_inside_day"`
	EndInsideDay   time.Time `json:"end_inside_day"`
}

type ListSettingsResponse struct {
	Data []settingResponse `json:"data"`
}

type createSettingRequest struct {
	Start              *int64  `json:"start_time" validate:"required"`
	End                *int64  `json:"end_time" validate:"required"`
	StartTimeInsideDay *int64  `json:"start_inside_day" validate:"required"`
	EndTimeInsideDay   *int64  `json:"end_inside_day" validate:"required"`
	Ticker             *string `json:"ticker" validate:"required"`
	StrategyName       *string `json:"strategy_name" validate:"required"`
}

type createSettingResponse struct {
	ID int64 `json:"id"`
}

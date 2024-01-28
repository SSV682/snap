package v1

import (
	"time"
	"worker/internal/dto"
)

type backtestRequest struct {
	StartTime        *int64   `json:"start_time" validate:"required"`
	EndTime          *int64   `json:"end_time" validate:"required"`
	Ticker           *string  `json:"ticker" validate:"required"`
	BrokerTaxPercent *float64 `json:"broker_tax_percent" validate:"required"`
	StrategyName     *string  `json:"strategy_name" validate:"required"`
}

func (r *backtestRequest) ToDTO() dto.Filter {
	return dto.Filter{
		StartTime:        time.Unix(*r.StartTime, 0),
		EndTime:          time.Unix(*r.EndTime, 0),
		Ticker:           *r.Ticker,
		BrokerTaxPercent: *r.BrokerTaxPercent,
		StrategyName:     *r.StrategyName,
	}
}

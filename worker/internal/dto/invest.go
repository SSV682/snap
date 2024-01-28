package dto

import "time"

type Filter struct {
	StartTime        time.Time
	EndTime          time.Time
	Ticker           string
	BrokerTaxPercent float64
	StrategyName     string
}

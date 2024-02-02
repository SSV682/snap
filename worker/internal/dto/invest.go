package dto

import "time"

type Filter struct {
	StartTime    time.Time
	EndTime      time.Time
	Ticker       string
	StrategyName string
}

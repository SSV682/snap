package strategies

import (
	"context"
	"sync"

	"worker/internal/entity"
)

const (
	VWAP = "VWAP"
)

type baseStrategy struct {
	inCh                       chan entity.Candle
	outCh                      chan entity.Event
	thresholdTakeProfitPercent float64
	thresholdStopLostPercent   float64
	cancelFn                   context.CancelFunc
	wg                         sync.WaitGroup
}

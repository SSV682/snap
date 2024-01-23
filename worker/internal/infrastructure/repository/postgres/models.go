package postgres

import (
	"snap/worker/internal/entity"
	"time"
)

type CandleRow struct {
	Open   float64   `db:"open"`
	Close  float64   `db:"close"`
	High   float64   `db:"high"`
	Low    float64   `db:"low"`
	Volume int64     `db:"volume"`
	Time   time.Time `db:"time"`
}

func NewCandleRow(candle *entity.Candle) *CandleRow {
	return &CandleRow{
		Open:   candle.Open,
		Close:  candle.Close,
		High:   candle.High,
		Low:    candle.Low,
		Volume: candle.Volume,
		Time:   candle.Time,
	}

}

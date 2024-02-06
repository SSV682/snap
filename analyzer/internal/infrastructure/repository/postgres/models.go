package postgres

import (
	"database/sql/driver"
	"fmt"
	"time"
	"worker/internal/entity"
)

//type CandleRow struct {
//	Open   float64   `db:"open"`
//	Close  float64   `db:"close"`
//	High   float64   `db:"high"`
//	Low    float64   `db:"low"`
//	Volume int64     `db:"volume"`
//	Time   time.Time `db:"time"`
//}
//
//func NewCandleRow(candle *entity.Candle) *CandleRow {
//	return &CandleRow{
//		Open:   candle.Open,
//		Close:  candle.Close,
//		High:   candle.High,
//		Low:    candle.Low,
//		Volume: candle.Volume,
//		Time:   candle.Time,
//	}
//}

const MyTimeFormat = "15:04:05"

type MyTime time.Time

func (t *MyTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = MyTime(v)
	case nil:
		*t = MyTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(MyTimeFormat)), nil
}

func (t *MyTime) UnmarshalText(value string) error {
	dd, err := time.Parse(MyTimeFormat, value)
	if err != nil {
		return err
	}
	*t = MyTime(dd)
	return nil
}

type strategySettingsRow struct {
	Ticker           string    `db:"ticker"`
	Strategy         string    `db:"strategy"`
	StrategyTimeFrom time.Time `db:"strategy_time_from"`
	StrategyTimeTo   time.Time `db:"strategy_time_to"`
	TradingTimeFrom  MyTime    `db:"trading_time_from"`
	TradingTimeTo    MyTime    `db:"trading_time_to"`
}

func (r strategySettingsRow) ToModel() *entity.StrategySettings {
	return &entity.StrategySettings{
		Ticker:           r.Ticker,
		Strategy:         r.Strategy,
		StrategyTimeFrom: r.StrategyTimeFrom,
		StrategyTimeTo:   r.StrategyTimeTo,
		TradingTimeFrom:  time.Time(r.TradingTimeFrom),
		TradingTimeTo:    time.Time(r.TradingTimeTo),
	}
}

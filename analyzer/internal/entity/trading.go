package entity

import (
	"time"
)

type TaxFn func(price float64) float64

type EventType string

const (
	Buy  EventType = "buy"
	Sell EventType = "sell"
)

type Event struct {
	Typ    EventType
	Ticker string
	//TODO: remove this attributes
	Price     float64
	Period    int
	BestPrice float64
}

type Candle struct {
	Ticker string
	Open   float64   //Цена открытия за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	High   float64   //Максимальная цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Low    float64   //Минимальная цена за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Close  float64   //Цена закрытия за 1 инструмент. Для получения стоимости лота требуется умножить на лотность инструмента.
	Volume int64     //Объём торгов в лотах.
	Time   time.Time //Время свечи в часовом поясе UTC.
}

type BackTestResult struct {
	NumberDial int
	PNL        float64
	Dials      []Dial
}

type Dial struct {
	Buy    float64
	Sell   float64
	PNL    float64
	Period int
}

func (d *Dial) CalculatePNL() {
	d.PNL = d.Sell - d.Buy
}

type Instrument struct {
	Name   string
	FIGI   string
	Ticker string
}

type StrategySettings struct {
	ID               int64
	Ticker           string
	Strategy         string
	StrategyTimeFrom time.Time
	StrategyTimeTo   time.Time
	TradingTimeFrom  time.Time
	TradingTimeTo    time.Time
}

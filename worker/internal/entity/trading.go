package entity

import (
	"time"
)

type TaxFn func(price float64) float64

type Candle struct {
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
	Figi   string
	Ticker string
}

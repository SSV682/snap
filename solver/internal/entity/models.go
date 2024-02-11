package entity

type EventType string

const (
	Buy  EventType = "buy"
	Sell EventType = "sell"
)

type Event struct {
	Ticker    string
	EventType EventType
	Price     float64
}

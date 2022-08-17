package backtest

import "time"

type DealKind int

const (
	Sell DealKind = 100
	Buy  DealKind = 200
)

type Deal struct {
	D            time.Time
	TickerId     string
	Volume       int
	Lots         int
	Kind         DealKind
	Price        float64
	InvestResult float64
}

func (d *Deal) Sum() float64 {
	return float64(d.Volume) * d.Price
}

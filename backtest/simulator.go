package backtest

import (
	"math"
	"time"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/ent"
)

func ActualPortfolio(Strategy ent.Strategy, Market cube.Cube, D time.Time, Source *Portfolio) *Portfolio {

	fixedTickersShare := 0
	for _, r := range Strategy.Edges.FixedTickers {
		if r.IsUsed {
			fixedTickersShare += r.Share
		}
	}

	result := Portfolio{}
	if (len(Strategy.Edges.FixedTickers) == 0 && len(Strategy.Edges.Factors) == 0) || (Strategy.MaxTickers == 0) {
		return &result
	}

	// take into account fixed tickers
	startFuel := (Source.CurrentBalance() + Source.RUB) / 100.0 * (100.0 - float64(fixedTickersShare))
	result.RUB = Source.CurrentBalance() + Source.RUB

	if startFuel > 0 {

		//factors, etc.

	}

	// fixed tickers
	for _, r := range Strategy.Edges.FixedTickers {
		if !r.IsUsed {
			continue
		}
		cell := Market.CellsByTickerByDate(r.Ticker, D, true)
		if cell != nil {
			lots := int(math.Trunc((Source.CurrentBalance() + Source.RUB) / 100 * float64(r.Share) / (float64(cell.Emission.LotSize) * cell.Quote.C)))
			if lots > 0 {
				result.BuyLots(cell, lots)
			}
		}
	}

	return &result

}

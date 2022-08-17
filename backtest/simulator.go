package backtest

import (
	"math"
	"sort"
	"time"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/schema"
)

func ActualPortfolio(Strategy ent.Strategy, Market *cube.Cube, D time.Time, Source *Portfolio) *Portfolio {

	//TODO samepolicy

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
	flexRUB := (Source.CurrentBalance() + Source.RUB) / 100.0 * (100.0 - float64(fixedTickersShare))
	result.RUB = Source.CurrentBalance() + Source.RUB

	// first, process factor+filter combinations
	if flexRUB > 0 {

		all := Market.GetCellsByDate(D)

		buffer := make(map[*cube.Cell]map[int]float64, 0)

		for _, c := range all {
			if c.R2 == nil {
				continue
			}
			if k := c.Quote.Edges.Ticker.Kind; k != schema.TickerKind_Stock && k != schema.TickerKind_Bond || k == schema.TickerKind_StockPref {
				continue
			}
			if !Filter(c, Strategy, Market, D) {
				continue
			}

			pfItem := Source.FindByTickerId(c.TickerId())

			for _, factor := range Strategy.Edges.Factors {
				if !factor.IsUsed {
					continue
				}
				rv := c.RepValue(Market, factor.RK, factor.RVT)
				if factor.Inverse {
					if rv == 0 {
						rv = 0.00000001
					} else {
						rv = 1.0 / rv
					}
				}
				r2 := rv * factor.K
				if pfItem != nil {
					r2 += factor.Gist
				}

				mapt, ok := buffer[c]
				if !ok {
					mapt = make(map[int]float64, 0)
					buffer[c] = mapt
				}
				mapt[factor.LineNum] = r2

			}

		}

		type piece struct {
			c   *cube.Cell
			w   float64
			sum float64
		}

		pieces := make([]piece, len(buffer))
		idx := 0
		for t, v := range buffer {
			pieces[idx].c = t
			for _, r2 := range v {
				pieces[idx].w += r2
			}
			idx++
		}

		sort.Slice(pieces, func(i, j int) bool { return pieces[i].w > pieces[j].w })
		pieces = pieces[:Strategy.MaxTickers]
		ws := 0.0
		for _, p := range pieces {
			ws += p.w
		}
		for idx, p := range pieces {
			pieces[idx].sum = p.w / ws * flexRUB
		}
		for _, sp := range pieces {
			lotprice := float64(sp.c.Emission.LotSize) * sp.c.Quote.C
			if lotprice > sp.sum {
				continue
			}
			lots := int(math.Trunc(sp.sum / lotprice))
			result.BuyLots(sp.c, lots)
		}

	}

	// second, apply fixed tickers
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

func Filter(Src *cube.Cell, Strategy ent.Strategy, Market *cube.Cube, D time.Time) bool {

	for _, f := range Strategy.Edges.Filters {

		if !f.IsUsed {
			continue
		}

		switch f.LeftValueKind {
		case domains.FVK_Ticker:
			{
				if f.Operation == domains.FilterOp_Eq && Src.TickerId() != f.RightValueStr {
					return false
				}
				if f.Operation == domains.FilterOp_Ne && Src.TickerId() == f.RightValueStr {
					return false
				}
			}
		case domains.FVK_Industry:
			{
				if f.Operation == domains.FilterOp_Eq && Src.Industry.ID != f.RightValueStr {
					return false
				}
				if f.Operation == domains.FilterOp_Ne && Src.Industry.ID == f.RightValueStr {
					return false
				}
			}
		case domains.FVK_ReportValue:
			{
				lval := Src.RepValue(cube.Market, f.LeftReportValue, f.LeftReportValueType)
				rval := f.RightValueFloat
				if f.Operation == domains.FilterOp_Eq && lval != rval {
					return false
				}
				if f.Operation == domains.FilterOp_Ge && lval < rval {
					return false
				}
				if f.Operation == domains.FilterOp_Gt && lval <= rval {
					return false
				}
				if f.Operation == domains.FilterOp_Le && lval > rval {
					return false
				}
				if f.Operation == domains.FilterOp_Lt && lval >= rval {
					return false
				}
				if f.Operation == domains.FilterOp_Ne && lval == rval {
					return false
				}
			}
		}
	}
	return true
}

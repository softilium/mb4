package backtest

import (
	"log"
	"math"
	"sort"
	"time"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/schema"
)

func ActualPortfolio(Strategy *ent.Strategy, Market *cube.Cube, D time.Time, Source *Portfolio, ImplicitTicker *ent.Ticker) *Portfolio {

	if ImplicitTicker != nil {
		ir := Portfolio{RUB: Source.CurrentBalance() + Source.RUB}
		cell := Market.CellsByTickerByDate(ImplicitTicker.ID, D, true)
		if cell == nil {
			log.Fatalf("ActualPortfolio: no cell for %s on %s\n", ImplicitTicker.Descr, D)
			return nil
		}
		lots := int(math.Trunc(Source.CurrentBalance() + Source.RUB/(float64(cell.LotSize())+cell.Quote.C)))
		ir.BuyLots(cell, lots)
		return &ir
	}

	fixedTickersShare := 0
	fixedTickersCnt := 0
	for _, r := range Strategy.Edges.FixedTickers {
		if r.IsUsed {
			fixedTickersShare += r.Share
			fixedTickersCnt++
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
			if k := c.Quote.Edges.Ticker.Kind; !(k == schema.TickerKind_Stock || k == schema.TickerKind_StockPref) {
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

		sort.Slice(pieces, func(i, j int) bool {
			if pieces[i].w == pieces[j].w {
				return pieces[i].c.Quote.C < pieces[j].c.Quote.C
			}
			return pieces[i].w > pieces[j].w
		})

		//step 1 - all tickers
		ws := 0.0
		for _, p := range pieces {
			ws += p.w
		}
		for idx, p := range pieces {
			pieces[idx].sum = p.w / ws * flexRUB
		}
		pieces2 := make([]piece, 0, len(pieces))

		stocks := make(map[*ent.Emitent]int, 0)
		prefs := make(map[*ent.Emitent]int, 0)

		for _, p := range pieces {
			lotprice := float64(p.c.LotSize()) * p.c.Quote.C
			if lotprice > p.sum {
				continue
			}

			// apply sameEmitent policy
			switch Strategy.SameEmitent {
			case domains.SameEmitentPolicy_PreferPrefs:
				if p.c.Quote.Edges.Ticker.Kind == schema.TickerKind_StockPref {
					idx, ok := stocks[p.c.Quote.Edges.Ticker.Edges.Emitent]
					if ok {
						pieces2 = append(pieces2[:idx], pieces2[idx+1:]...)
					}
				} else if p.c.Quote.Edges.Ticker.Kind == schema.TickerKind_Stock {
					_, ok := prefs[p.c.Quote.Edges.Ticker.Edges.Emitent]
					if ok {
						continue
					}
				}
			case domains.SameEmitentPolicy_PreferOrd:
				if p.c.Quote.Edges.Ticker.Kind == schema.TickerKind_Stock {
					idx, ok := prefs[p.c.Quote.Edges.Ticker.Edges.Emitent]
					if ok {
						pieces2 = append(pieces2[:idx], pieces2[idx+1:]...)
					}
				} else if p.c.Quote.Edges.Ticker.Kind == schema.TickerKind_StockPref {
					_, ok := stocks[p.c.Quote.Edges.Ticker.Edges.Emitent]
					if ok {
						continue
					}
				}
			case domains.SameEmitentPolicy_AllowOnlyOne:
				{
					_, oks := stocks[p.c.Quote.Edges.Ticker.Edges.Emitent]
					_, okp := prefs[p.c.Quote.Edges.Ticker.Edges.Emitent]
					if oks || okp {
						continue
					}
				}
			}

			switch p.c.Quote.Edges.Ticker.Kind {
			case schema.TickerKind_Stock:
				stocks[p.c.Quote.Edges.Ticker.Edges.Emitent] = len(pieces2)
			case schema.TickerKind_StockPref:
				prefs[p.c.Quote.Edges.Ticker.Edges.Emitent] = len(pieces2)
			}

			pieces2 = append(pieces2, p)
			if len(pieces2) >= (Strategy.MaxTickers - fixedTickersCnt) {
				break
			}

		}

		//step 2 - needed tickers
		ws = 0.0
		for _, p := range pieces2 {
			ws += p.w
		}
		for idx, p := range pieces2 {
			pieces2[idx].sum = p.w / ws * flexRUB
		}

		for _, sp := range pieces2 {
			lotprice := float64(sp.c.LotSize()) * sp.c.Quote.C
			lots := int(math.Trunc(sp.sum / lotprice))
			deals := result.BuyLots(sp.c, lots)
			if len(deals) == 0 {
				log.Println("no deals for " + sp.c.TickerId())
			}
		}

	}

	// second, apply fixed tickers
	for _, r := range Strategy.Edges.FixedTickers {
		if !r.IsUsed {
			continue
		}
		cell := Market.CellsByTickerByDate(r.Ticker, D, true)
		if cell != nil {
			lots := int(math.Trunc((Source.CurrentBalance() + Source.RUB) / 100 * float64(r.Share) / (float64(cell.LotSize()) * cell.Quote.C)))
			if lots > 0 {
				result.BuyLots(cell, lots)
			}
		}
	}

	result.ApplyCurrentPrices(Market, Market.LastDate())

	sort.Slice(result.Items, func(i, j int) bool { return result.Items[i].CurrentPercent > result.Items[j].CurrentPercent })
	return &result

}

func Filter(Src *cube.Cell, Strategy *ent.Strategy, Market *cube.Cube, D time.Time) bool {

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

func Simulate(Strategy *ent.Strategy, Market *cube.Cube, From *time.Time, StartAmount float64, WeekRefillAmount float64, Implicit *ent.Ticker) *SimulationResult {

	if StartAmount < 0 {
		StartAmount = Strategy.StartAmount
	}
	if WeekRefillAmount < 0 {
		WeekRefillAmount = Strategy.WeekRefillAmount
	}

	result := &SimulationResult{}
	prtf := &Portfolio{RUB: StartAmount}

	AllTime_Equity := StartAmount
	AllTime_Dividends := 0.0

	today := time.Now()
	for _, D := range Market.GetAllDates(From, &today) {

		Refill := 0.0
		if D.Weekday() == time.Thursday {
			Refill = WeekRefillAmount
		}
		AllTime_Equity += Refill
		Divs := 0.0
		for _, item := range prtf.Items {
			c := Market.CellsByTickerByDate(item.Ticker.ID, D, false)
			if c != nil {
				continue
			}
			Divs += c.DivPayout
			result.TickerDividendResults = append(result.TickerDividendResults, &SimulationTickerDividendResultItem{
				Ticker:    c.Quote.Edges.Ticker,
				D:         D,
				Dividends: c.DivPayout * float64(item.Position),
			})
		}
		AllTime_Dividends += Divs

		prtf.RUB += Refill + Divs

		SD := &SimulationDay{D: D, Dividends: Divs, Refill: Refill}

		prtf.ApplyCurrentPrices(Market, D)

		var Ideal = ActualPortfolio(Strategy, Market, D, prtf, Implicit)
		Ideal.ApplyCurrentPrices(Market, D)

		// close missing positions
		for _, c := range prtf.Items {

			var found *PortfolioItem
			for _, idealItem := range Ideal.Items {
				if idealItem.Ticker.ID == c.Ticker.ID {
					found = idealItem
					break
				}
			}
			if found == nil {

				if Strategy.AllowLossWhenSell || c.balPrice < c.CurrentPrice {
					qc := Market.CellsByTickerByDate(c.Ticker.ID, D, true)
					if qc == nil {
						log.Printf("Simulation fail. Close position. No quote for %s for date %s\n", c.Ticker.ID, D)
					} else {
						deals := prtf.SellLots(qc, c.Position/c.LotSize)
						SD.Deals = append(SD.Deals, deals...)
					}
				}
			} else {
				if Strategy.AllowSellToFit {
					if found.Lots() > c.Lots() {
						if Strategy.AllowLossWhenSell || c.balPrice > found.balPrice {
							qc := Market.CellsByTickerByDate(c.Ticker.ID, D, true)
							if qc == nil {
								log.Printf("Simulation fail. Sell position. No quote for %s for date %s\n", c.Ticker.ID, D)
							} else {
								deals := prtf.SellLots(qc, found.Lots()-c.Lots())
								SD.Deals = append(SD.Deals, deals...)
							}
						}
					}
				}
			}
		}

		// open/update positions

		for _, idealItem := range Ideal.Items {

			var found *PortfolioItem
			for _, c := range prtf.Items {
				if c.Ticker.ID == idealItem.Ticker.ID {
					found = c
					break
				}
			}
			if found == nil {
				qc := Market.CellsByTickerByDate(idealItem.Ticker.ID, D, true)
				if qc == nil {
					log.Printf("Simulation fail. Open position. No quote for %s for date %s\n", idealItem.Ticker.ID, D)
				} else {
					deals := prtf.BuyLots(qc, idealItem.Position/idealItem.LotSize)
					SD.Deals = append(SD.Deals, deals...)
				}
			} else {
				if idealItem.Lots() > found.Lots() {
					qc := Market.CellsByTickerByDate(idealItem.Ticker.ID, D, true)
					if qc == nil {
						log.Printf("Simulation fail. No quote for %s for date %s\n", idealItem.Ticker.ID, D)
					} else {
						deals := prtf.BuyLots(qc, idealItem.Lots()-found.Lots())
						SD.Deals = append(SD.Deals, deals...)
					}
				}
			}
		}

		prtf.ApplyCurrentPrices(Market, D)
		SD.PortfolioRUB = prtf.RUB
		SD.PortfolioBalance = prtf.CurrentBalance()

		SD.Accu_InvestResult = SD.PortfolioBalance + SD.PortfolioRUB - AllTime_Equity - AllTime_Dividends
		SD.Accu_Equity = AllTime_Equity
		SD.Accu_Dividends = AllTime_Dividends

		result.Days = append(result.Days, SD)
	}

	result.Dates = make([]string, 0)
	result.Equity = make([]float64, 0)
	result.InvestResults = make([]float64, 0)
	result.Divs = make([]float64, 0)
	result.StrategyLevels = make([]float64, 0)

	for _, d := range result.Days {
		if d.D.Weekday() != time.Thursday {
			continue
		}
		result.Dates = append(result.Dates, d.D.Format("2006-01-02"))
		result.Equity = append(result.Equity, d.Accu_Equity)
		result.InvestResults = append(result.InvestResults, d.Accu_InvestResult)
		result.Divs = append(result.Divs, d.Accu_Dividends)
		result.StrategyLevels = append(result.StrategyLevels, d.PortfolioBalance+d.PortfolioRUB)

	}

	result.BaseLevels = make([]float64, len(result.StrategyLevels))
	if Strategy.BaseIndex != "" {
		BaseStrategy := &ent.Strategy{
			StartAmount:      Strategy.StartAmount,
			WeekRefillAmount: Strategy.WeekRefillAmount,
			StartSimulation:  Strategy.StartSimulation,
		}

		BaseResult := Simulate(
			BaseStrategy,
			Market,
			(*time.Time)(Strategy.StartSimulation),
			BaseStrategy.StartAmount,
			BaseStrategy.WeekRefillAmount,
			Market.GetAllTickers()[Strategy.BaseIndex],
		)

		idx := 0
		for _, d := range result.Days {
			if d.D.Weekday() != time.Thursday {
				continue
			}
			result.BaseLevels[idx] = BaseResult.Days[idx].PortfolioBalance + BaseResult.Days[idx].PortfolioRUB
			idx++
		}

	}

	result.Calc(Market, prtf)

	EmptyPortfolio := &Portfolio{RUB: Strategy.StartAmount}
	result.ActualPortfolio = ActualPortfolio(Strategy, Market, Market.LastDate(), EmptyPortfolio, nil)
	result.ActualPortfolio.ApplyCurrentPrices(Market, Market.LastDate())

	return result

}

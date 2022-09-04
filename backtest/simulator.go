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

func ActualPortfolio(Strategy *ent.Strategy, Market *cube.Cube, D time.Time, Source *Portfolio, ImplicitTicker *ent.Ticker, debug bool) *Portfolio {

	if ImplicitTicker != nil {
		ir := Portfolio{RUB: Source.CurrentBalance() + Source.RUB}
		cell := Market.CellsByTickerByDate(ImplicitTicker.ID, D, cube.LookBack)
		if cell == nil {
			log.Fatalf("ActualPortfolio: no cell for %s on %s\n", ImplicitTicker.Descr, D)
			return nil
		}
		lots := int(math.Trunc(Source.CurrentBalance() + Source.RUB/(float64(cell.LotSize())+cell.Quote.C)))
		ir.BuyLots(cell, lots)
		ir.ApplyCurrentPrices(Market, D)
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

		debugs := make(map[string]map[int]float64) // tickerId - factor.LineNum - factor weight

		buffer := make(map[*cube.Cell]map[int]float64, 0)

		for _, c := range all {
			if c.R2 == nil {
				continue
			}
			if k := c.Quote.Edges.Ticker.Kind; !(k == schema.TickerKind_Stock || k == schema.TickerKind_StockPref) {
				continue
			}
			if !Filter(c, Strategy, D) {
				continue
			}

			pfItem := Source.FindByTickerId(c.TickerId())

			for _, factor := range Strategy.Edges.Factors {
				if !factor.IsUsed {
					continue
				}
				rv := c.RepValue(factor.RK, factor.RVT)
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

				if debug {
					tm, ok := debugs[c.TickerId()]
					if !ok {
						tm = make(map[int]float64)
						debugs[c.TickerId()] = tm
					}
					tm[factor.LineNum] = r2
				}

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
		if ws == 0.0 {
			ws = 0.000000001
		}
		for idx, p := range pieces2 {
			pieces2[idx].sum = p.w / ws * flexRUB
		}

		for _, sp := range pieces2 {
			lotprice := float64(sp.c.LotSize()) * sp.c.Quote.C
			lots := int(math.Trunc(sp.sum / lotprice))
			if lots == 0 {
				continue
			}
			deals := result.BuyLots(sp.c, lots)
			if len(deals) == 0 {
				log.Println("ActualPortfolio. No deals for " + sp.c.TickerId())
			}
		}

		if debug {
			for _, pi := range result.Items {
				pi.DebugFactors = make(map[int]float64)
				for _, fi := range Strategy.Edges.Factors {
					if tm, ok := debugs[pi.Ticker.ID]; ok {
						pi.DebugFactors[fi.LineNum] = tm[fi.LineNum]
					}
				}
			}
		}

	}

	// second, apply fixed tickers
	for _, r := range Strategy.Edges.FixedTickers {
		if !r.IsUsed {
			continue
		}
		cell := Market.CellsByTickerByDate(r.Ticker, D, cube.LookBack)
		if cell != nil {
			lots := int(math.Trunc((Source.CurrentBalance() + Source.RUB) / 100 * float64(r.Share) / (float64(cell.LotSize()) * cell.Quote.C)))
			if lots > 0 {
				result.BuyLots(cell, lots)
			}
		}
	}

	result.ApplyCurrentPrices(Market, D)

	sort.Slice(result.Items, func(i, j int) bool { return result.Items[i].CurrentPercent > result.Items[j].CurrentPercent })
	return &result

}

func Filter(s0 *cube.Cell, Strategy *ent.Strategy, D time.Time) bool {

	for _, f := range Strategy.Edges.Filters {

		if !f.IsUsed {
			continue
		}

		switch f.LeftValueKind {
		case domains.FVK_Ticker:
			{
				if f.Operation == domains.FilterOp_Eq && s0.TickerId() != f.RightValueStr {
					return false
				}
				if f.Operation == domains.FilterOp_Ne && s0.TickerId() == f.RightValueStr {
					return false
				}
			}
		case domains.FVK_Industry:
			{
				if f.Operation == domains.FilterOp_Eq && s0.Industry.ID != f.RightValueStr {
					return false
				}
				if f.Operation == domains.FilterOp_Ne && s0.Industry.ID == f.RightValueStr {
					return false
				}
			}
		case domains.FVK_ReportValue:
			{
				lval := s0.RepValue(f.LeftReportValue, f.LeftReportValueType)
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
	result.TickerDividendResults = make([]*SimulationTickerDividendResultItem, 0)
	prtf := &Portfolio{RUB: StartAmount}

	AllTime_Equity := StartAmount
	AllTime_Dividends := 0.0

	today := time.Now()

	for _, D := range Market.GetAllDates(From, &today) {

		if D.Weekday() != time.Thursday {
			continue
		}
		Refill := WeekRefillAmount
		AllTime_Equity += Refill

		divsSum := 0.0
		for _, item := range prtf.Items {

			D0 := D
			D6 := D.Add(time.Hour * 24 * 7).Truncate(0)

			for D0.Before(D6) {

				c := Market.CellsByTickerByDate(item.Ticker.ID, D0, cube.LookBack)
				if c == nil {
					log.Fatalf("Simulation.divs. Unable to bind divpayout from date=%v for %s\n", D0, item.Ticker.ID)
				}
				if c.DivPayout > 0 {
					divsSum += c.DivPayout * float64(item.Position)

					var dtr *SimulationTickerDividendResultItem = nil
					for _, dr := range result.TickerDividendResults {
						if dr.Ticker.ID == c.TickerId() {
							dtr = dr
							break
						}
					}
					if dtr == nil {
						dtr = &SimulationTickerDividendResultItem{Ticker: c.Quote.Edges.Ticker, ByYears: make(map[int]float64)}
						result.TickerDividendResults = append(result.TickerDividendResults, dtr)
					}
					if _, ok := dtr.ByYears[D0.Year()]; !ok {
						dtr.ByYears[D0.Year()] = 0
					}
					dtr.ByYears[D0.Year()] += c.DivPayout * float64(item.Position)
				}

				D0 = D0.Add(time.Hour * 24).Truncate(0)
			}
		}
		AllTime_Dividends += divsSum

		prtf.RUB += Refill + divsSum

		SD := &SimulationDay{D: D, Refill: Refill}

		prtf.ApplyCurrentPrices(Market, D)

		var Ideal = ActualPortfolio(Strategy, Market, D, prtf, Implicit, false)
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
					qc := Market.CellsByTickerByDate(c.Ticker.ID, D, cube.LookBack)
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
							qc := Market.CellsByTickerByDate(c.Ticker.ID, D, cube.LookBack)
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
				qc := Market.CellsByTickerByDate(idealItem.Ticker.ID, D, cube.LookBack)
				if qc == nil {
					log.Printf("Simulation fail. Open position. No quote for %s for date %s\n", idealItem.Ticker.ID, D)
				} else {
					deals := prtf.BuyLots(qc, idealItem.Position/idealItem.LotSize)
					SD.Deals = append(SD.Deals, deals...)
				}
			} else {
				if idealItem.Lots() > found.Lots() {
					qc := Market.CellsByTickerByDate(idealItem.Ticker.ID, D, cube.LookBack)
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
	if Implicit == nil && Strategy.BaseIndex != "" {

		bResult := Simulate(
			Strategy,
			Market,
			(*time.Time)(Strategy.StartSimulation),
			Strategy.StartAmount,
			Strategy.WeekRefillAmount,
			Market.GetAllTickers()[Strategy.BaseIndex],
		)

		idx := 0
		for _, d := range bResult.Days {
			if d.D.Weekday() != time.Thursday {
				continue
			}
			result.BaseLevels[idx] = bResult.Days[idx].PortfolioBalance + bResult.Days[idx].PortfolioRUB
			idx++
		}

	}

	result.Calc(Market, prtf, true)

	EmptyPortfolio := &Portfolio{RUB: Strategy.StartAmount}
	result.ActualPortfolio = ActualPortfolio(Strategy, Market, Market.LastDate(), EmptyPortfolio, nil, true)

	return result

}

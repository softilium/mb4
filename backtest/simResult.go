package backtest

import (
	"fmt"
	"sort"
	"time"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/ent"
)

type SimulationDay struct {
	D                 time.Time
	Refill            float64
	PortfolioRUB      float64
	PortfolioBalance  float64
	Accu_InvestResult float64
	Accu_Yield        float64
	Accu_Dividends    float64
	Accu_Equity       float64
	Deals             []*Deal
}

type SimulationTickerTradeResultItem struct {
	TickerId   string
	Profit     float64
	DealsCount int
}

type SimulationTickerDividendResultItem struct {
	Ticker    *ent.Ticker
	D         time.Time
	Dividends float64
}

type SimulationIndustryTradeResultItem struct {
	IndustryId   string
	IndustryDesc string
	Profit       float64
	DealsCount   int
}

type SimulationIndustryDividendResultItem struct {
	Industry  *ent.Industry
	Year      int
	Dividends float64
}

type SimulationResult struct {
	// simulation result containers
	Days                  []*SimulationDay
	TickerDividendResults []*SimulationTickerDividendResultItem

	// simulation result statistics
	TickerTradeResults      []*SimulationTickerTradeResultItem
	IndustryTradeResults    []*SimulationIndustryTradeResultItem
	IndustryDividendResults []*SimulationIndustryDividendResultItem

	//render info
	ActualPortfolio *Portfolio
	Dates           []string
	Equity          []float64
	InvestResults   []float64
	Divs            []float64
	StrategyLevels  []float64
	BaseLevels      []float64
	YearsInResult   []int
	DivIndustries   []string
	ProfitTotal     float64
	DivsTotal       float64
}

func (t *SimulationResult) Calc(Market *cube.Cube, FinalPortfolio *Portfolio) {

	min := t.Days[0].D
	for _, d := range t.Days {
		if d.D.Before(min) {
			min = d.D
		}
	}

	ibt := make(map[string]*ent.Industry) // industry by tickerId
	indd := make(map[string]string)       // industry desc by id
	for _, c := range Market.GetCellsByDate(Market.LastDate()) {
		if c.Industry != nil {
			ibt[c.TickerId()] = c.Industry
			indd[c.Industry.ID] = c.Industry.Descr
		}
	}

	t.TickerTradeResults = make([]*SimulationTickerTradeResultItem, 0)

	yearsMap := make(map[int]bool)

	// profit for completed deals
	for _, D := range t.Days {

		yearsMap[D.D.Year()] = true

		for _, Deal := range D.Deals {
			ttrIdx := -1
			for idx, tc := range t.TickerTradeResults {
				if tc.TickerId == Deal.TickerId {
					ttrIdx = idx
					break
				}
			}
			if ttrIdx == -1 {
				t.TickerTradeResults = append(t.TickerTradeResults, &SimulationTickerTradeResultItem{
					TickerId:   Deal.TickerId,
					DealsCount: 1,
					Profit:     Deal.InvestResult,
				})
			} else {
				t.TickerTradeResults[ttrIdx].DealsCount++
				t.TickerTradeResults[ttrIdx].Profit += Deal.InvestResult
			}
		}
		ti_days := D.D.Sub(min).Hours() / 24
		D.Accu_Yield = 0
		if D.Accu_Equity != 0 && ti_days > 0 {
			D.Accu_Yield = D.Accu_InvestResult / D.Accu_Equity / ti_days * 365 * 100
		}

	}

	t.YearsInResult = make([]int, len(yearsMap))
	q := 0
	for y := range yearsMap {
		t.YearsInResult[q] = y
		q++
	}
	sort.Slice(t.YearsInResult, func(i, j int) bool { return t.YearsInResult[i] < t.YearsInResult[j] })

	// profit for opened positions (today)
	for _, i := range FinalPortfolio.Items {

		ttrIdx := 0
		for idx, tc := range t.TickerTradeResults {
			if tc.TickerId == i.Ticker.ID {
				ttrIdx = idx
				break
			}
		}

		c := Market.CellsByTickerByDate(i.Ticker.ID, Market.LastDate(), cube.LookBack)

		prof := float64(i.Position) * (c.Quote.C - i.BalPrice())

		if ttrIdx == -1 {
			t.TickerTradeResults = append(t.TickerTradeResults, &SimulationTickerTradeResultItem{
				TickerId:   i.Ticker.ID,
				DealsCount: 1,
				Profit:     prof,
			})
		} else {
			t.TickerTradeResults[ttrIdx].DealsCount++
			t.TickerTradeResults[ttrIdx].Profit += prof
		}

	}

	for _, R := range t.TickerTradeResults {
		for _, item := range FinalPortfolio.Items {
			if item.Ticker.ID == R.TickerId {
				cell := Market.CellsByTickerByDate(item.Ticker.ID, Market.LastDate(), cube.LookBack)
				for _, q := range item.Rests {
					R.Profit += (cell.Quote.C - q.Price) * float64(q.Position)
				}
				break
			}
		}
	}
	sort.Slice(t.TickerTradeResults, func(i, j int) bool { return t.TickerTradeResults[i].Profit > t.TickerTradeResults[j].Profit })

	// profit by industries
	indProfit := make(map[string]*struct {
		profit     float64
		dealsCount int
	})
	for _, r := range t.TickerTradeResults {

		indu, ok := ibt[r.TickerId]
		if !ok {
			continue
		}

		indRec, ok := indProfit[indu.ID]
		if !ok {
			indRec = &struct {
				profit     float64
				dealsCount int
			}{profit: 0, dealsCount: 0}
			indProfit[indu.ID] = indRec
		}
		indRec.profit += r.Profit
		indRec.dealsCount += r.DealsCount
	}
	t.IndustryTradeResults = make([]*SimulationIndustryTradeResultItem, len(indProfit))
	tidx := 0
	for k, ind := range indProfit {
		t.IndustryTradeResults[tidx] = &SimulationIndustryTradeResultItem{
			IndustryId:   k,
			IndustryDesc: indd[k],
			Profit:       ind.profit,
			DealsCount:   ind.dealsCount,
		}
		tidx++
	}
	sort.Slice(t.IndustryTradeResults, func(i, j int) bool {
		return t.IndustryTradeResults[i].Profit > t.IndustryTradeResults[j].Profit
	})

	// dividends by industries (Year + Industry)
	divMap := make(map[string]*SimulationIndustryDividendResultItem)
	for _, R := range t.TickerDividendResults {
		if R.Dividends <= 0 {
			continue
		}
		key := fmt.Sprintf("%v--%v", R.Ticker.Edges.Emitent.Edges.Industry.ID, R.D.Year())
		rec, ok := divMap[key]
		if !ok {
			rec = &SimulationIndustryDividendResultItem{Industry: R.Ticker.Edges.Emitent.Edges.Industry, Year: R.D.Year()}
			divMap[key] = rec
		}
		rec.Dividends += R.Dividends
	}
	t.IndustryDividendResults = make([]*SimulationIndustryDividendResultItem, 0, len(divMap))
	for _, v := range divMap {
		t.IndustryDividendResults = append(t.IndustryDividendResults, v)
	}
	sort.Slice(t.IndustryDividendResults, func(i, j int) bool {
		return t.IndustryDividendResults[i].Dividends > t.IndustryDividendResults[j].Dividends
	})

	q = len(t.Dates) - 1
	t.ProfitTotal = t.InvestResults[q]
	t.DivsTotal = t.Divs[q]

}

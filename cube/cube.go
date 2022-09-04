package cube

import (
	"context"

	"log"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/emission"
	"github.com/softilium/mb4/ent/schema"
	"github.com/softilium/mb4/ent/ticker"
)

func RoundX(x float64, dec int) float64 {
	mul := math.Pow10(dec)
	return math.Round(x*mul) / mul
}

func Avg(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

type Cube struct {
	l                   *sync.Mutex
	allDays             []time.Time // sorted
	allTickets          map[string]*ent.Ticker
	prefTickers         map[xid.ID]*ent.Ticker         // prefered tickers by emitents
	cellsByTickerByDate map[string]map[time.Time]*Cell // cell by ticker
	cellsByDate         map[time.Time][]*Cell          // cell by date
	repsByEmitent       map[xid.ID][]*Report2          // map by Emitent.ID sorted by reportdate

	cellsByIndustryByDate map[string]map[time.Time]*Cell // cell by industry
}

func (c *Cube) LoadCube() (err error) {

	c.l.Lock()
	defer c.l.Unlock()

	c.cellsByTickerByDate = make(map[string]map[time.Time]*Cell)
	c.cellsByDate = make(map[time.Time][]*Cell)
	c.allTickets = make(map[string]*ent.Ticker)
	allDaysMap := make(map[time.Time]bool)
	c.repsByEmitent = make(map[xid.ID][]*Report2)

	// for quick book value calc in reports we need emitent-pref map
	c.prefTickers = make(map[xid.ID]*ent.Ticker)
	prefraw, err := db.DB.Ticker.Query().WithEmitent().
		Where(ticker.KindEQ(schema.TickerKind_StockPref)).
		All(context.Background())
	if err != nil {
		return err
	}
	for _, pref := range prefraw {
		c.prefTickers[pref.Edges.Emitent.ID] = pref
	}

	q, err := db.DB.Quote.Query().
		WithTicker(
			func(q *ent.TickerQuery) {
				q.
					WithEmitent(
						func(q *ent.EmitentQuery) {
							q.WithIndustry()
						}).
					WithEmissions(
						func(q *ent.EmissionQuery) {
							q.Order(ent.Desc(emission.FieldRecDate))
						})

			}).
		All(context.Background())
	if err != nil {
		return err
	}

	for _, v := range q {
		allDaysMap[v.D] = true

		c.allTickets[v.Edges.Ticker.ID] = v.Edges.Ticker

		tdm, ok := c.cellsByTickerByDate[v.Edges.Ticker.ID]
		if !ok {
			tdm = make(map[time.Time]*Cell)
			c.cellsByTickerByDate[v.Edges.Ticker.ID] = tdm
		}

		oneCell := &Cell{Quote: v, D: v.D}
		oneCell.Industry = v.Edges.Ticker.Edges.Emitent.Edges.Industry

		tdm[v.D] = oneCell

		dm, ok := c.cellsByDate[v.D]
		if !ok {
			dm := make([]*Cell, 0)
			c.cellsByDate[v.D] = dm
		}
		dm = append(dm, oneCell)
		c.cellsByDate[v.D] = dm
	}

	c.allDays = make([]time.Time, len(allDaysMap))
	i := 0
	for k := range allDaysMap {
		c.allDays[i] = k
		i++
	}
	sort.Slice(c.allDays, func(i, j int) bool { return c.allDays[i].Before(c.allDays[j]) })

	if err = c.addMissingCells(); err != nil {
		return err
	}

	if err = c.linkEmissions(); err != nil {
		return err
	}

	if len(c.allDays) == 0 {
		return nil
	}

	if err = c.loadDivsAndCaps(); err != nil {
		return err
	}

	err = c.loadReports()
	if err != nil {
		return err
	}

	err = c.loadIndustries()
	if err != nil {
		return err
	}

	c.calcIndustryUpsides()

	//TODO считать положение котировки на 52-недельном цикле

	return nil

}

func (c *Cube) linkEmissions() error {

	em, err := db.DB.Emission.Query().WithTicker().All(context.Background())
	if err != nil {
		return err
	}

	type emplus struct { // data range (RecDate .. endDate)
		ent.Emission
		endDate *time.Time
	}

	byTicker := make(map[string][]*emplus)
	for _, v := range em {
		newemplus := &emplus{Emission: *v}
		if _, ok := byTicker[v.Edges.Ticker.ID]; !ok {
			ts := []*emplus{newemplus}
			byTicker[v.Edges.Ticker.ID] = ts
		} else {
			byTicker[v.Edges.Ticker.ID] = append(byTicker[v.Edges.Ticker.ID], newemplus)
		}
	}
	for _, v := range byTicker {
		sort.Slice(v, func(i, j int) bool {
			return v[i].RecDate.After(v[j].RecDate)
		})
		for i := len(v) - 1; i > 0; i-- {
			x := v[i-1].RecDate.AddDate(0, 0, -1)
			v[i].endDate = &x
		}
	}

	for tKey := range c.allTickets {
		em := byTicker[tKey]
		for _, v := range c.cellsByTickerByDate[tKey] {
			for i := 0; i < len(em); i++ {
				if v.D.Unix() >= em[i].RecDate.Unix() && (em[i].endDate == nil || v.D.Unix() <= em[i].endDate.Unix()) {
					v.emission = &em[i].Emission
					break
				}
			}
		}
	}

	// ensure we have all stocks and prefs are linked
	for tickerket, tickermap := range c.cellsByTickerByDate {
		for _, cell := range tickermap {
			to := c.allTickets[tickerket]
			if cell.emission == nil && (to.Kind == schema.TickerKind_Stock || to.Kind == schema.TickerKind_StockPref) {
				log.Panicf("Ticker %v, quote for %v has no emission info\n", tickerket, cell.D)
			}
		}
	}

	return nil

}

func (c *Cube) loadDivsAndCaps() error {

	dpRaw, err := db.DB.DivPayout.Query().WithTickers().All(context.Background())
	if err != nil {
		return err
	}

	for _, v := range dpRaw {
		cell := c._cellsByTickerByDate(v.CloseDate, v.Edges.Tickers.ID, LookAhead)
		if cell == nil {
			log.Println("No cell for div payout", v.CloseDate, v.Edges.Tickers.ID)
			_ = c._cellsByTickerByDate(v.CloseDate, v.Edges.Tickers.ID, LookAhead)
		} else {
			cell.DivPayout = v.DPS
		}
	}

	divPayoutMap := make(map[string]map[int]float64) // tickerID -> year -> divpayout
	for _, v := range dpRaw {
		tmap, ok := divPayoutMap[v.Edges.Tickers.ID]
		if !ok {
			tmap = make(map[int]float64)
			divPayoutMap[v.Edges.Tickers.ID] = tmap
		}
		if _, ok = tmap[v.ForYear]; !ok {
			tmap[v.ForYear] = v.DPS
		} else {
			tmap[v.ForYear] += v.DPS
		}
	}

	type dsiRecord struct {
		Stability byte
		Grow      byte
		dsi       float64
	}
	dsimap := make(map[string]map[int]dsiRecord)
	minYear := c.allDays[0].Year() + 6
	maxYear := c.allDays[len(c.allDays)-1].Year()
	for year := minYear; year <= maxYear; year++ {
		for ticker := range c.allTickets {
			cur := 0.0
			stability := 0.0
			grow := 0.0
			for i := 0; i < 6; i++ {
				actual := divPayoutMap[ticker][year-6+i]
				if actual > 0 && cur > 0 {
					stability++
				}
				if actual > 0 && cur > actual {
					grow += 1
				} else if actual > 0 && cur >= (actual*0.93) {
					grow += 0.5
				}
				cur = actual
			}
			if (stability + grow) > 0 {
				dsi := (stability + grow) / 14.0
				if _, ok := dsimap[ticker]; !ok {
					dsimap[ticker] = make(map[int]dsiRecord)
				}
				dsimap[ticker][year] = dsiRecord{Stability: byte(stability), Grow: byte(grow), dsi: RoundX(dsi, 1)}
			}
		}
	}

	for ticker, tickerv := range c.cellsByTickerByDate {
		dft, ok := divPayoutMap[ticker]
		if !ok {
			continue
		}

		lastYear, lastDivSum3, lastDivSum5 := -1, -1.0, -1.0
		minYear5, minYear3, maxYear := 0, 0, 0
		for _, day := range c.allDays {
			if _, ok := tickerv[day]; !ok {
				continue
			}
			if day.Year() != lastYear {
				lastYear, lastDivSum5, lastDivSum3 = day.Year(), 0.0, 0.0
				minYear5, minYear3, maxYear = lastYear-5, lastYear-3, lastYear-1
				for y, v := range dft {
					if y >= minYear5 && y <= maxYear {
						lastDivSum5 += v
					}
					if y >= minYear3 && y <= maxYear {
						lastDivSum3 += v
					}
				}
			}
			if lastDivSum5 < 0.0001 {
				continue
			}

			cl := c.cellsByTickerByDate[ticker][day]

			cl.DivSum5Y.S = lastDivSum5
			cl.DivSum3Y.S = lastDivSum3
			cl.DivYield5Y.S = RoundX(lastDivSum5/cl.Quote.C*100, 1)
			cl.DivYield3Y.S = RoundX(lastDivSum3/cl.Quote.C*100, 1)
		} //day
	} //ticker

	for _, day := range c.allDays {
		for _, cell := range c.cellsByDate[day] {
			if cell.Quote != nil && cell.emission != nil {
				cell.Cap.S = (cell.Quote.C * float64(cell.emission.Size)) / 1000000 // in mln. according to report values
			}
			if _, ok := dsimap[cell.TickerId()]; ok {
				if dsi, ok := dsimap[cell.TickerId()][day.Year()-1]; ok {
					cell.DSI.S = RoundX(dsi.dsi, 1)
				}
			}
		}
	}

	return nil
}

func (c *Cube) addMissingCells() error {

	// if quotes for ticker has holes in dates, fill them with last known cell

	lk := make(map[string]*Cell)
	for _, day := range c.allDays {
		for _, ticker := range c.allTickets {
			_, ok := c.cellsByTickerByDate[ticker.ID][day]
			if !ok && lk[ticker.ID] != nil {
				sc := lk[ticker.ID]

				newCell := &Cell{D: day, Quote: sc.Quote, emission: sc.emission, R2: sc.R2, IsMissed: true}

				c.cellsByDate[day] = append(c.cellsByDate[day], newCell)
				c.cellsByTickerByDate[ticker.ID][day] = sc
			}
			lk[ticker.ID] = c.cellsByTickerByDate[ticker.ID][day]
		}
	}
	return nil
}

func (c *Cube) loadReports() error {

	rd, err := db.DB.Report.Query().WithEmitent().All(context.Background())
	if err != nil {
		return err
	}

	for _, ticker := range c.allTickets {

		treps := make([]*ent.Report, 0, 10)
		for _, r := range rd {
			if r.Edges.Emitent.ID == ticker.Edges.Emitent.ID {
				treps = append(treps, r)
			}
		}
		//sort desc by reportdate
		sort.Slice(treps, func(i, j int) bool { return treps[i].ReportDate.Before(treps[j].ReportDate) })

		// convert reports to reports2
		r2reports := make([]*Report2, 0, 10)
		prevMaps := make(map[int]map[int]*Report2) // Year - Quarter - Report
		for _, r := range treps {

			r2 := Report2{}

			var prevY *Report2 = nil
			if _, ok := prevMaps[r.ReportYear-1]; ok {
				prevY = prevMaps[r.ReportYear-1][4]
			}
			var prevQ *Report2 = nil
			if _, ok := prevMaps[r.ReportYear-1]; ok {
				prevQ = prevMaps[r.ReportYear-1][r.ReportQuarter]
			}
			r2.LoadFromRawReport(r, prevY, prevQ)

			ymap, ok := prevMaps[r.ReportYear]
			if !ok {
				ymap = make(map[int]*Report2)
				prevMaps[r.ReportYear] = ymap
			}
			ymap[r.ReportQuarter] = &r2

			// bind report dates to quote dates for cube purposes.
			// Because we fild quotes by report dates later (see tickers page)
			for i := len(c.allDays) - 1; i >= 0; i-- {
				if c.allDays[i].Before(r2.ReportDate) {
					r2.ReportDate = c.allDays[i]
					break
				}
			}

			r2reports = append(r2reports, &r2)
		}

		//sort cells by dates
		c.repsByEmitent[ticker.Edges.Emitent.ID] = r2reports
		cells := c.cellsByTickerByDate[ticker.ID]
		dates := make([]time.Time, len(cells))
		i := 0
		for k := range cells {
			dates[i] = k
			i++
		}
		sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j]) })

		for _, D := range dates {
			cell := cells[D]
			for i := len(r2reports) - 1; i >= 0; i-- {
				if D.Unix() >= r2reports[i].ReportDate.Unix() {

					y := time.Date(D.Year()-1, D.Month(), D.Day(), D.Hour(), D.Minute(), D.Second(), D.Nanosecond(), D.Location())
					pyCell := c._cellsByTickerByDate(y, ticker.ID, LookBack)

					cell.R2 = r2reports[i]

					cell.CalcR3(c, pyCell)
					break
				}
			}
		}

	}

	return nil

}

func (c *Cube) loadIndustries() error {

	c.cellsByIndustryByDate = make(map[string]map[time.Time]*Cell) // slice of dsi by industry for date

	for _, day := range c.allDays {

		dsiArr := make(map[string][]float64)
		repsArr := make(map[string]*Cell)

		for _, cell := range c.cellsByDate[day] {
			if cell.Industry == nil || cell.R2 == nil {
				continue
			}
			indMap, ok := c.cellsByIndustryByDate[cell.Industry.ID]
			if !ok {
				indMap = make(map[time.Time]*Cell)
				c.cellsByIndustryByDate[cell.Industry.ID] = indMap
			}
			ir, ok := indMap[day]
			if !ok {
				ir = &Cell{D: day, Industry: cell.Industry}
				indMap[day] = ir
				ir.R2 = &Report2{ReportQuarter: cell.R2.ReportQuarter, ReportYear: cell.R2.ReportYear}
				ir.R2.Init()
			}
			repsArr[cell.Industry.ID] = ir

			ir.Cap.S += cell.Cap.S
			ir.DivSum3Y.S += cell.DivSum3Y.S
			ir.DivSum5Y.S += cell.DivSum5Y.S

			ir.R2.Revenue.YtdAdj += cell.R2.Revenue.YtdAdj
			ir.R2.Revenue.Ltm += cell.R2.Revenue.Ltm

			ir.R2.Amortization.YtdAdj += cell.R2.Amortization.YtdAdj
			ir.R2.Amortization.Ltm += cell.R2.Amortization.Ltm

			ir.R2.OperatingIncome.YtdAdj += cell.R2.OperatingIncome.YtdAdj
			ir.R2.OperatingIncome.Ltm += cell.R2.OperatingIncome.Ltm

			ir.R2.InterestIncome.YtdAdj += cell.R2.InterestIncome.YtdAdj
			ir.R2.InterestIncome.Ltm += cell.R2.InterestIncome.Ltm

			ir.R2.InterestExpenses.YtdAdj += cell.R2.InterestExpenses.YtdAdj
			ir.R2.InterestExpenses.Ltm += cell.R2.InterestExpenses.Ltm

			ir.R2.IncomeTax.YtdAdj += cell.R2.IncomeTax.YtdAdj
			ir.R2.IncomeTax.Ltm += cell.R2.IncomeTax.Ltm

			ir.R2.NetIncome.YtdAdj += cell.R2.NetIncome.YtdAdj
			ir.R2.NetIncome.Ltm += cell.R2.NetIncome.Ltm

			ir.R2.EBITDA.YtdAdj += cell.R2.EBITDA.YtdAdj
			ir.R2.EBITDA.Ltm += cell.R2.EBITDA.Ltm

			ir.R2.OIBDA.YtdAdj += cell.R2.OIBDA.YtdAdj
			ir.R2.OIBDA.Ltm += cell.R2.OIBDA.Ltm

			ir.R2.Cash.S += cell.R2.Cash.S
			ir.R2.NonCurrentLiabilities.S += cell.R2.NonCurrentLiabilities.S
			ir.R2.CurrentLiabilities.S += cell.R2.CurrentLiabilities.S
			ir.R2.NonControlling.S += cell.R2.NonControlling.S
			ir.R2.Equity.S += cell.R2.Equity.S
			ir.R2.Total.S += cell.R2.Total.S

			ir.R2.NetDebt.S += cell.R2.NetDebt.S

			_, ok = dsiArr[cell.Industry.ID]
			if !ok {
				dsiArr[cell.Industry.ID] = make([]float64, 0)
			}
			cell.IndustryCell = ir
			dsiArr[cell.Industry.ID] = append(dsiArr[cell.Industry.ID], cell.DSI.S)

		}
		for _, ir := range repsArr {
			ir.R2.CalcMults()
			ir.CalcR3(c, nil)

			dsiSlice, ok := dsiArr[ir.Industry.ID]
			if ok {
				ir.DSI.S = Avg(dsiSlice) // averate DSI for industry
			}
		}

	}

	return nil

}

func (c *Cube) getSortedCellsForTicker(ticker string) []*Cell {

	qm := c.cellsByTickerByDate[ticker]
	quotes := make([]*Cell, len(qm))
	idx := 0
	for _, v := range qm {
		quotes[idx] = v
		idx++
	}
	sort.Slice(quotes, func(i, j int) bool { return quotes[i].D.Before(quotes[j].D) })
	return quotes
}

func (c *Cube) calcIndustryUpsides() {

	for _, d := range c.allDays {
		for _, cell := range c.cellsByDate[d] {
			if cell.IndustryCell == nil {
				continue
			}

			cell.EV_on_EBITDA.IndustryUpside_YtdAdj = -Growth(cell.EV_on_EBITDA.YtdAdj, cell.IndustryCell.EV_on_EBITDA.YtdAdj, 1)
			cell.EV_on_EBITDA.IndustryUpside_Ltm = -Growth(cell.EV_on_EBITDA.Ltm, cell.IndustryCell.EV_on_EBITDA.Ltm, 1)

			cell.R2.Debt_on_EBITDA.IndustryUpside_YtdAdj = -Growth(cell.R2.Debt_on_EBITDA.YtdAdj, cell.IndustryCell.R2.Debt_on_EBITDA.YtdAdj, 1)
			cell.R2.Debt_on_EBITDA.IndustryUpside_Ltm = -Growth(cell.R2.Debt_on_EBITDA.Ltm, cell.IndustryCell.R2.Debt_on_EBITDA.Ltm, 1)

			cell.P_on_E.IndustryUpside_YtdAdj = -Growth(cell.P_on_E.YtdAdj, cell.IndustryCell.P_on_E.YtdAdj, 1)
			cell.P_on_E.IndustryUpside_Ltm = -Growth(cell.P_on_E.Ltm, cell.IndustryCell.P_on_E.Ltm, 1)

			cell.P_on_BV.IndustryUpside = -Growth(cell.P_on_BV.S, cell.IndustryCell.P_on_BV.S, 1)

			cell.P_on_S.IndustryUpside_YtdAdj = -Growth(cell.P_on_S.YtdAdj, cell.IndustryCell.P_on_S.YtdAdj, 1)
			cell.P_on_S.IndustryUpside_Ltm = -Growth(cell.P_on_S.Ltm, cell.IndustryCell.P_on_S.Ltm, 1)

			cell.R2.ROE.IndustryUpside_YtdAdj = -Growth(cell.R2.ROE.YtdAdj, cell.IndustryCell.R2.ROE.YtdAdj, 1)
			cell.R2.ROE.IndustryUpside_Ltm = -Growth(cell.R2.ROE.Ltm, cell.IndustryCell.R2.ROE.Ltm, 1)

			{ // margins

				cell.R2.OIBDAMargin.IndustryUpside_YtdAdj = -Growth(cell.R2.OIBDAMargin.YtdAdj, cell.IndustryCell.R2.OIBDAMargin.YtdAdj, 1)
				cell.R2.OIBDAMargin.IndustryUpside_Ltm = -Growth(cell.R2.OIBDAMargin.Ltm, cell.IndustryCell.R2.OIBDAMargin.Ltm, 1)

				cell.R2.EBITDAMargin.IndustryUpside_YtdAdj = -Growth(cell.R2.EBITDAMargin.YtdAdj, cell.IndustryCell.R2.EBITDAMargin.YtdAdj, 1)
				cell.R2.EBITDAMargin.IndustryUpside_Ltm = -Growth(cell.R2.EBITDAMargin.Ltm, cell.IndustryCell.R2.EBITDAMargin.Ltm, 1)

				cell.R2.OperationalMargin.IndustryUpside_YtdAdj = -Growth(cell.R2.OperationalMargin.YtdAdj, cell.IndustryCell.R2.OperationalMargin.YtdAdj, 1)
				cell.R2.OperationalMargin.IndustryUpside_Ltm = -Growth(cell.R2.OperationalMargin.Ltm, cell.IndustryCell.R2.OperationalMargin.Ltm, 1)

				cell.R2.NetMargin.IndustryUpside_YtdAdj = -Growth(cell.R2.NetMargin.YtdAdj, cell.IndustryCell.R2.NetMargin.YtdAdj, 1)
				cell.R2.NetMargin.IndustryUpside_Ltm = -Growth(cell.R2.NetMargin.Ltm, cell.IndustryCell.R2.NetMargin.Ltm, 1)

			}

			{ // dividends
				cell.DivYield3Y.IndustryUpside = Growth(cell.DivYield3Y.S, cell.IndustryCell.DivYield3Y.S, 1)
				cell.DivYield5Y.IndustryUpside = Growth(cell.DivYield5Y.S, cell.IndustryCell.DivYield5Y.S, 1)

				cell.DSI.IndustryUpside = Growth(cell.DSI.S, cell.IndustryCell.DSI.S, 1)
			}

		}
	}

}

var Market *Cube = &Cube{l: &sync.Mutex{}}

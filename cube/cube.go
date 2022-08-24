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
	"github.com/softilium/mb4/domains"
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

type Cell struct {
	D                       time.Time
	Quote                   *ent.Quote //nil means industry card for day
	emission                *ent.Emission
	emission_lotsize_cached int
	R2                      *Report2      //same report for all cells between published IFRS reports
	Industry                *ent.Industry // flat industry from quote
	IsMissed                bool          //indicates than cell was copied for missing quotes from prevous days
	DivPayout               float64       //div payout for day

	//R3
	BookValue  RepV
	P_on_E     RepV
	P_on_BV    RepV
	Cap        RepV
	P_on_S     RepV
	DivSum5Y   RepV
	DivSum3Y   RepV
	DivYield5Y RepV
	DivYield3Y RepV
	DSI        RepV
}

func (r *Cell) LotSize() int {
	if r.emission != nil {
		ls := r.emission.LotSize
		if ls == 0 {

			// emissions can contains 0 in lotsSize. We scan for different non-zero lots sizes.
			// If we found 1 non-zero value, we will use it in

			if r.emission_lotsize_cached != 0 {
				return r.emission_lotsize_cached
			}
			minls := 1000000000
			maxls := 0
			for _, rec := range r.Quote.Edges.Ticker.Edges.Emissions {
				if rec.LotSize < minls && rec.LotSize > 0 {
					minls = rec.LotSize
				}
				if rec.LotSize > maxls {
					maxls = rec.LotSize
				}
			}
			if minls == maxls {
				r.emission_lotsize_cached = minls
				return minls
			} else {
				log.Fatalf("LotSize is not uniform for ticker %s", r.Quote.Edges.Ticker.ID)
				return 0
			}
		}
		return ls
	}
	return 1

}

func (r *Cell) TickerId() string {
	return r.Quote.Edges.Ticker.ID
}

func (r *Cell) CalcAfterLoad(cb *Cube) {

	r.BookValue.V = 0
	prefCap := 0.0
	if prefTicker, ok := cb.prefTickers[r.Quote.Edges.Ticker.Edges.Emitent.ID]; ok {
		if prefCells, ok := cb.cellsByTickerByDate[prefTicker.ID]; ok {
			if prefcell, ok := prefCells[r.D]; ok {
				prefCap = prefcell.Cap.V
			}
		}
	}
	r.BookValue.V = r.R2.Total.V - r.R2.CurrentLiabilities.V - r.R2.NonCurrentLiabilities.V - r.R2.NonControlling.V - prefCap

	if r.BookValue.V != 0 {
		r.P_on_BV.V = r.Cap.V / r.BookValue.V
	}

	if r.R2.NetIncome.V != 0 {
		r.P_on_E.V = r.Cap.V / r.R2.NetIncome.V
		r.P_on_E.Ltm = r.Cap.V / r.R2.NetIncome.Ltm
	}

	r.P_on_S.V = r.Cap.V / r.R2.Revenue.V
	r.P_on_S.Ltm = r.Cap.V / r.R2.Revenue.Ltm

}

func (c *Cell) RepValue(market *Cube, rv domains.ReportValue, rvt domains.ReportValueType) float64 {

	var r2 *Report2
	var r3 *Cell

	switch rvt {
	case domains.RVT_Src, domains.RVT_YtdAdj, domains.RVT_Ltm, domains.RVT_AG, domains.RVT_AG_Ltm:
		r2 = c.R2
		r3 = c
	case domains.RVT_Ind_Src, domains.RVT_Ind_YtdAdj, domains.RVT_Ind_Ltm, domains.RVT_Ind_AG, domains.RVT_Ind_AG_Ltm:
		r3 = market.cellsByIndustryByDate[c.Industry.ID][c.D]
		r2 = r3.R2
	}

	var rc RepV

	switch rv {
	case domains.RK_Revenue:
		rc = r2.Revenue
	case domains.RK_Amortization:
		rc = r2.Amortization
	case domains.RK_OperatingIncome:
		rc = r2.OperatingIncome
	case domains.RK_InterestIncome:
		rc = r2.InterestIncome
	case domains.RK_InterestExpenses:
		rc = r2.InterestExpenses
	case domains.RK_IncomeTax:
		rc = r2.IncomeTax
	case domains.RK_NetIncome:
		rc = r2.NetIncome
	case domains.RK_OIBDA:
		rc = r2.OIBDA
	case domains.RK_EBITDA:
		rc = r2.EBITDA
	case domains.RK_OIBDAMargin:
		rc = r2.OIBDAMargin
	case domains.RK_EBITDAMargin:
		rc = r2.EBITDAMargin
	case domains.RK_OperationalMargin:
		rc = r2.OperationalMargin
	case domains.RK_NetMargin:
		rc = r2.NetMargin
	case domains.RK_Debt_on_EBITDA:
		rc = r2.Debt_on_EBITDA
	case domains.RK_EV_on_EBITDA:
		rc = r2.EV_on_EBITDA
	case domains.RK_ROE:
		rc = r2.ROE
	case domains.RK_Cash:
		rc = r2.Cash
	case domains.RK_NonCurrentLiabilities:
		rc = r2.NonCurrentLiabilities
	case domains.RK_CurrentLiabilities:
		rc = r2.CurrentLiabilities
	case domains.RK_NonControlling:
		rc = r2.NonControlling
	case domains.RK_Equity:
		rc = r2.Equity
	case domains.RK_Total:
		rc = r2.Total
	case domains.RK_NetDebt:
		rc = r2.NetDebt
	case domains.RK_EV:
		rc = r2.EV
	case domains.RK_BookValue:
		rc = r3.BookValue
	case domains.RK_P_on_E:
		rc = r3.P_on_E
	case domains.RK_P_on_BV:
		rc = r3.P_on_BV
	case domains.RK_Cap:
		rc = r3.Cap
	case domains.RK_P_on_S:
		rc = r3.P_on_S
	case domains.RK_DivSum5Y:
		rc = r3.DivSum5Y
	case domains.RK_DivSum3Y:
		rc = r3.DivSum3Y
	case domains.RK_DivYield5Y:
		rc = r3.DivYield5Y
	case domains.RK_DivYield3Y:
		rc = r3.DivYield3Y
	case domains.RK_DSI:
		rc = r3.DSI

	default:
		log.Fatalf("Unable to get report value for %v", rv)
	}

	switch rvt {
	case domains.RVT_Src, domains.RVT_Ind_Src:
		return rc.V
	case domains.RVT_YtdAdj, domains.RVT_Ind_YtdAdj:
		return rc.YtdAdj
	case domains.RVT_Ltm, domains.RVT_Ind_Ltm:
		return rc.Ltm
	case domains.RVT_AG, domains.RVT_Ind_AG:
		return rc.AG
	case domains.RVT_AG_Ltm, domains.RVT_Ind_AG_Ltm:
		return rc.AGLtm
	default:
		log.Fatalf("Unable to get report value for %v (%v)", rv, rvt)
		return 0.0
	}

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

	c.loadReports()

	c.loadIndustries()

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
			//TODO 4 strings there !
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

			cl.DivSum5Y.V = lastDivSum5
			cl.DivSum3Y.V = lastDivSum3
			cl.DivYield5Y.V = RoundX(lastDivSum5/cl.Quote.C*100, 1)
			cl.DivYield3Y.V = RoundX(lastDivSum3/cl.Quote.C*100, 1)
		} //day
	} //ticker

	for _, day := range c.allDays {
		for _, cell := range c.cellsByDate[day] {
			if cell.Quote != nil && cell.emission != nil {
				cell.Cap.V = cell.Quote.C * float64(cell.emission.Size) / 1000000 // in mln. according to report values
			}
			if _, ok := dsimap[cell.TickerId()]; ok {
				if dsi, ok := dsimap[cell.TickerId()][day.Year()-1]; ok {
					cell.DSI.V = RoundX(dsi.dsi, 1)
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

		c.repsByEmitent[ticker.Edges.Emitent.ID] = r2reports
		for D, cell := range c.cellsByTickerByDate[ticker.ID] { // make report view for each day/quote
			for i := len(r2reports) - 1; i >= 0; i-- {
				if D.Unix() >= r2reports[i].ReportDate.Unix() {
					cell.R2 = r2reports[i]
					cell.CalcAfterLoad(c)
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
			}
			repsArr[cell.Industry.ID] = ir

			ir.Cap.V += cell.Cap.V
			ir.DivSum3Y.V += cell.DivSum3Y.V
			ir.DivSum5Y.V += cell.DivSum5Y.V

			ir.R2 = &Report2{ReportQuarter: 4, ReportYear: cell.R2.ReportYear}
			ir.R2.Init()

			ir.R2.Revenue.V += cell.R2.Revenue.V
			ir.R2.Revenue.Ltm += cell.R2.Revenue.Ltm

			ir.R2.Amortization.V += cell.R2.Amortization.V
			ir.R2.Amortization.Ltm += cell.R2.Amortization.Ltm

			ir.R2.OperatingIncome.V += cell.R2.OperatingIncome.V
			ir.R2.OperatingIncome.Ltm += cell.R2.OperatingIncome.Ltm

			ir.R2.InterestIncome.V += cell.R2.InterestIncome.V
			ir.R2.InterestIncome.Ltm += cell.R2.InterestIncome.Ltm

			ir.R2.InterestExpenses.V += cell.R2.InterestExpenses.V
			ir.R2.InterestExpenses.Ltm += cell.R2.InterestExpenses.Ltm

			ir.R2.IncomeTax.V += cell.R2.IncomeTax.V
			ir.R2.IncomeTax.Ltm += cell.R2.IncomeTax.Ltm

			ir.R2.NetIncome.V += cell.R2.NetIncome.V
			ir.R2.NetIncome.Ltm += cell.R2.NetIncome.Ltm

			ir.R2.Cash.V += cell.R2.Cash.V
			ir.R2.NonCurrentLiabilities.V += cell.R2.NonCurrentLiabilities.V
			ir.R2.CurrentLiabilities.V += cell.R2.CurrentLiabilities.V
			ir.R2.NonControlling.V += cell.R2.NonControlling.V
			ir.R2.Equity.V += cell.R2.Equity.V
			ir.R2.Total.V += cell.R2.Total.V

			ir.R2.Calc(nil, nil)

			//TODO calc previous reports
			//TODO calc avg.values from mults for industry+day. Now it is value-weighted average, we need to just flat average

			_, ok = dsiArr[cell.Industry.ID]
			if !ok {
				dsiArr[cell.Industry.ID] = make([]float64, 0)
			}
			dsiArr[cell.Industry.ID] = append(dsiArr[cell.Industry.ID], cell.DSI.V)

		}
		for _, ir := range repsArr {
			dsiSlice, ok := dsiArr[ir.Industry.ID]
			if ok {
				ir.DSI.V = Avg(dsiSlice)
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

var Market *Cube = &Cube{l: &sync.Mutex{}}

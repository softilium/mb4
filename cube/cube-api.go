package cube

import (
	"sort"
	"time"

	"github.com/softilium/mb4/ent"
)

func (c *Cube) TopDivYields5Y(HowMany int) []*Cell {
	c.l.Lock()
	defer c.l.Unlock()

	if len(c.allDays) == 0 {
		return make([]*Cell, 0)
	}

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DivYield5Y.S > newslice[j].DivYield5Y.S })
	return newslice[:HowMany]
}

func (c *Cube) TopDSI(HowMany int) []*Cell {
	c.l.Lock()
	defer c.l.Unlock()

	if len(c.allDays) == 0 {
		return make([]*Cell, 0)
	}

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DSI.S > newslice[j].DSI.S })
	return newslice[:HowMany]
}

type ItemPriceChange struct {
	Cell               *Cell
	PercentPriceChange float64
}

func (c *Cube) TopFallenRaise(HowMany int, Raise bool) []ItemPriceChange {

	c.l.Lock()
	defer c.l.Unlock()

	if len(c.allDays) == 0 {
		return make([]ItemPriceChange, 0)
	}

	lastDate := c.allDays[len(c.allDays)-1]
	oneMonthAgo := lastDate.Add(-30 * 24 * time.Hour)
	for i := len(c.allDays) - 1; i >= 0; i-- {
		if c.allDays[i].Before(oneMonthAgo) {
			oneMonthAgo = c.allDays[i]
			break
		}
	}
	sliceNow := c.cellsByDate[lastDate]

	changes := make([]ItemPriceChange, len(sliceNow))
	for idx, v := range sliceNow {
		prevPrice := c.cellsByTickerByDate[v.TickerId()][oneMonthAgo].Quote.C
		changes[idx] = ItemPriceChange{v, RoundX((v.Quote.C-prevPrice)/prevPrice*100, 1)}
	}
	if !Raise {
		sort.Slice(changes, func(i, j int) bool { return changes[i].PercentPriceChange < changes[j].PercentPriceChange })
	} else {
		sort.Slice(changes, func(i, j int) bool { return changes[i].PercentPriceChange > changes[j].PercentPriceChange })
	}
	return changes[:HowMany]

}

type TickerRenderInfo struct {
	Cell          *Cell
	CandleDates   []string
	CandleOCLH    [][4]float64
	CandleVolumes []float64
}

func (c *Cube) GetTickerRenderInfo(ticker string, loadQ bool) *TickerRenderInfo {

	c.l.Lock()
	defer c.l.Unlock()

	result := &TickerRenderInfo{}
	result.Cell = c.cellsByTickerByDate[ticker][c.allDays[len(c.allDays)-1]]

	if !loadQ {
		return result
	}

	quotes := c.getSortedCellsForTicker(ticker)
	result.CandleDates = make([]string, len(quotes))
	result.CandleOCLH = make([][4]float64, len(quotes))
	result.CandleVolumes = make([]float64, len(quotes))
	for idx, v := range quotes {
		result.CandleDates[idx] = v.D.Format("2006-01-02")
		result.CandleOCLH[idx] = [4]float64{v.Quote.O, v.Quote.C, v.Quote.L, v.Quote.H}
		result.CandleVolumes[idx] = v.Quote.V
	}

	return result

}

func (c *Cube) GetReports2(ticker string) []*Report2 {

	c.l.Lock()
	defer c.l.Unlock()

	tobj, ok := c.allTickets[ticker]
	if !ok {
		return []*Report2{}
	}

	return c.repsByEmitent[tobj.Edges.Emitent.ID]

}

type LookDirection int

const (
	LookBack  LookDirection = -1
	LookNone  LookDirection = 0
	LookAhead LookDirection = 1
)

func (c *Cube) CellsByTickerByDate(ticker string, d time.Time, lookDir LookDirection) *Cell {
	c.l.Lock()
	defer c.l.Unlock()
	return c._cellsByTickerByDate(d, ticker, lookDir)
}

func (c *Cube) _cellsByTickerByDate(d time.Time, ticker string, lookDir LookDirection) *Cell {
	d0 := d
	r, ok := c.cellsByTickerByDate[ticker][d0]
	if lookDir == LookBack {
		minDate := c.allDays[0]
		for !ok && d0.After(minDate) {
			d0 = d0.AddDate(0, 0, -1)
			r, ok = c.cellsByTickerByDate[ticker][d0]
		}
	}
	if lookDir == LookAhead {
		maxDate := c.allDays[len(c.allDays)-1]
		for !ok && d0.Before(maxDate) {
			d0 = d0.AddDate(0, 0, 1)
			r, ok = c.cellsByTickerByDate[ticker][d0]
		}
	}
	return r
}

func (c *Cube) GetAllTickers() map[string]*ent.Ticker {

	c.l.Lock()
	defer c.l.Unlock()

	return c.allTickets

}

func (c *Cube) GetAllDates(Min *time.Time, Max *time.Time) []time.Time {
	c.l.Lock()
	defer c.l.Unlock()

	res := make([]time.Time, 0)
	for _, v := range c.allDays {
		if Min != nil && v.Before(*Min) {
			continue
		}
		if Max != nil && v.After(*Max) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func (c *Cube) GetIndustryCell(industry string, d time.Time) *Cell {

	c.l.Lock()
	defer c.l.Unlock()

	return c.cellsByIndustryByDate[industry][d]

}

func (c *Cube) LastDate() time.Time {
	c.l.Lock()
	defer c.l.Unlock()

	return c.allDays[len(c.allDays)-1]
}

func (c *Cube) GetCellsByDate(D time.Time) []*Cell {

	c.l.Lock()
	defer c.l.Unlock()

	return c.cellsByDate[D]

}

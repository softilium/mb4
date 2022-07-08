package cube

import (
	"sort"
	"time"
)

func (c *Cube) TopDivYields5Y(HowMany int) []*Cell {
	c.l.Lock()
	defer c.l.Unlock()

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DivYield5Y > newslice[j].DivYield5Y })
	return newslice[:HowMany]
}

func (c *Cube) TopDSI(HowMany int) []*Cell {
	c.l.Lock()
	defer c.l.Unlock()

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DSI > newslice[j].DSI })
	return newslice[:HowMany]
}

type ItemPriceChange struct {
	Cell               *Cell
	PercentPriceChange float64
}

func (c *Cube) TopFallenRaise(HowMany int, Raise bool) []ItemPriceChange {

	c.l.Lock()
	defer c.l.Unlock()

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
		prevPrice := c.cellsByTickerByDate[v.Quote.Edges.Ticker.ID][oneMonthAgo].Quote.C
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

func (c *Cube) CellsByTickerByDate(ticker string, d time.Time) *Cell {

	c.l.Lock()
	defer c.l.Unlock()

	return c.cellsByTickerByDate[ticker][d]

}
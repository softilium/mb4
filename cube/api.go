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

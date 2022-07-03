package cube

import (
	"sort"
)

func (c *Cube) TopDivYields5Y(HowMany int) ([]*Cell, error) {

	c.l.Lock()
	defer c.l.Unlock()

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DivYield5Y > newslice[j].DivYield5Y })
	return newslice[:HowMany], nil

}

func (c *Cube) TopDSI(HowMany int) ([]*Cell, error) {

	c.l.Lock()
	defer c.l.Unlock()

	slice := c.cellsByDate[c.allDays[len(c.allDays)-1]]
	newslice := make([]*Cell, len(slice))
	copy(newslice, slice)
	sort.Slice(newslice, func(i, j int) bool { return newslice[i].DSI > newslice[j].DSI })
	return newslice[:HowMany], nil

}

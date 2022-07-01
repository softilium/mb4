package cube

import (
	"context"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/schema"
)

type Cell struct {
	D        time.Time // missed quotes propagades to the next day
	Quote    *ent.Quote
	Emission *ent.Emission
	Report   *ent.Report
	IsMissed bool //indicates than cell was copied for missing quotes
}

type Cube struct {
	l          *sync.Mutex
	allDays    []time.Time // sorted
	allTickets map[string]*ent.Ticker
	//allReports          map[*ent.Emitent][]*ent.Report // slice sorted by Year, Quarter
	cellsByTickerByDate map[string]map[time.Time]*Cell // cell by ticker
	cellsByDate         map[time.Time][]*Cell          // cell by date
}

func (c *Cube) LoadCube() (err error) {

	c.l.Lock()
	defer c.l.Unlock()

	c.cellsByTickerByDate = make(map[string]map[time.Time]*Cell)
	c.cellsByDate = make(map[time.Time][]*Cell)
	c.allTickets = make(map[string]*ent.Ticker)
	allDaysMap := make(map[time.Time]bool)

	q, err := db.DB.Quote.Query().WithTicker().
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
		tdm[v.D] = &Cell{Quote: v, D: v.D}

		dm, ok := c.cellsByDate[v.D]
		if !ok {
			dm := make([]*Cell, 0)
			c.cellsByDate[v.D] = dm
		}
		dm = append(dm, &Cell{Quote: v, D: v.D})
		c.cellsByDate[v.D] = dm
	}

	c.allDays = make([]time.Time, len(allDaysMap))
	i := 0
	for k := range allDaysMap {
		c.allDays[i] = k
		i++
	}
	sort.Slice(c.allDays, func(i, j int) bool { return c.allDays[i].Before(c.allDays[j]) })

	if err = c.linkEmissions(); err != nil {
		return err
	}

	if err = c.loadDivs(); err != nil {
		return err
	}

	if err = c.addMissingCells(); err != nil {
		return err
	}

	//статические поля отчета  в отчет++
	//Расставить отчеты++
	//досчитать динамику по дням в отчеты++++
	//досчитать рост
	//досчитать отраслевые отчеты++ и отчеты++++
	//досчитать отраслевой рост

	return nil

}

func (c *Cube) linkEmissions() error {

	em, err := db.DB.Emission.Query().WithTicker().All(context.Background())
	if err != nil {
		return err
	}

	type emplus struct {
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
					v.Emission = &em[i].Emission
					break
				}
			}
		}
	}

	// ensure we have all stocks and prefs are linked
	for tickerket, tickermap := range c.cellsByTickerByDate {
		for _, cell := range tickermap {
			to := c.allTickets[tickerket]
			if cell.Emission == nil && (to.Kind == schema.TickerKind_Stock || to.Kind == schema.TickerKind_StockPref) {
				log.Panicf("Ticker %v, quote for %v has no emission info\n", tickerket, cell.D)
			}
		}
	}

	return nil

}

func (c *Cube) loadDivs() error {
	//расчитать свои дивидендные коэффициенты, суммы и доходность по ячейкам
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
				c.cellsByDate[day] = append(c.cellsByDate[day],
					&Cell{D: day, Quote: sc.Quote, Emission: sc.Emission, Report: sc.Report, IsMissed: true})
				c.cellsByTickerByDate[ticker.ID][day] = sc
			}
			lk[ticker.ID] = c.cellsByTickerByDate[ticker.ID][day]
		}
	}
	return nil
}

var Market *Cube = &Cube{l: &sync.Mutex{}}

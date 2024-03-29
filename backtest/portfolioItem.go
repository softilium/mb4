package backtest

import (
	"encoding/json"
	"log"
	"sort"
	"time"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/ent"
)

type PortfolioItemRest struct {
	D        time.Time
	Position int
	Price    float64
	Cell     *cube.Cell
}

type PortfolioItem struct {
	Ticker         *ent.Ticker
	LotSize        int
	Rests          []*PortfolioItemRest
	Position       int
	balance        float64 //read-only
	balPrice       float64 //read-only
	CurrentPrice   float64
	CurrentPercent float64
	DebugFactors   map[int]float64
}

func (p *PortfolioItem) Balance() float64 {
	return p.balance
}

func (d *PortfolioItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Ticker         *ent.Ticker
		Lots           int
		Balance        float64
		CurrentPercent float64
		DebugFactors   map[int]float64
	}{
		Ticker:         d.Ticker,
		Lots:           d.Lots(),
		Balance:        d.Balance(),
		CurrentPercent: d.CurrentPercent,
		DebugFactors:   d.DebugFactors,
	})
}

func (p *PortfolioItem) BalPrice() float64 {
	return p.balPrice
}

func (p *PortfolioItem) Lots() int {
	return p.Position / p.LotSize
}

func (p *PortfolioItem) CurrentBalance() float64 {
	return p.CurrentPrice * float64(p.Position)
}

func (p *PortfolioItem) Refresh() {
	p.Position = 0
	p.balance = 0
	for _, r := range p.Rests {
		p.Position += r.Position
		p.balance += float64(r.Position) * r.Price
	}
	p.balPrice = p.balance / float64(p.Position)
}

func (p *PortfolioItem) Buy(c *cube.Cell, delta int) []*Deal {

	if c.LotSize() == 0 {
		log.Fatalf("Lot size is 0 for %s (%v)", p.Ticker.ID, c.LotSize())
		return nil
	}

	newd := Deal{
		D:            c.D,
		Kind:         Buy,
		Price:        c.Quote.C,
		TickerId:     p.Ticker.ID,
		Volume:       delta,
		Lots:         delta / c.LotSize(),
		InvestResult: 0,
	}

	p.Rests = append(p.Rests, &PortfolioItemRest{D: c.D, Position: delta, Price: c.Quote.C, Cell: c})
	p.Refresh()

	result := []*Deal{&newd}
	return result

}

func (p *PortfolioItem) Sell(Cell *cube.Cell, Delta int) []*Deal {

	if p.Position < Delta {
		return make([]*Deal, 0)
	}

	toProcess := Delta
	result := make([]*Deal, 0)
	sort.Slice(p.Rests, func(i, j int) bool { return p.Rests[i].D.Before(p.Rests[j].D) })

	for toProcess > 0 && len(p.Rests) > 0 {

		rest := p.Rests[0]

		piece := toProcess
		if piece > rest.Position {
			piece = rest.Position
		}

		rest.Position -= piece
		if rest.Position == 0 {
			p.Rests = p.Rests[1:]
		}
		toProcess -= piece

		if Cell.LotSize() == 0 {
			log.Fatalf("Lot size is 0 for %s (%v)", p.Ticker.ID, Cell.LotSize())
			return nil
		}

		newDial := Deal{
			D:            Cell.D,
			Kind:         Sell,
			TickerId:     p.Ticker.ID,
			Volume:       piece,
			Price:        Cell.Quote.C,
			Lots:         piece / Cell.LotSize(),
			InvestResult: (Cell.Quote.C - rest.Price) * float64(piece),
		}
		result = append(result, &newDial)
	}
	return result

}

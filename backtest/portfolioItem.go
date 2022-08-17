package backtest

import (
	"encoding/json"
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
	}{
		Ticker:         d.Ticker,
		Lots:           d.Lots(),
		Balance:        d.Balance(),
		CurrentPercent: d.CurrentPercent,
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
		p.balPrice = p.balance / float64(p.Position)
	}

}

func (p *PortfolioItem) Buy(Cell *cube.Cell, Delta int) []*Deal {

	newDial := Deal{
		D:            Cell.D,
		Kind:         Buy,
		Price:        Cell.Quote.C,
		TickerId:     p.Ticker.ID,
		Volume:       Delta,
		Lots:         Delta / Cell.Emission.LotSize,
		InvestResult: 0,
	}

	p.Rests = append(p.Rests, &PortfolioItemRest{D: Cell.D, Position: Delta, Price: Cell.Quote.C, Cell: Cell})
	p.Refresh()

	result := []*Deal{&newDial}
	return result

}

func (p *PortfolioItem) Sell(Cell *cube.Cell, Delta int) []*Deal {

	if p.Position < Delta {
		return make([]*Deal, 0)
	}
	toProcess := Delta
	result := make([]*Deal, 0)
	restsSorted := make([]*PortfolioItemRest, len(p.Rests))
	copy(restsSorted, p.Rests)
	sort.Slice(restsSorted, func(i, j int) bool { return restsSorted[i].D.Before(restsSorted[j].D) })

	for toProcess > 0 && len(p.Rests) > 0 {

		rest := restsSorted[0]

		piece := toProcess
		if piece > rest.Position {
			piece = rest.Position
		}

		rest.Position -= piece
		if rest.Position == 0 {
			p.Rests = p.Rests[1:]
		}
		toProcess -= piece

		newDial := Deal{
			D:            Cell.D,
			Kind:         Sell,
			TickerId:     p.Ticker.ID,
			Volume:       piece,
			Price:        Cell.Quote.C,
			Lots:         piece / Cell.Emission.LotSize,
			InvestResult: (Cell.Quote.C - rest.Price) * float64(piece),
		}
		result = append(result, &newDial)
	}
	return result

}

package backtest

import (
	"math"
	"time"

	"github.com/softilium/mb4/cube"
)

type Portfolio struct {
	RUB   float64
	Items []*PortfolioItem
}

func (p *Portfolio) CurrentBalance() float64 {
	res := 0.0
	for _, item := range p.Items {
		res += item.CurrentBalance()
	}
	return res
}

func (p *Portfolio) BuyLots(Cell *cube.Cell, LotsToBuy int) []*Deal {

	lots := LotsToBuy
	maxLots := int(math.Trunc(p.RUB / float64(Cell.Emission.LotSize*int(Cell.Quote.C))))
	if lots > maxLots {
		lots = maxLots
	}
	if lots == 0 {
		return make([]*Deal, 0)
	}

	tIdx := -1
	for idx, item := range p.Items {
		if item.Ticker.ID == Cell.Quote.Edges.Ticker.ID {
			tIdx = idx
			break
		}
	}
	if tIdx == -1 {
		p.Items = append(p.Items, &PortfolioItem{
			Ticker:  Cell.Quote.Edges.Ticker,
			LotSize: Cell.Emission.LotSize,
		})
		tIdx = len(p.Items) - 1
	}
	deals := p.Items[tIdx].Buy(Cell, lots*Cell.Emission.LotSize)
	for _, deal := range deals {
		p.RUB -= deal.Sum()
	}
	return deals

}

func (p *Portfolio) SellLots(Cell *cube.Cell, lots int) []*Deal {

	tIdx := -1
	for idx, item := range p.Items {
		if item.Ticker.ID == Cell.Quote.Edges.Ticker.ID {
			tIdx = idx
			break
		}
	}
	if tIdx == -1 {
		return make([]*Deal, 0)
	}

	item := p.Items[tIdx]
	lotsToSell := lots
	if lotsToSell > item.Lots() {
		lotsToSell = item.Lots()
	}

	deals := item.Sell(Cell, lotsToSell*item.LotSize)

	if len(item.Rests) == 0 {
		p.Items = append(p.Items[:tIdx], p.Items[tIdx+1:]...)
	}

	for _, deal := range deals {
		p.RUB += deal.Sum()
	}

	return deals

}

func (p *Portfolio) ApplyCurrentPrices(market *cube.Cube, D time.Time) {

	// step one: apply current prices to all items
	for _, item := range p.Items {
		cell := market.CellsByTickerByDate(item.Ticker.ID, D, true)
		item.CurrentPrice = 0
		if cell != nil {
			item.CurrentPrice = cell.Quote.C
		}
	}

	// step two: calc percent shares
	for _, item := range p.Items {
		item.CurrentPercent = item.CurrentBalance() / p.CurrentBalance() * 100
	}

}

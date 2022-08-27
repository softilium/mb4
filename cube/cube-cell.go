package cube

import (
	"log"
	"time"

	"github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
)

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
	EV           RepV
	EV_on_EBITDA RepV
	BookValue    RepV
	P_on_E       RepV
	P_on_BV      RepV
	Cap          RepV
	P_on_S       RepV
	DivSum5Y     RepV
	DivSum3Y     RepV
	DivYield5Y   RepV
	DivYield3Y   RepV
	DSI          RepV
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
	if r.Quote != nil {
		if prefTicker, ok := cb.prefTickers[r.Quote.Edges.Ticker.Edges.Emitent.ID]; ok {
			if prefCells, ok := cb.cellsByTickerByDate[prefTicker.ID]; ok {
				if prefcell, ok := prefCells[r.D]; ok {
					prefCap = prefcell.Cap.V
				}
			}
		}
	}

	r.EV.V = r.Cap.V + r.R2.Cash.V + r.R2.NonControlling.V + r.R2.NonCurrentLiabilities.V + r.R2.CurrentLiabilities.V
	r.EV.Ltm = r.EV.V
	r.EV.YtdAdj = r.EV.V

	r.EV_on_EBITDA.V = r.EV.V / r.R2.EBITDA.V
	r.EV_on_EBITDA.YtdAdj = r.EV.YtdAdj / r.R2.EBITDA.YtdAdj
	r.EV_on_EBITDA.Ltm = r.EV.Ltm / r.R2.EBITDA.Ltm

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

	r.EV_on_EBITDA.InverseGrowth = true
	r.P_on_E.InverseGrowth = true
	r.P_on_BV.InverseGrowth = true

}

func (c *Cell) GetRepV(k domains.ReportValue) *RepV {

	switch k {
	case domains.RK_Revenue:
		return &c.R2.Revenue
	case domains.RK_Amortization:
		return &c.R2.Amortization
	case domains.RK_OperatingIncome:
		return &c.R2.OperatingIncome
	case domains.RK_InterestIncome:
		return &c.R2.InterestIncome
	case domains.RK_InterestExpenses:
		return &c.R2.InterestExpenses
	case domains.RK_IncomeTax:
		return &c.R2.IncomeTax
	case domains.RK_NetIncome:
		return &c.R2.NetIncome
	case domains.RK_OIBDA:
		return &c.R2.OIBDA
	case domains.RK_EBITDA:
		return &c.R2.EBITDA
	case domains.RK_OIBDAMargin:
		return &c.R2.OIBDAMargin
	case domains.RK_EBITDAMargin:
		return &c.R2.EBITDAMargin
	case domains.RK_OperationalMargin:
		return &c.R2.OperationalMargin
	case domains.RK_NetMargin:
		return &c.R2.NetMargin
	case domains.RK_Debt_on_EBITDA:
		return &c.R2.Debt_on_EBITDA
	case domains.RK_ROE:
		return &c.R2.ROE
	case domains.RK_Cash:
		return &c.R2.Cash
	case domains.RK_NonCurrentLiabilities:
		return &c.R2.NonCurrentLiabilities
	case domains.RK_CurrentLiabilities:
		return &c.R2.CurrentLiabilities
	case domains.RK_NonControlling:
		return &c.R2.NonControlling
	case domains.RK_Equity:
		return &c.R2.Equity
	case domains.RK_Total:
		return &c.R2.Total
	case domains.RK_NetDebt:
		return &c.R2.NetDebt
	case domains.RK_EV:
		return &c.EV
	case domains.RK_EV_on_EBITDA:
		return &c.EV_on_EBITDA
	case domains.RK_BookValue:
		return &c.BookValue
	case domains.RK_P_on_E:
		return &c.P_on_E
	case domains.RK_P_on_BV:
		return &c.P_on_BV
	case domains.RK_Cap:
		return &c.Cap
	case domains.RK_P_on_S:
		return &c.P_on_S
	case domains.RK_DivSum5Y:
		return &c.DivSum5Y
	case domains.RK_DivSum3Y:
		return &c.DivSum3Y
	case domains.RK_DivYield5Y:
		return &c.DivYield5Y
	case domains.RK_DivYield3Y:
		return &c.DivYield3Y
	case domains.RK_DSI:
		return &c.DSI

	default:
		log.Fatalf("Unable to get report value for %v", k)
	}

	return nil

}

func (c *Cell) RepValue(market *Cube, rv domains.ReportValue, rvt domains.ReportValueType) float64 {

	var ind, rc *RepV
	rc = c.GetRepV(rv)
	switch rvt {
	case domains.RVT_Ind_Src, domains.RVT_Ind_YtdAdj, domains.RVT_Ind_Ltm, domains.RVT_Ind_AG, domains.RVT_Ind_AG_Ltm,
		domains.RVT_IndUpside_Src, domains.RVT_IndUpside_YtdAdj, domains.RVT_IndUpside_Ltm, domains.RVT_IndUpside_AG, domains.RVT_IndUpside_AG_Ltm:
		ind = market.cellsByIndustryByDate[c.Industry.ID][c.D].GetRepV(rv)
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
	case domains.RVT_IndUpside_Src:
		return rc.CalcIndUpside_V(ind)
	case domains.RVT_IndUpside_YtdAdj:
		return rc.CalcIndUpside_YtdAdj(ind)
	case domains.RVT_IndUpside_Ltm:
		return rc.CalcIndUpside_Ltm(ind)
	default:
		log.Fatalf("Unable to get report value for %v (%v)", rv, rvt)
		return 0.0
	}

}

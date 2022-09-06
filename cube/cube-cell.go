package cube

import (
	"log"
	"time"

	d "github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
)

type Cell struct {
	D                       time.Time
	Quote                   *ent.Quote //nil means industry card for day
	emission                *ent.Emission
	emission_lotsize_cached int
	R2                      *Report2      //same report for all cells between published IFRS reports
	Industry                *ent.Industry // flat industry from quote
	IndustryCell            *Cell         //linked industry-cell
	IsMissed                bool          //indicates than cell was copied for missing quotes from prevous days
	DivPayout               float64       //div payout for day

	//R3
	EV           RepS
	EV_on_EBITDA RepV
	BookValue    RepS
	P_on_E       RepV
	P_on_BV      RepS
	Cap          RepS
	P_on_S       RepV
	DivSum5Y     RepS
	DivSum3Y     RepS
	DivYield5Y   RepS
	DivYield3Y   RepS
	DSI          RepS
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
				log.Fatalf("LotSize==0 && is not uniform for ticker %s", r.Quote.Edges.Ticker.ID)
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

func (r *Cell) CalcR3(cb *Cube, py *Cell) {

	r.BookValue.S = 0
	prefCap := 0.0
	if r.Quote != nil {
		if prefTicker, ok := cb.prefTickers[r.Quote.Edges.Ticker.Edges.Emitent.ID]; ok {
			if prefCells, ok := cb.cellsByTickerByDate[prefTicker.ID]; ok {
				if prefcell, ok := prefCells[r.D]; ok {
					prefCap = prefcell.Cap.S
				}
			}
		}
	}

	r.EV.S = r.Cap.S + r.R2.Cash.S + r.R2.NonControlling.S + r.R2.NonCurrentLiabilities.S + r.R2.CurrentLiabilities.S

	r.EV_on_EBITDA.YtdAdj = IIF(r.R2.EBITDA.YtdAdj == 0, 0, r.EV.S/r.R2.EBITDA.YtdAdj)
	r.EV_on_EBITDA.Ltm = IIF(r.R2.EBITDA.Ltm == 0, 0, r.EV.S/r.R2.EBITDA.Ltm)

	r.BookValue.S = r.R2.Total.S - r.R2.CurrentLiabilities.S - r.R2.NonCurrentLiabilities.S - r.R2.NonControlling.S - prefCap

	r.P_on_BV.S = IIF(r.BookValue.S == 0, 0, r.Cap.S/r.BookValue.S)

	r.P_on_E.YtdAdj = IIF(r.R2.NetIncome.YtdAdj == 0, 0, r.Cap.S/r.R2.NetIncome.YtdAdj)
	r.P_on_E.Ltm = IIF(r.R2.NetIncome.Ltm == 0, 0, r.Cap.S/r.R2.NetIncome.Ltm)

	r.P_on_S.YtdAdj = IIF(r.R2.Revenue.YtdAdj == 0, 0, r.Cap.S/r.R2.Revenue.YtdAdj)
	r.P_on_S.Ltm = IIF(r.R2.Revenue.Ltm == 0, 0, r.Cap.S/r.R2.Revenue.Ltm)

	if py != nil {
		r.Cap.AG = Growth(r.Cap.S, py.Cap.S, 1)

		r.EV.AG = Growth(r.EV.S, py.EV.S, 1)

		r.EV_on_EBITDA.AGYtdAdj = Growth(r.EV_on_EBITDA.YtdAdj, py.EV_on_EBITDA.YtdAdj, 1)
		r.EV_on_EBITDA.AGLtm = Growth(r.EV_on_EBITDA.Ltm, py.EV_on_EBITDA.Ltm, 1)

		r.BookValue.AG = Growth(r.BookValue.S, py.BookValue.S, 1)

		r.P_on_BV.AG = Growth(r.P_on_BV.S, py.P_on_BV.S, 1)

		r.P_on_E.AGYtdAdj = Growth(r.P_on_E.YtdAdj, py.P_on_E.YtdAdj, 1)
		r.P_on_E.AGLtm = Growth(r.P_on_E.Ltm, py.P_on_E.Ltm, 1)

		r.P_on_S.AGYtdAdj = Growth(r.P_on_S.YtdAdj, py.P_on_S.YtdAdj, 1)
		r.P_on_S.Ltm = Growth(r.P_on_S.Ltm, py.P_on_S.Ltm, 1)

		r.DSI.AG = Growth(r.DSI.S, py.DSI.S, 1)

		r.DivSum5Y.AG = Growth(r.DivSum5Y.S, py.DivSum5Y.S, 1)
		r.DivSum3Y.AG = Growth(r.DivSum3Y.S, py.DivSum3Y.S, 1)
		r.DivYield5Y.AG = Growth(r.DivYield5Y.S, py.DivYield5Y.S, 1)
		r.DivYield3Y.AG = Growth(r.DivYield3Y.S, py.DivYield3Y.S, 1)
	}

}

func (c *Cell) GetRepV(k d.ReportValue) *RepV {

	switch k {
	case d.RK_Revenue:
		return &c.R2.Revenue
	case d.RK_Amortization:
		return &c.R2.Amortization
	case d.RK_OperatingIncome:
		return &c.R2.OperatingIncome
	case d.RK_InterestIncome:
		return &c.R2.InterestIncome
	case d.RK_InterestExpenses:
		return &c.R2.InterestExpenses
	case d.RK_IncomeTax:
		return &c.R2.IncomeTax
	case d.RK_NetIncome:
		return &c.R2.NetIncome
	case d.RK_OIBDA:
		return &c.R2.OIBDA
	case d.RK_EBITDA:
		return &c.R2.EBITDA
	case d.RK_OIBDAMargin:
		return &c.R2.OIBDAMargin
	case d.RK_EBITDAMargin:
		return &c.R2.EBITDAMargin
	case d.RK_OperationalMargin:
		return &c.R2.OperationalMargin
	case d.RK_NetMargin:
		return &c.R2.NetMargin
	case d.RK_Debt_on_EBITDA:
		return &c.R2.Debt_on_EBITDA
	case d.RK_ROE:
		return &c.R2.ROE
	case d.RK_EV_on_EBITDA:
		return &c.EV_on_EBITDA
	case d.RK_P_on_E:
		return &c.P_on_E
	case d.RK_P_on_S:
		return &c.P_on_S

	}

	return nil

}

func (c *Cell) GetRepS(k d.ReportValue) *RepS {

	switch k {
	case d.RK_Cash:
		return &c.R2.Cash
	case d.RK_NonCurrentLiabilities:
		return &c.R2.NonCurrentLiabilities
	case d.RK_CurrentLiabilities:
		return &c.R2.CurrentLiabilities
	case d.RK_NonControlling:
		return &c.R2.NonControlling
	case d.RK_Equity:
		return &c.R2.Equity
	case d.RK_Total:
		return &c.R2.Total
	case d.RK_NetDebt:
		return &c.R2.NetDebt
	case d.RK_EV:
		return &c.EV
	case d.RK_BookValue:
		return &c.BookValue
	case d.RK_P_on_BV:
		return &c.P_on_BV
	case d.RK_Cap:
		return &c.Cap
	case d.RK_DivSum5Y:
		return &c.DivSum5Y
	case d.RK_DivSum3Y:
		return &c.DivSum3Y
	case d.RK_DivYield5Y:
		return &c.DivYield5Y
	case d.RK_DivYield3Y:
		return &c.DivYield3Y
	case d.RK_DSI:
		return &c.DSI
	default:
		return nil
	}

}

func (c *Cell) RepValue(rv d.ReportValue, rvt d.ReportValueType) float64 {

	valV := c.GetRepV(rv)
	valS := c.GetRepS(rv)

	switch rvt {
	case d.RVT_Src:
		if valV == nil {
			return 0
		}
		return valV.Src
	case d.RVT_S:
		if valS == nil {
			return 0
		}
		return valS.S
	case d.RVT_YtdAdj:
		if valV == nil {
			return 0
		}
		return valV.YtdAdj
	case d.RVT_Ltm:
		if valV == nil {
			return 0
		}
		return valV.Ltm
	case d.RVT_AG:
		if valS == nil {
			return 0
		}
		return valS.AG
	case d.RVT_AG_YtdAdj:
		if valV == nil {
			return 0
		}
		return valV.AGYtdAdj
	case d.RVT_AG_Ltm:
		if valV == nil {
			return 0
		}
		return valV.AGLtm
	case d.RVT_IndUpside:
		if valS == nil {
			return 0
		}
		return valS.IndustryUpside
	case d.RVT_IndUpside_YtdAdj:
		if valV == nil {
			return 0
		}
		return valV.IndustryUpside_YtdAdj
	case d.RVT_IndUpside_Ltm:
		if valV == nil {
			return 0
		}
		return valV.IndustryUpside_Ltm
	default:
		log.Fatalf("Unable to call Cell.RepValue for rv=%v, rvt=%v", rv, rvt)
		return 0.0
	}

}

func (c *Cell) EvalChange(rv d.ReportValue, rvt d.ReportValueType) ChangeValulation {

	upGood := func(c float64) ChangeValulation {
		if c <= float64(VeryBad) {
			return VeryBad
		}
		if c < float64(Bad) {
			return Bad
		}
		if c > float64(VeryGood) {
			return VeryGood
		}
		if c > float64(Good) {
			return Good
		}
		return Neutral
	}

	negIsBad := func(c float64) ChangeValulation {
		if c < 0 {
			return VeryBad
		}
		return Neutral
	}

	pozIsBad := func(c float64) ChangeValulation {
		if c > 0 {
			return VeryBad
		}
		return Neutral
	}

	downGood := func(c float64) ChangeValulation { return -upGood(c) }
	neutral := func(c float64) ChangeValulation { return Neutral }

	switch rvt {
	case d.RVT_AG, d.RVT_AG_YtdAdj, d.RVT_AG_Ltm, d.RVT_IndUpside, d.RVT_IndUpside_Ltm, d.RVT_IndUpside_YtdAdj:
		switch rv {
		case d.RK_Revenue, d.RK_OperatingIncome, d.RK_NetIncome, d.RK_OIBDA,
			d.RK_EBITDA, d.RK_OIBDAMargin, d.RK_EBITDAMargin, d.RK_OperationalMargin,
			d.RK_NetMargin, d.RK_ROE, d.RK_Equity, d.RK_Total, d.RK_BookValue,
			d.RK_Cap, d.RK_DivYield5Y, d.RK_DivYield3Y, d.RK_DSI:
			return upGood(c.RepValue(rv, rvt))
		case d.RK_Debt_on_EBITDA, d.RK_EV_on_EBITDA, d.RK_NetDebt, d.RK_P_on_E:
			return downGood(c.RepValue(rv, rvt))
		default:
			return neutral(0)
		}

	case d.RVT_S, d.RVT_Src, d.RVT_Ltm, d.RVT_YtdAdj:
		switch rv {
		case d.RK_Revenue, d.RK_OperatingIncome, d.RK_InterestIncome, d.RK_NetIncome, d.RK_OIBDA,
			d.RK_EBITDA, d.RK_OIBDAMargin, d.RK_EBITDAMargin, d.RK_OperationalMargin,
			d.RK_NetMargin, d.RK_ROE, d.RK_Equity, d.RK_Total, d.RK_BookValue,
			d.RK_Cap, d.RK_DivYield5Y, d.RK_DivYield3Y, d.RK_Cash, d.RK_NonCurrentLiabilities, d.RK_CurrentLiabilities,
			d.RK_NonControlling, d.RK_EV, d.RK_DivSum5Y, d.RK_DivSum3Y:
			return negIsBad(c.RepValue(rv, rvt))
		case d.RK_Amortization, d.RK_InterestExpenses, d.RK_IncomeTax:
			return pozIsBad(c.RepValue(rv, rvt))
		default:
			return neutral(0)
		}
	}

	log.Panicf("Unable to determine report value for eval: %v\n", rv)
	return neutral(0)

}

// TODO move to report page code (UI)
func (c *Cell) EvalChangeAsClass(rv d.ReportValue, rvt d.ReportValueType) string {

	r := c.EvalChange(rv, rvt)

	switch r {
	case VeryBad:
		return "table-danger fw-bold"
	case Bad:
		return "table-danger"
	case VeryGood:
		return "table-success fw-bold"
	case Good:
		return "table-success"
	}

	return ""

}

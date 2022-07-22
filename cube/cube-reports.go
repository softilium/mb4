package cube

import (
	"math"
	"time"

	"github.com/softilium/mb4/ent"
)

const (
	//Pnl src
	RK2Revenue          = 100
	RK2Amortization     = 110
	RK2OperatingIncome  = 120
	RK2InterestIncome   = 130
	RK2InterestExpenses = 140
	RK2IncomeTax        = 150
	RK2NetIncome        = 160

	//Pnl calculated
	RK2OIBDA             = 200
	RK2EBITDA            = 210
	RK2OIBDAMargin       = 220
	RK2EBITDAMargin      = 230
	RK2OperationalMargin = 240
	RK2NetMargin         = 250
	RK2Debt_On_EBITDA    = 260
	RK2EV_On_EBITDA      = 270
	RK2ROE               = 280

	// CF src
	RK2Cash                  = 300
	RK2NonCurrentLiabilities = 305
	RK2CurrentLiabilities    = 310
	RK2NonControlling        = 315
	RK2Equity                = 320
	RK2Total                 = 325

	// CF calculated
	RK2NetDebt = 400
	RK2EV      = 410
)

type Pnl2Value struct {
	RK     int
	V      float64 //Ytd
	YtdAdj float64
	Ltm    float64
	AG     float64 // Annual Growth
	AGLtm  float64 // Annual Growth LTM
}

func (p *Pnl2Value) Calc(src, prevQ, prevY *Report2) {

	if p.V == 0 {
		return
	}

	p.YtdAdj = p.V / float64(src.ReportQuarter) * 4
	if p.Ltm == 0 { // skip when we assign it before
		if prevQ == nil || prevY == nil {
			p.Ltm = p.YtdAdj
		} else {
			p.Ltm = prevY.YV[p.RK].V - prevQ.YV[p.RK].V + p.V
		}
	}
	if prevY == nil {
		p.AG = 0
		p.AGLtm = 0
	} else {
		p.AG = RoundX(p.V/prevY.YV[p.RK].V*100, 1) - 100
		p.AGLtm = RoundX(p.Ltm/prevY.YV[p.RK].Ltm*100, 1) - 100
	}

}

type Cf2Value struct {
	RK int
	V  float64
	AG float64
}

func (p *Cf2Value) Calc(prevY *Report2) {

	if prevY == nil {
		p.AG = 0
	} else {
		p.AG = RoundX(p.V/prevY.SV[p.RK].V*100, 1) - 100
	}

}

type Report2 struct { // enriched report with calculated fields
	ReportQuarter int
	ReportDate    time.Time

	prevYear    *Report2
	prevQuarter *Report2

	YV map[int]*Pnl2Value // year-based values (PNL, ...)
	SV map[int]*Cf2Value  // saldo-based values (CF, ...)

}

func (r *Report2) Load(s *ent.Report, prevY, prevQ *Report2) {

	r.ReportQuarter = s.ReportQuarter
	r.ReportDate = s.ReportDate

	r.YV = make(map[int]*Pnl2Value, 15)
	r.SV = make(map[int]*Cf2Value, 6)

	r.prevQuarter = prevQ
	r.prevYear = prevY

	// Pnl src
	r.YV[RK2Revenue] = &Pnl2Value{RK: RK2Revenue, V: s.PnlRevenueYtd}
	r.YV[RK2Amortization] = &Pnl2Value{RK: RK2Amortization, V: s.PnlAmortizationYtd}
	r.YV[RK2OperatingIncome] = &Pnl2Value{RK: RK2OperatingIncome, V: s.PnlOperatingIncomeYtd}
	r.YV[RK2InterestIncome] = &Pnl2Value{RK: RK2InterestIncome, V: s.PnlInterestIncomeYtd}
	r.YV[RK2InterestExpenses] = &Pnl2Value{RK: RK2InterestExpenses, V: s.PnlInterestExpensesYtd}
	r.YV[RK2IncomeTax] = &Pnl2Value{RK: RK2IncomeTax, V: s.PnlIncomeTaxYtd}
	r.YV[RK2NetIncome] = &Pnl2Value{RK: RK2NetIncome, V: s.PnlNetIncomeYtd}
	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}

	// CF src
	r.SV[RK2Cash] = &Cf2Value{V: s.CfCashSld}
	r.SV[RK2NonCurrentLiabilities] = &Cf2Value{V: s.CfNonCurrentLiabilitiesSld}
	r.SV[RK2CurrentLiabilities] = &Cf2Value{V: s.CfCurrentLiabilitesSld}
	r.SV[RK2NonControlling] = &Cf2Value{V: s.CfNonControllingSld}
	r.SV[RK2Equity] = &Cf2Value{V: s.CfEquitySld}
	r.SV[RK2Total] = &Cf2Value{V: s.CfTotalSld}

	// CF calculated
	r.SV[RK2NetDebt] = &Cf2Value{
		RK: RK2NetDebt,
		V:  r.SV[RK2NonCurrentLiabilities].V + r.SV[RK2CurrentLiabilities].V - r.SV[RK2Cash].V,
	}
	r.SV[RK2EV] = &Cf2Value{
		RK: RK2EV,
		V:  r.SV[RK2Cash].V + r.SV[RK2NonControlling].V + r.SV[RK2NonCurrentLiabilities].V + r.SV[RK2CurrentLiabilities].V,
	}

	// Pnl calculated

	r.YV[RK2OIBDA] = &Pnl2Value{
		RK:  RK2OIBDA,
		V:   r.YV[RK2Revenue].V - r.YV[RK2Amortization].V,
		Ltm: r.YV[RK2Revenue].Ltm - r.YV[RK2Amortization].Ltm,
	}
	r.YV[RK2EBITDA] = &Pnl2Value{
		RK:  RK2EBITDA,
		V:   r.YV[RK2NetIncome].V - r.YV[RK2Amortization].V - r.YV[RK2InterestIncome].V - r.YV[RK2InterestExpenses].V - r.YV[RK2IncomeTax].V,
		Ltm: r.YV[RK2NetIncome].Ltm - r.YV[RK2Amortization].Ltm - r.YV[RK2InterestIncome].Ltm - r.YV[RK2InterestExpenses].Ltm - r.YV[RK2IncomeTax].Ltm,
	}

	r.YV[RK2OIBDAMargin] = &Pnl2Value{RK: RK2OIBDAMargin}
	r.YV[RK2EBITDAMargin] = &Pnl2Value{RK: RK2EBITDAMargin}
	r.YV[RK2OperationalMargin] = &Pnl2Value{RK: RK2OperationalMargin}
	r.YV[RK2NetMargin] = &Pnl2Value{RK: RK2NetMargin}
	if math.Abs(s.PnlRevenueYtd) >= 0.01 {

		r.YV[RK2OIBDAMargin].V = RoundX(r.YV[RK2OIBDA].V/r.YV[RK2Revenue].V*100, 1)
		r.YV[RK2OIBDAMargin].Ltm = RoundX(r.YV[RK2OIBDA].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2EBITDAMargin].V = RoundX(r.YV[RK2EBITDA].V/r.YV[RK2Revenue].V*100, 1)
		r.YV[RK2EBITDAMargin].Ltm = RoundX(r.YV[RK2EBITDA].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2OperationalMargin].V = RoundX(r.YV[RK2OperatingIncome].V/r.YV[RK2Revenue].V*100, 1)
		r.YV[RK2OperationalMargin].Ltm = RoundX(r.YV[RK2OperatingIncome].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2NetMargin].V = RoundX(r.YV[RK2NetIncome].V/r.YV[RK2Revenue].V*100, 1)
		r.YV[RK2NetMargin].Ltm = RoundX(r.YV[RK2NetIncome].Ltm/r.YV[RK2Revenue].Ltm*100, 1)
	}

	r.YV[RK2ROE] = &Pnl2Value{RK: RK2ROE}
	if math.Abs(r.SV[RK2Total].V) >= 0.01 {
		r.YV[RK2ROE].V = RoundX(r.YV[RK2NetIncome].V/r.SV[RK2Total].V*100, 1)
		r.YV[RK2ROE].Ltm = RoundX(r.YV[RK2NetIncome].Ltm/r.SV[RK2Total].V*100, 1)
	}
	r.YV[RK2Debt_On_EBITDA] = &Pnl2Value{RK: RK2Debt_On_EBITDA}
	r.YV[RK2EV_On_EBITDA] = &Pnl2Value{RK: RK2EV_On_EBITDA}
	if math.Abs(r.YV[RK2EBITDA].V) >= 0.01 {
		r.YV[RK2Debt_On_EBITDA].V = r.SV[RK2NetDebt].V / r.YV[RK2EBITDA].V
		r.YV[RK2Debt_On_EBITDA].Ltm = r.SV[RK2NetDebt].V / r.YV[RK2EBITDA].Ltm
		r.YV[RK2EV_On_EBITDA].V = r.SV[RK2EV].V / r.YV[RK2EBITDA].V
		r.YV[RK2EV_On_EBITDA].Ltm = r.SV[RK2EV].V / r.YV[RK2EBITDA].Ltm
	}

	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}
	for _, v := range r.SV {
		v.Calc(r.prevYear)
	}

}

type Cf3Value struct {
	RK int
	V  float64
	AG float64 // anual growth
}

const (
	RK3BookValue  = 1010 //
	RK3P_On_E     = 1030
	RK3P_On_S     = 1040
	RK3P_On_BV    = 1050
	RK3Cap        = 1060 //
	RK3DivSum5Y   = 1070 //
	RK3DivSum3Y   = 1080 //
	RK3DivYield5Y = 1090 //
	RK3DivYield3Y = 1100 //
	RK3DSI        = 1110 //
)

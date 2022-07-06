package cube

import (
	"math"
	"time"

	"github.com/softilium/mb4/ent"
)

const (
	//Pnl src
	rk2Revenue          = 100
	rk2Amortization     = 110
	rk2OperatingIncome  = 120
	rk2InterestIncome   = 130
	rk2InterestExpenses = 140
	rk2IncomeTax        = 150
	rk2NetIncome        = 160

	//Pnl calculated
	rk2OIBDA             = 200
	rk2EBITDA            = 210
	rk2OIBDAMargin       = 220
	rk2EBITDAMargin      = 230
	rk2OperationalMargin = 240
	rk2NetMargin         = 250
	rk2Debt_On_EBITDA    = 260
	rk2EV_On_EBITDA      = 270
	rk2ROE               = 280

	// CF src
	rk2Cash                  = 300
	rk2NonCurrentLiabilities = 305
	rk2CurrentLiabilities    = 310
	rk2NonControlling        = 315
	rk2Equity                = 320
	rk2Total                 = 325

	// CF calculated
	rk2NetDebt = 400
	rk2EV      = 410
)

type PnlValue struct {
	RK           int
	Ytd          float64
	YtdAdjust    float64
	Ltm          float64
	AnnualGrow   float64
	AnualGrowLtm float64
}

func (p *PnlValue) Calc(src, prevQ, prevY *Report2) {

	if p.Ytd == 0 {
		return
	}

	p.YtdAdjust = p.Ytd / float64(src.ReportQuarter) * 4
	if p.Ltm == 0 { // skip when we assign it before
		if prevQ == nil || prevY == nil {
			p.Ltm = p.YtdAdjust
		} else {
			p.Ltm = prevY.YV[p.RK].Ytd - prevQ.YV[p.RK].Ytd + p.Ytd
		}
	}
	if prevY == nil {
		p.AnnualGrow = 0
		p.AnualGrowLtm = 0
	} else {
		p.AnnualGrow = RoundX(p.Ytd/prevY.YV[p.RK].Ytd*100, 1) - 100
		p.AnualGrowLtm = RoundX(p.Ltm/prevY.YV[p.RK].Ltm*100, 1) - 100
	}

}

type CfValue struct {
	RK         int
	Sld        float64
	AnnualGrow float64
}

func (p *CfValue) Calc(prevY *Report2) {

	if prevY == nil {
		p.AnnualGrow = 0
	} else {
		p.AnnualGrow = RoundX(p.Sld/prevY.SV[p.RK].Sld*100, 1) - 100
	}

}

type Report2 struct { // enriched report with calculated fields
	ReportQuarter int
	ReportDate    time.Time

	prevYear    *Report2
	prevQuarter *Report2

	YV map[int]*PnlValue // year-based values
	SV map[int]*CfValue  // saldo-based values

}

func (r *Report2) Load(s *ent.Report, prevY, prevQ *Report2) {

	r.ReportQuarter = s.ReportQuarter
	r.ReportDate = s.ReportDate

	r.YV = make(map[int]*PnlValue, 15)
	r.SV = make(map[int]*CfValue, 6)

	r.prevQuarter = prevQ
	r.prevYear = prevY

	// Pnl src
	r.YV[rk2Revenue] = &PnlValue{RK: rk2Revenue, Ytd: s.PnlRevenueYtd}
	r.YV[rk2Amortization] = &PnlValue{RK: rk2Amortization, Ytd: s.PnlAmortizationYtd}
	r.YV[rk2OperatingIncome] = &PnlValue{RK: rk2OperatingIncome, Ytd: s.PnlOperatingIncomeYtd}
	r.YV[rk2InterestIncome] = &PnlValue{RK: rk2InterestIncome, Ytd: s.PnlInterestIncomeYtd}
	r.YV[rk2InterestExpenses] = &PnlValue{RK: rk2InterestExpenses, Ytd: s.PnlInterestExpensesYtd}
	r.YV[rk2IncomeTax] = &PnlValue{RK: rk2IncomeTax, Ytd: s.PnlIncomeTaxYtd}
	r.YV[rk2NetIncome] = &PnlValue{RK: rk2NetIncome, Ytd: s.PnlNetIncomeYtd}
	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}

	// CF src
	r.SV[rk2Cash] = &CfValue{Sld: s.CfCashSld}
	r.SV[rk2NonCurrentLiabilities] = &CfValue{Sld: s.CfNonCurrentLiabilitiesSld}
	r.SV[rk2CurrentLiabilities] = &CfValue{Sld: s.CfCurrentLiabilitesSld}
	r.SV[rk2NonControlling] = &CfValue{Sld: s.CfNonControllingSld}
	r.SV[rk2Equity] = &CfValue{Sld: s.CfEquitySld}
	r.SV[rk2Total] = &CfValue{Sld: s.CfTotalSld}
	for _, v := range r.SV {
		v.Calc(r.prevYear)
	}

	// CF calculated
	r.SV[rk2NetDebt] = &CfValue{
		RK:  rk2NetDebt,
		Sld: r.SV[rk2NonCurrentLiabilities].Sld + r.SV[rk2CurrentLiabilities].Sld - r.SV[rk2Cash].Sld,
	}
	r.SV[rk2EV] = &CfValue{
		RK:  rk2EV,
		Sld: r.SV[rk2Cash].Sld + r.SV[rk2NonControlling].Sld + r.SV[rk2NonCurrentLiabilities].Sld + r.SV[rk2CurrentLiabilities].Sld,
	}

	// Pnl calculated

	r.YV[rk2OIBDA] = &PnlValue{
		RK:  rk2OIBDA,
		Ytd: r.YV[rk2Revenue].Ytd - r.YV[rk2Amortization].Ytd,
		Ltm: r.YV[rk2Revenue].Ltm - r.YV[rk2Amortization].Ltm,
	}
	r.YV[rk2EBITDA] = &PnlValue{
		RK:  rk2EBITDA,
		Ytd: r.YV[rk2NetIncome].Ytd - r.YV[rk2Amortization].Ytd - r.YV[rk2InterestIncome].Ytd - r.YV[rk2InterestExpenses].Ytd - r.YV[rk2IncomeTax].Ytd,
		Ltm: r.YV[rk2NetIncome].Ltm - r.YV[rk2Amortization].Ltm - r.YV[rk2InterestIncome].Ltm - r.YV[rk2InterestExpenses].Ltm - r.YV[rk2IncomeTax].Ltm,
	}

	r.YV[rk2OIBDAMargin] = &PnlValue{RK: rk2OIBDAMargin}
	r.YV[rk2EBITDAMargin] = &PnlValue{RK: rk2EBITDAMargin}
	r.YV[rk2OperationalMargin] = &PnlValue{RK: rk2OperationalMargin}
	r.YV[rk2NetMargin] = &PnlValue{RK: rk2NetMargin}
	if math.Abs(s.PnlRevenueYtd) >= 0.01 {

		r.YV[rk2OIBDAMargin].Ytd = RoundX(r.YV[rk2OIBDA].Ytd/r.YV[rk2Revenue].Ytd*100, 1)
		r.YV[rk2OIBDAMargin].Ltm = RoundX(r.YV[rk2OIBDA].Ltm/r.YV[rk2Revenue].Ltm*100, 1)

		r.YV[rk2EBITDAMargin].Ytd = RoundX(r.YV[rk2EBITDA].Ytd/r.YV[rk2Revenue].Ytd*100, 1)
		r.YV[rk2EBITDAMargin].Ltm = RoundX(r.YV[rk2EBITDA].Ltm/r.YV[rk2Revenue].Ltm*100, 1)

		r.YV[rk2OperationalMargin].Ytd = RoundX(r.YV[rk2OperatingIncome].Ytd/r.YV[rk2Revenue].Ytd*100, 1)
		r.YV[rk2OperationalMargin].Ltm = RoundX(r.YV[rk2OperatingIncome].Ltm/r.YV[rk2Revenue].Ltm*100, 1)

		r.YV[rk2NetMargin].Ytd = RoundX(r.YV[rk2NetIncome].Ytd/r.YV[rk2Revenue].Ytd*100, 1)
		r.YV[rk2NetMargin].Ltm = RoundX(r.YV[rk2NetIncome].Ltm/r.YV[rk2Revenue].Ltm*100, 1)
	}

	r.YV[rk2ROE] = &PnlValue{RK: rk2ROE}
	if math.Abs(r.SV[rk2Total].Sld) >= 0.01 {
		r.YV[rk2ROE].Ytd = RoundX(r.YV[rk2NetIncome].Ytd/r.SV[rk2Total].Sld*100, 1)
		r.YV[rk2ROE].Ltm = RoundX(r.YV[rk2NetIncome].Ltm/r.SV[rk2Total].Sld*100, 1)
	}
	r.YV[rk2Debt_On_EBITDA] = &PnlValue{RK: rk2Debt_On_EBITDA}
	r.YV[rk2EV_On_EBITDA] = &PnlValue{RK: rk2EV_On_EBITDA}
	if math.Abs(r.YV[rk2EBITDA].Ytd) >= 0.01 {
		r.YV[rk2Debt_On_EBITDA].Ytd = r.SV[rk2NetDebt].Sld / r.YV[rk2EBITDA].Ytd
		r.YV[rk2Debt_On_EBITDA].Ltm = r.SV[rk2NetDebt].Sld / r.YV[rk2EBITDA].Ltm
		r.YV[rk2EV_On_EBITDA].Ytd = r.SV[rk2EV].Sld / r.YV[rk2EBITDA].Ytd
		r.YV[rk2EV_On_EBITDA].Ltm = r.SV[rk2EV].Sld / r.YV[rk2EBITDA].Ltm
	}

	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}

}

const (
	rk3BookValue = 1010
	rk3P_On_E    = 1030
	rk3P_On_S    = 1040
	rk3P_On_BV   = 1050
)

type CellReport struct { // enriched cell with calculated fields for day (from Cell.D)
	R2 *Report2
	V  map[int]*CfValue
}

func (r *CellReport) Calc(c *Cell) {

	r.V = make(map[int]*CfValue, 4)

}

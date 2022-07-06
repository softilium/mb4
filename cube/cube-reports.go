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

	YV map[int]*PnlValue // year-based values (PNL, ...)
	SV map[int]*CfValue  // saldo-based values (CF, ...)

}

func (r *Report2) Load(s *ent.Report, prevY, prevQ *Report2) {

	r.ReportQuarter = s.ReportQuarter
	r.ReportDate = s.ReportDate

	r.YV = make(map[int]*PnlValue, 15)
	r.SV = make(map[int]*CfValue, 6)

	r.prevQuarter = prevQ
	r.prevYear = prevY

	// Pnl src
	r.YV[RK2Revenue] = &PnlValue{RK: RK2Revenue, Ytd: s.PnlRevenueYtd}
	r.YV[RK2Amortization] = &PnlValue{RK: RK2Amortization, Ytd: s.PnlAmortizationYtd}
	r.YV[RK2OperatingIncome] = &PnlValue{RK: RK2OperatingIncome, Ytd: s.PnlOperatingIncomeYtd}
	r.YV[RK2InterestIncome] = &PnlValue{RK: RK2InterestIncome, Ytd: s.PnlInterestIncomeYtd}
	r.YV[RK2InterestExpenses] = &PnlValue{RK: RK2InterestExpenses, Ytd: s.PnlInterestExpensesYtd}
	r.YV[RK2IncomeTax] = &PnlValue{RK: RK2IncomeTax, Ytd: s.PnlIncomeTaxYtd}
	r.YV[RK2NetIncome] = &PnlValue{RK: RK2NetIncome, Ytd: s.PnlNetIncomeYtd}
	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}

	// CF src
	r.SV[RK2Cash] = &CfValue{Sld: s.CfCashSld}
	r.SV[RK2NonCurrentLiabilities] = &CfValue{Sld: s.CfNonCurrentLiabilitiesSld}
	r.SV[RK2CurrentLiabilities] = &CfValue{Sld: s.CfCurrentLiabilitesSld}
	r.SV[RK2NonControlling] = &CfValue{Sld: s.CfNonControllingSld}
	r.SV[RK2Equity] = &CfValue{Sld: s.CfEquitySld}
	r.SV[RK2Total] = &CfValue{Sld: s.CfTotalSld}

	// CF calculated
	r.SV[RK2NetDebt] = &CfValue{
		RK:  RK2NetDebt,
		Sld: r.SV[RK2NonCurrentLiabilities].Sld + r.SV[RK2CurrentLiabilities].Sld - r.SV[RK2Cash].Sld,
	}
	r.SV[RK2EV] = &CfValue{
		RK:  RK2EV,
		Sld: r.SV[RK2Cash].Sld + r.SV[RK2NonControlling].Sld + r.SV[RK2NonCurrentLiabilities].Sld + r.SV[RK2CurrentLiabilities].Sld,
	}

	// Pnl calculated

	r.YV[RK2OIBDA] = &PnlValue{
		RK:  RK2OIBDA,
		Ytd: r.YV[RK2Revenue].Ytd - r.YV[RK2Amortization].Ytd,
		Ltm: r.YV[RK2Revenue].Ltm - r.YV[RK2Amortization].Ltm,
	}
	r.YV[RK2EBITDA] = &PnlValue{
		RK:  RK2EBITDA,
		Ytd: r.YV[RK2NetIncome].Ytd - r.YV[RK2Amortization].Ytd - r.YV[RK2InterestIncome].Ytd - r.YV[RK2InterestExpenses].Ytd - r.YV[RK2IncomeTax].Ytd,
		Ltm: r.YV[RK2NetIncome].Ltm - r.YV[RK2Amortization].Ltm - r.YV[RK2InterestIncome].Ltm - r.YV[RK2InterestExpenses].Ltm - r.YV[RK2IncomeTax].Ltm,
	}

	r.YV[RK2OIBDAMargin] = &PnlValue{RK: RK2OIBDAMargin}
	r.YV[RK2EBITDAMargin] = &PnlValue{RK: RK2EBITDAMargin}
	r.YV[RK2OperationalMargin] = &PnlValue{RK: RK2OperationalMargin}
	r.YV[RK2NetMargin] = &PnlValue{RK: RK2NetMargin}
	if math.Abs(s.PnlRevenueYtd) >= 0.01 {

		r.YV[RK2OIBDAMargin].Ytd = RoundX(r.YV[RK2OIBDA].Ytd/r.YV[RK2Revenue].Ytd*100, 1)
		r.YV[RK2OIBDAMargin].Ltm = RoundX(r.YV[RK2OIBDA].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2EBITDAMargin].Ytd = RoundX(r.YV[RK2EBITDA].Ytd/r.YV[RK2Revenue].Ytd*100, 1)
		r.YV[RK2EBITDAMargin].Ltm = RoundX(r.YV[RK2EBITDA].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2OperationalMargin].Ytd = RoundX(r.YV[RK2OperatingIncome].Ytd/r.YV[RK2Revenue].Ytd*100, 1)
		r.YV[RK2OperationalMargin].Ltm = RoundX(r.YV[RK2OperatingIncome].Ltm/r.YV[RK2Revenue].Ltm*100, 1)

		r.YV[RK2NetMargin].Ytd = RoundX(r.YV[RK2NetIncome].Ytd/r.YV[RK2Revenue].Ytd*100, 1)
		r.YV[RK2NetMargin].Ltm = RoundX(r.YV[RK2NetIncome].Ltm/r.YV[RK2Revenue].Ltm*100, 1)
	}

	r.YV[RK2ROE] = &PnlValue{RK: RK2ROE}
	if math.Abs(r.SV[RK2Total].Sld) >= 0.01 {
		r.YV[RK2ROE].Ytd = RoundX(r.YV[RK2NetIncome].Ytd/r.SV[RK2Total].Sld*100, 1)
		r.YV[RK2ROE].Ltm = RoundX(r.YV[RK2NetIncome].Ltm/r.SV[RK2Total].Sld*100, 1)
	}
	r.YV[RK2Debt_On_EBITDA] = &PnlValue{RK: RK2Debt_On_EBITDA}
	r.YV[RK2EV_On_EBITDA] = &PnlValue{RK: RK2EV_On_EBITDA}
	if math.Abs(r.YV[RK2EBITDA].Ytd) >= 0.01 {
		r.YV[RK2Debt_On_EBITDA].Ytd = r.SV[RK2NetDebt].Sld / r.YV[RK2EBITDA].Ytd
		r.YV[RK2Debt_On_EBITDA].Ltm = r.SV[RK2NetDebt].Sld / r.YV[RK2EBITDA].Ltm
		r.YV[RK2EV_On_EBITDA].Ytd = r.SV[RK2EV].Sld / r.YV[RK2EBITDA].Ytd
		r.YV[RK2EV_On_EBITDA].Ltm = r.SV[RK2EV].Sld / r.YV[RK2EBITDA].Ltm
	}

	for _, v := range r.YV {
		v.Calc(r, r.prevQuarter, r.prevYear)
	}
	for _, v := range r.SV {
		v.Calc(r.prevYear)
	}

}

const (
	RK3BookValue = 1010
	RK3P_On_E    = 1030
	RK3P_On_S    = 1040
	RK3P_On_BV   = 1050
)

type CellReport struct { // enriched cell with calculated fields for day (from Cell.D)
	R2 *Report2
	V  map[int]*CfValue
}

func (r *CellReport) Calc(c *Cell) {

	r.V = make(map[int]*CfValue, 4)

}

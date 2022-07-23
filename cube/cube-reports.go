package cube

import (
	"math"
	"time"

	"github.com/softilium/mb4/ent"
)

type RepV struct {
	V      float64 //Ytd/Sld
	YtdAdj float64
	Ltm    float64
	AG     float64
	AGLtm  float64
	FromR2 func(*Report2) RepV `json:"-"`
}

func (p *RepV) SetFromPnl(newV, newLtm float64, src, prevQ, prevY *Report2) {

	p.V = newV
	p.Ltm = newLtm

	p.YtdAdj = p.V / float64(src.ReportQuarter) * 4
	if p.Ltm == 0 { // skip when we assign it before
		if prevQ == nil || prevY == nil {
			p.Ltm = p.YtdAdj
		} else {
			p.Ltm = p.FromR2(prevY).V - p.FromR2(prevQ).V + p.V
		}
	}
	if prevY == nil {
		p.AG = 0
		p.AGLtm = 0
	} else {
		p.AG = RoundX(p.V/p.FromR2(prevY).V*100, 1) - 100
		p.AGLtm = RoundX(p.Ltm/p.FromR2(prevY).Ltm*100, 1) - 100
	}

}

func (p *RepV) CalcCashflowAnnualGrowth(prevY *Report2) {

	if prevY == nil {
		p.AG = 0
	} else {
		p.AG = RoundX(p.V/p.FromR2(prevY).V*100, 1) - 100
	}

}

type Report2 struct { // enriched report with calculated fields
	ReportQuarter int
	ReportDate    time.Time

	prevYear    *Report2
	prevQuarter *Report2

	// Pnl src
	Revenue          RepV
	Amortization     RepV
	OperatingIncome  RepV
	InterestIncome   RepV
	InterestExpenses RepV
	IncomeTax        RepV
	NetIncome        RepV

	// Pnl calculated
	OIBDA             RepV
	EBITDA            RepV
	OIBDAMargin       RepV
	EBITDAMargin      RepV
	OperationalMargin RepV
	NetMargin         RepV
	Debt_on_EBITDA    RepV
	EV_on_EBITDA      RepV
	ROE               RepV

	// Cf src
	Cash                  RepV
	NonCurrentLiabilities RepV
	CurrentLiabilities    RepV
	NonControlling        RepV
	Equity                RepV
	Total                 RepV

	// Cf calculated
	NetDebt RepV
	EV      RepV
}

func (r *Report2) Load(s *ent.Report, prevY, prevQ *Report2) {

	r.ReportQuarter = s.ReportQuarter
	r.ReportDate = s.ReportDate

	r.prevQuarter = prevQ
	r.prevYear = prevY

	// init rep fields

	r.Revenue.FromR2 = func(r *Report2) RepV { return r.Revenue }
	r.Amortization.FromR2 = func(r *Report2) RepV { return r.Amortization }
	r.OperatingIncome.FromR2 = func(r *Report2) RepV { return r.OperatingIncome }
	r.InterestIncome.FromR2 = func(r *Report2) RepV { return r.InterestIncome }
	r.InterestExpenses.FromR2 = func(r *Report2) RepV { return r.InterestExpenses }
	r.IncomeTax.FromR2 = func(r *Report2) RepV { return r.IncomeTax }
	r.NetIncome.FromR2 = func(r *Report2) RepV { return r.NetIncome }
	r.OIBDA.FromR2 = func(r *Report2) RepV { return r.OIBDA }
	r.EBITDA.FromR2 = func(r *Report2) RepV { return r.EBITDA }
	r.OIBDAMargin.FromR2 = func(r *Report2) RepV { return r.OIBDAMargin }
	r.EBITDAMargin.FromR2 = func(r *Report2) RepV { return r.EBITDAMargin }
	r.OperationalMargin.FromR2 = func(r *Report2) RepV { return r.OperationalMargin }
	r.NetMargin.FromR2 = func(r *Report2) RepV { return r.NetMargin }
	r.Debt_on_EBITDA.FromR2 = func(r *Report2) RepV { return r.Debt_on_EBITDA }
	r.EV_on_EBITDA.FromR2 = func(r *Report2) RepV { return r.EV_on_EBITDA }
	r.ROE.FromR2 = func(r *Report2) RepV { return r.ROE }

	// Pnl src

	r.Revenue.SetFromPnl(s.PnlRevenueYtd, 0, r, r.prevQuarter, r.prevYear)
	r.Amortization.SetFromPnl(s.PnlAmortizationYtd, 0, r, r.prevQuarter, r.prevYear)
	r.OperatingIncome.SetFromPnl(s.PnlOperatingIncomeYtd, 0, r, r.prevQuarter, r.prevYear)
	r.InterestIncome.SetFromPnl(s.PnlInterestIncomeYtd, 0, r, r.prevQuarter, r.prevYear)
	r.InterestExpenses.SetFromPnl(s.PnlInterestExpensesYtd, 0, r, r.prevQuarter, r.prevYear)
	r.IncomeTax.SetFromPnl(s.PnlIncomeTaxYtd, 0, r, r.prevQuarter, r.prevYear)
	r.NetIncome.SetFromPnl(s.PnlNetIncomeYtd, 0, r, r.prevQuarter, r.prevYear)

	// CF src
	r.Cash.V = s.CfCashSld
	r.Cash.CalcCashflowAnnualGrowth(r.prevYear)

	r.NonCurrentLiabilities.V = s.CfNonCurrentLiabilitiesSld
	r.NonCurrentLiabilities.CalcCashflowAnnualGrowth(r.prevYear)

	r.CurrentLiabilities.V = s.CfCurrentLiabilitesSld
	r.CurrentLiabilities.CalcCashflowAnnualGrowth(r.prevYear)

	r.NonControlling.V = s.CfNonControllingSld
	r.NonControlling.CalcCashflowAnnualGrowth(r.prevYear)

	r.Equity.V = s.CfEquitySld
	r.Equity.CalcCashflowAnnualGrowth(r.prevYear)

	r.Total.V = s.CfTotalSld
	r.Total.CalcCashflowAnnualGrowth(r.prevYear)

	// CF calculated
	r.NetDebt.V = r.NonCurrentLiabilities.V + r.CurrentLiabilities.V - r.Cash.V
	r.NetDebt.CalcCashflowAnnualGrowth(r.prevYear)

	r.EV.V = r.Cash.V + r.NonControlling.V + r.NonCurrentLiabilities.V + r.CurrentLiabilities.V
	r.EV.CalcCashflowAnnualGrowth(r.prevYear)

	// Pnl calculated

	r.OIBDA.SetFromPnl(
		r.Revenue.V-r.Amortization.V,
		r.Revenue.Ltm-r.Amortization.Ltm,
		r, r.prevQuarter, r.prevYear)

	r.EBITDA.SetFromPnl(
		r.NetIncome.V-r.Amortization.V-r.InterestIncome.V-r.InterestExpenses.V-r.IncomeTax.V,
		r.NetIncome.Ltm-r.Amortization.Ltm-r.InterestIncome.Ltm-r.InterestExpenses.Ltm-r.IncomeTax.Ltm,
		r, r.prevQuarter, r.prevYear)

	if math.Abs(s.PnlRevenueYtd) >= 0.01 {

		r.OIBDAMargin.SetFromPnl(
			RoundX(r.OIBDA.V/r.Revenue.V*100, 1),
			RoundX(r.OIBDA.Ltm/r.Revenue.Ltm*100, 1),
			r, r.prevQuarter, r.prevYear)

		r.EBITDAMargin.SetFromPnl(
			RoundX(r.EBITDA.V/r.Revenue.V*100, 1),
			RoundX(r.EBITDA.Ltm/r.Revenue.Ltm*100, 1),
			r, r.prevQuarter, r.prevYear)

		r.OperationalMargin.SetFromPnl(
			RoundX(r.OperatingIncome.V/r.Revenue.V*100, 1),
			RoundX(r.OperatingIncome.Ltm/r.Revenue.Ltm*100, 1),
			r, r.prevQuarter, r.prevYear)

		r.NetMargin.SetFromPnl(
			RoundX(r.NetIncome.V/r.Revenue.V*100, 1),
			RoundX(r.NetIncome.Ltm/r.Revenue.Ltm*100, 1),
			r, r.prevQuarter, r.prevYear)
	}

	if math.Abs(r.Total.V) >= 0.01 {
		r.ROE.SetFromPnl(
			RoundX(r.NetIncome.V/r.Total.V*100, 1),
			RoundX(r.NetIncome.Ltm/r.Total.V*100, 1),
			r, r.prevQuarter, r.prevYear)
	}
	if math.Abs(r.EBITDA.V) >= 0.01 {

		r.Debt_on_EBITDA.SetFromPnl(
			r.NetDebt.V/r.EBITDA.V,
			r.NetDebt.V/r.EBITDA.Ltm,
			r, r.prevQuarter, r.prevYear)

		r.EV_on_EBITDA.SetFromPnl(
			r.EV.V/r.EBITDA.V,
			r.EV.V/r.EBITDA.Ltm,
			r, r.prevQuarter, r.prevYear)
	}

}

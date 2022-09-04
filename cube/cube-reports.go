package cube

import (
	"math"
	"time"

	"github.com/softilium/mb4/ent"
)

// TODO Отношение EV к BookValue - When compared to the company's market value, book value can indicate whether a stock is under- or overpriced.
type Report2 struct { // enriched report with calculated fields
	ReportYear    int
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
	ROE               RepV

	// Cf src
	Cash                  RepS
	NonCurrentLiabilities RepS
	CurrentLiabilities    RepS
	NonControlling        RepS
	Equity                RepS
	Total                 RepS

	// Cf calculated
	NetDebt RepS
}

func (r *Report2) Init() {

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
	r.ROE.FromR2 = func(r *Report2) RepV { return r.ROE }

	r.Cash.FromR2 = func(r *Report2) RepS { return r.Cash }
	r.NonCurrentLiabilities.FromR2 = func(r *Report2) RepS { return r.NonCurrentLiabilities }
	r.CurrentLiabilities.FromR2 = func(r *Report2) RepS { return r.CurrentLiabilities }
	r.NonControlling.FromR2 = func(r *Report2) RepS { return r.NonControlling }
	r.Equity.FromR2 = func(r *Report2) RepS { return r.Equity }
	r.Total.FromR2 = func(r *Report2) RepS { return r.Total }
	r.NetDebt.FromR2 = func(r *Report2) RepS { return r.NetDebt }

}

func (r *Report2) LoadFromRawReport(s *ent.Report, prevY, prevQ *Report2) {

	r.ReportQuarter = s.ReportQuarter
	r.ReportYear = s.ReportYear
	r.ReportDate = s.ReportDate

	r.prevQuarter = prevQ
	r.prevYear = prevY

	r.Init()

	// Pnl src

	r.Revenue.SetFromPnl(s.PnlRevenueYtd, 0, r)
	r.Amortization.SetFromPnl(s.PnlAmortizationYtd, 0, r)
	r.OperatingIncome.SetFromPnl(s.PnlOperatingIncomeYtd, 0, r)
	r.InterestIncome.SetFromPnl(s.PnlInterestIncomeYtd, 0, r)
	r.InterestExpenses.SetFromPnl(s.PnlInterestExpensesYtd, 0, r)
	r.IncomeTax.SetFromPnl(s.PnlIncomeTaxYtd, 0, r)
	r.NetIncome.SetFromPnl(s.PnlNetIncomeYtd, 0, r)

	// CF src
	r.Cash.S = s.CfCashSld
	r.Cash.CalcAG(r.prevYear)

	r.NonCurrentLiabilities.S = s.CfNonCurrentLiabilitiesSld
	r.NonCurrentLiabilities.CalcAG(r.prevYear)

	r.CurrentLiabilities.S = s.CfCurrentLiabilitesSld
	r.CurrentLiabilities.CalcAG(r.prevYear)

	r.NonControlling.S = s.CfNonControllingSld
	r.NonControlling.CalcAG(r.prevYear)

	r.Equity.S = s.CfEquitySld
	r.Equity.CalcAG(r.prevYear)

	r.Total.S = s.CfTotalSld
	r.Total.CalcAG(r.prevYear)

	// CF calculated
	r.NetDebt.S = r.NonCurrentLiabilities.S + r.CurrentLiabilities.S - r.Cash.S
	r.NetDebt.CalcAG(r.prevYear)

	// Pnl calculated

	r.OIBDA.SetFromPnl(r.Revenue.Src-r.Amortization.Src, r.Revenue.Ltm-r.Amortization.Ltm, r)

	r.EBITDA.SetFromPnl(
		r.NetIncome.Src-r.Amortization.Src-r.InterestIncome.Src-r.InterestExpenses.Src-r.IncomeTax.Src,
		r.NetIncome.Ltm-r.Amortization.Ltm-r.InterestIncome.Ltm-r.InterestExpenses.Ltm-r.IncomeTax.Ltm,
		r)

	r.CalcMults()

}

func (r *Report2) CalcMults() {

	r.OIBDAMargin.YtdAdj = IIF(math.Abs(r.Revenue.YtdAdj) < 0.01, 0, RoundX(r.OIBDA.YtdAdj/r.Revenue.YtdAdj*100, 1))
	r.OIBDAMargin.Ltm = IIF(math.Abs(r.Revenue.Ltm) < 0.01, 0, RoundX(r.OIBDA.Ltm/r.Revenue.Ltm*100, 1))
	r.OIBDAMargin.CalcAG(r.prevYear)

	r.EBITDAMargin.YtdAdj = IIF(math.Abs(r.Revenue.YtdAdj) < 0.01, 0, RoundX(r.EBITDA.YtdAdj/r.Revenue.YtdAdj*100, 1))
	r.EBITDAMargin.Ltm = IIF(math.Abs(r.Revenue.Ltm) < 0.01, 0, RoundX(r.EBITDA.Ltm/r.Revenue.Ltm*100, 1))
	r.EBITDAMargin.CalcAG(r.prevYear)

	r.OperationalMargin.YtdAdj = IIF(math.Abs(r.Revenue.YtdAdj) < 0.01, 0, RoundX(r.OperatingIncome.YtdAdj/r.Revenue.YtdAdj*100, 1))
	r.OperationalMargin.Ltm = IIF(math.Abs(r.Revenue.Ltm) < 0.01, 0, RoundX(r.OperatingIncome.Ltm/r.Revenue.Ltm*100, 1))
	r.OperationalMargin.CalcAG(r.prevYear)

	r.NetMargin.YtdAdj = IIF(math.Abs(r.Revenue.YtdAdj) < 0.01, 0, RoundX(r.NetIncome.YtdAdj/r.Revenue.YtdAdj*100, 1))
	r.NetMargin.Ltm = IIF(math.Abs(r.Revenue.Ltm) < 0.01, 0, RoundX(r.NetIncome.Ltm/r.Revenue.Ltm*100, 1))
	r.NetMargin.CalcAG(r.prevYear)

	r.ROE.YtdAdj = IIF(math.Abs(r.Total.S) < 0.01, 0, RoundX(r.NetIncome.YtdAdj/r.Total.S*100, 1))
	r.ROE.Ltm = IIF(math.Abs(r.Total.S) < 0.01, 0, RoundX(r.NetIncome.Ltm/r.Total.S*100, 1))
	r.ROE.CalcAG(r.prevYear)

	r.Debt_on_EBITDA.YtdAdj = IIF(math.Abs(r.EBITDA.YtdAdj) < 0.01, 0, RoundX(r.NetDebt.S/r.EBITDA.YtdAdj, 1))
	r.Debt_on_EBITDA.Ltm = IIF(math.Abs(r.EBITDA.Ltm) < 0.01, 0, RoundX(r.NetDebt.S/r.EBITDA.Ltm, 1))
	r.Debt_on_EBITDA.CalcAG(r.prevYear)

}

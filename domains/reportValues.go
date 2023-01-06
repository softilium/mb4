package domains

type ReportValue int

// ReportKinds - enum all possible values for a report (for strategy editor or so)
const (
	// Pnl Src
	RK_Revenue          ReportValue = 100
	RK_Amortization     ReportValue = 200
	RK_OperatingIncome  ReportValue = 300
	RK_InterestIncome   ReportValue = 400
	RK_InterestExpenses ReportValue = 500
	RK_IncomeTax        ReportValue = 600
	RK_NetIncome        ReportValue = 700

	// Pnl calculated
	RK_OIBDA             ReportValue = 1100
	RK_EBITDA            ReportValue = 1200
	RK_OIBDAMargin       ReportValue = 1300
	RK_EBITDAMargin      ReportValue = 1400
	RK_OperationalMargin ReportValue = 1500
	RK_NetMargin         ReportValue = 1600
	RK_Debt_on_EBITDA    ReportValue = 1700
	RK_EV_on_EBITDA      ReportValue = 1800
	RK_ROE               ReportValue = 1900

	// Cf src
	RK_Cash                  ReportValue = 2100
	RK_NonCurrentLiabilities ReportValue = 2200
	RK_CurrentLiabilities    ReportValue = 2300
	RK_NonControlling        ReportValue = 2400
	RK_Equity                ReportValue = 2500
	RK_Total                 ReportValue = 2600

	// Cf calculated
	RK_NetDebt ReportValue = 2700
	RK_EV      ReportValue = 2800

	//R3
	RK_BookValue  ReportValue = 3100
	RK_P_on_E     ReportValue = 3200
	RK_P_on_BV    ReportValue = 3300
	RK_Cap        ReportValue = 3400
	RK_P_on_S     ReportValue = 3500
	RK_DivSum5Y   ReportValue = 3600
	RK_DivSum3Y   ReportValue = 3700
	RK_DivYield5Y ReportValue = 3800
	RK_DivYield3Y ReportValue = 3900
	RK_DSI        ReportValue = 3950
)

var ReportValues Domain[ReportValue] = Domain[ReportValue]{}

func init() {

	ReportValues.init([]DomainItem[ReportValue]{

		// Pnl Src
		{Id: RK_Revenue, Descr: "Revenue"},
		{Id: RK_Amortization, Descr: "Amortization"},
		{Id: RK_OperatingIncome, Descr: "Operational income"},
		{Id: RK_InterestIncome, Descr: "Financial incoming"},
		{Id: RK_InterestExpenses, Descr: "Financial expences"},
		{Id: RK_IncomeTax, Descr: "Income tax"},
		{Id: RK_NetIncome, Descr: "Net income"},

		// Pnl calculated

		{Id: RK_OIBDA, Descr: "OIBDA"},
		{Id: RK_EBITDA, Descr: "EBITDA"},
		{Id: RK_OIBDAMargin, Descr: "OIBDA margin"},
		{Id: RK_EBITDAMargin, Descr: "EBITDA margin"},
		{Id: RK_OperationalMargin, Descr: "Operational margin"},
		{Id: RK_NetMargin, Descr: "Net margin"},
		{Id: RK_Debt_on_EBITDA, Descr: "Debt/EBITDA"},
		{Id: RK_EV_on_EBITDA, Descr: "EV/EBITDA"},
		{Id: RK_ROE, Descr: "ROE"},

		// Cf src
		{Id: RK_Cash, Descr: "Cash and equivalents"},
		{Id: RK_NonCurrentLiabilities, Descr: "Non-current liabilities"},
		{Id: RK_CurrentLiabilities, Descr: "Current liabilities"},
		{Id: RK_NonControlling, Descr: "Non-controlling shares"},
		{Id: RK_Equity, Descr: "Equity"},
		{Id: RK_Total, Descr: "Total"},

		// Cf calculated
		{Id: RK_NetDebt, Descr: "Net debt"},
		{Id: RK_EV, Descr: "EV"},

		//R3
		{Id: RK_BookValue, Descr: "Book value"},
		{Id: RK_P_on_E, Descr: "P/E"},
		{Id: RK_P_on_BV, Descr: "P/BV"},
		{Id: RK_Cap, Descr: "Market cap"},
		{Id: RK_P_on_S, Descr: "P/S"},
		{Id: RK_DivSum5Y, Descr: "Dividends for 5y"},
		{Id: RK_DivSum3Y, Descr: "Dividends for 3y"},
		{Id: RK_DivYield5Y, Descr: "Dividend yield for 5y"},
		{Id: RK_DivYield3Y, Descr: "Vididend yield for 3y"},
		{Id: RK_DSI, Descr: "DSI"},
	})

}

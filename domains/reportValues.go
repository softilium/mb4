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
		{Id: RK_Revenue, Descr: "Выручка"},
		{Id: RK_Amortization, Descr: "Амортизация"},
		{Id: RK_OperatingIncome, Descr: "Операционная прибыль"},
		{Id: RK_InterestIncome, Descr: "Финансовые доходы"},
		{Id: RK_InterestExpenses, Descr: "Финансовые расходы"},
		{Id: RK_IncomeTax, Descr: "Налог на прибыль"},
		{Id: RK_NetIncome, Descr: "Чистая прибыль"},

		// Pnl calculated

		{Id: RK_OIBDA, Descr: "OIBDA"},
		{Id: RK_EBITDA, Descr: "EBITDA"},
		{Id: RK_OIBDAMargin, Descr: "Маржинальность по OIBDA"},
		{Id: RK_EBITDAMargin, Descr: "Маржинальность по EBITDA"},
		{Id: RK_OperationalMargin, Descr: "Маржинальность операционная"},
		{Id: RK_NetMargin, Descr: "Маржинальность чистая"},
		{Id: RK_Debt_on_EBITDA, Descr: "Debt/EBITDA"},
		{Id: RK_EV_on_EBITDA, Descr: "EV/EBITDA"},
		{Id: RK_ROE, Descr: "ROE"},

		// Cf src
		{Id: RK_Cash, Descr: "Денежные средства"},
		{Id: RK_NonCurrentLiabilities, Descr: "Долгосрочные обязательства"},
		{Id: RK_CurrentLiabilities, Descr: "Краткосрочные обязательства"},
		{Id: RK_NonControlling, Descr: "Неконтроллируемые доли в уставном капитале"},
		{Id: RK_Equity, Descr: "Капитал"},
		{Id: RK_Total, Descr: "Активы и обязательства"},

		// Cf calculated
		{Id: RK_NetDebt, Descr: "Чистый долг"},
		{Id: RK_EV, Descr: "EV"},

		//R3
		{Id: RK_BookValue, Descr: "Балансовая стоимость"},
		{Id: RK_P_on_E, Descr: "P/E"},
		{Id: RK_P_on_BV, Descr: "P/BV"},
		{Id: RK_Cap, Descr: "Капитализация"},
		{Id: RK_P_on_S, Descr: "P/S"},
		{Id: RK_DivSum5Y, Descr: "Сумма дивидендов за 5 лет"},
		{Id: RK_DivSum3Y, Descr: "Сумма дивидендов за 3 года"},
		{Id: RK_DivYield5Y, Descr: "Дивидендная доходность за 5 лет"},
		{Id: RK_DivYield3Y, Descr: "Дивидендная доходность за 3 года"},
		{Id: RK_DSI, Descr: "DSI"},
	})

}

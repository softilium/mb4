package domains

type ReportValueType int

const (
	RVT_S      ReportValueType = 100 // SLD
	RVT_YtdAdj ReportValueType = 200 // Year adjusted for YTD
	RVT_Ltm    ReportValueType = 300 // Last 12 months

	RVT_AG        ReportValueType = 400 // Annual growth
	RVT_AG_Ltm    ReportValueType = 500 // Annual growth LTM
	RVT_AG_YtdSld ReportValueType = 600 // Annual growth YTD adj.

	RVT_IndUpside        ReportValueType = 2100 // Industry upside SLD
	RVT_IndUpside_YtdAdj ReportValueType = 2200 // Industry upside Year adjusted for YTD
	RVT_IndUpside_Ltm    ReportValueType = 2300 // Industry upside Last 12 months
)

var ReportValueTypes Domain[ReportValueType] = Domain[ReportValueType]{}

func init() {

	ReportValueTypes.init([]DomainItem[ReportValueType]{

		{Id: RVT_S, Descr: "Из отчета cashflow"},
		{Id: RVT_YtdAdj, Descr: "Из отчета PnL, выравненное по году"},
		{Id: RVT_Ltm, Descr: "LTM"},

		{Id: RVT_AG, Descr: "Рост г/г"},
		{Id: RVT_AG_YtdSld, Descr: "Рост г/г, выравненное по году"},
		{Id: RVT_AG_Ltm, Descr: "Рост г/г, LTM"},

		{Id: RVT_IndUpside, Descr: "Отраслевой апсайд"},
		{Id: RVT_IndUpside_YtdAdj, Descr: "Отраслевой апсайд, выравненно по году"},
		{Id: RVT_IndUpside_Ltm, Descr: "Отраслевой апсайд, LTM"},
	})

}

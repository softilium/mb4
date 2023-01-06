package domains

type ReportValueType int

const (
	RVT_Src    ReportValueType = 50  // Src
	RVT_S      ReportValueType = 100 // SLD
	RVT_YtdAdj ReportValueType = 200 // Year adjusted for YTD
	RVT_Ltm    ReportValueType = 300 // Last 12 months

	RVT_AG        ReportValueType = 400 // Annual growth
	RVT_AG_Ltm    ReportValueType = 500 // Annual growth LTM
	RVT_AG_YtdAdj ReportValueType = 600 // Annual growth YTD adj.

	RVT_IndUpside        ReportValueType = 2100 // Industry upside SLD
	RVT_IndUpside_YtdAdj ReportValueType = 2200 // Industry upside Year adjusted for YTD
	RVT_IndUpside_Ltm    ReportValueType = 2300 // Industry upside Last 12 months
)

var ReportValueTypes Domain[ReportValueType] = Domain[ReportValueType]{}

func init() {

	ReportValueTypes.init([]DomainItem[ReportValueType]{

		{Id: RVT_S, Descr: "CashFlow"},
		{Id: RVT_YtdAdj, Descr: "P&L, year aligned"},
		{Id: RVT_Ltm, Descr: "LTM"},

		{Id: RVT_AG, Descr: "Growth y/y"},
		{Id: RVT_AG_YtdAdj, Descr: "Growth y/y, year aligned"},
		{Id: RVT_AG_Ltm, Descr: "Growth y/y, LTM"},

		{Id: RVT_IndUpside, Descr: "Indistry upside"},
		{Id: RVT_IndUpside_YtdAdj, Descr: "Industry upside, year aligned"},
		{Id: RVT_IndUpside_Ltm, Descr: "Industry upside, LTM"},
	})

}

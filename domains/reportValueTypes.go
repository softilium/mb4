package domains

type ReportValueType int

const (
	RVT_Src    ReportValueType = 100 // YTD/SLD
	RVT_YtdAdj ReportValueType = 200 // Year adjusted for YTD
	RVT_Ltm    ReportValueType = 300 // Last 12 months
	RVT_AG     ReportValueType = 400 // Annual growth
	RVT_AG_Ltm ReportValueType = 500 // Annual growth LTM

	RVT_Ind_Src    ReportValueType = 1100 // Industry YTD/SLD
	RVT_Ind_YtdAdj ReportValueType = 1200 // Industry Year adjusted for YTD
	RVT_Ind_Ltm    ReportValueType = 1300 // Industry Last 12 months
	RVT_Ind_AG     ReportValueType = 1400 // Industry Annual growth
	RVT_Ind_AG_Ltm ReportValueType = 1500 // Industry Annual growth LTM

	RVT_IndUpside_Src    ReportValueType = 2100 // Industry upside YTD/SLD
	RVT_IndUpside_YtdAdj ReportValueType = 2200 // Industry upside Year adjusted for YTD
	RVT_IndUpside_Ltm    ReportValueType = 2300 // Industry upside Last 12 months
	RVT_IndUpside_AG     ReportValueType = 2400 // Industry upside Annual growth
	RVT_IndUpside_AG_Ltm ReportValueType = 2500 // Industry upside Annual growth LTM
)

var ReportValueTypes Domain[ReportValueType] = Domain[ReportValueType]{}

func init() {

	ReportValueTypes.init([]DomainItem[ReportValueType]{

		{Id: RVT_Src, Descr: "YTD/SLD"},
		{Id: RVT_YtdAdj, Descr: "Выровненое к году"},
		{Id: RVT_Ltm, Descr: "LTM"},
		{Id: RVT_AG, Descr: "Рост г/г"},
		{Id: RVT_AG_Ltm, Descr: "Рост LTM г/г"},

		{Id: RVT_Ind_Src, Descr: "Отраслевое YTD/SLD"},
		{Id: RVT_Ind_YtdAdj, Descr: "Отраслевое выровненое к году"},
		{Id: RVT_Ind_Ltm, Descr: "Отраслевое LTM"},
		{Id: RVT_Ind_AG, Descr: "Отраслевой рост г/г"},
		{Id: RVT_Ind_AG_Ltm, Descr: "Отраслевой рост LTM г/г"},

		{Id: RVT_IndUpside_Src, Descr: "Отраслевой апсайд YTD/SLD"},
		{Id: RVT_IndUpside_YtdAdj, Descr: "Отраслевой апсайд выровненый к году"},
		{Id: RVT_IndUpside_Ltm, Descr: "Отраслевой апсайд LTM"},
		{Id: RVT_IndUpside_AG, Descr: "Отраслевой апсайд роста г/г"},
		{Id: RVT_IndUpside_AG_Ltm, Descr: "Отраслевой апсайд роста LTM г/г"},
	})

}

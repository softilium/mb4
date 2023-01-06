package domains

type FilterValueKind int

const (
	FVK_Ticker      FilterValueKind = 100
	FVK_Industry    FilterValueKind = 200
	FVK_ReportValue FilterValueKind = 300
)

var FilterValueKinds Domain[FilterValueKind] = Domain[FilterValueKind]{}

func init() {

	FilterValueKinds.init([]DomainItem[FilterValueKind]{
		{Id: FVK_Ticker, Descr: "Ticker"},
		{Id: FVK_Industry, Descr: "Industry"},
		{Id: FVK_ReportValue, Descr: "Report parameter"},
	})

}

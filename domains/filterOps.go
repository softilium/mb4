package domains

type FilterOp int

const (
	FilterOp_Eq FilterOp = 10
	FilterOp_Lt FilterOp = 20
	FilterOp_Le FilterOp = 30
	FilterOp_Gt FilterOp = 40
	FilterOp_Ge FilterOp = 50
	FilterOp_Ne FilterOp = 60
)

var FilterOps Domain[FilterOp] = Domain[FilterOp]{}

func init() {

	FilterOps.init([]DomainItem[FilterOp]{
		{Id: FilterOp_Eq, Descr: "="},
		{Id: FilterOp_Lt, Descr: "<"},
		{Id: FilterOp_Le, Descr: "<="},
		{Id: FilterOp_Gt, Descr: ">"},
		{Id: FilterOp_Ge, Descr: ">="},
		{Id: FilterOp_Ne, Descr: "<>"},
	})

}

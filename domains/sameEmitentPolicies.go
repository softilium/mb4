package domains

type SameEmitentPolicy = int

const (
	SameEmitentPolicy_Allow        SameEmitentPolicy = 100
	SameEmitentPolicy_PreferPrefs  SameEmitentPolicy = 200
	SameEmitentPolicy_PreferOrd    SameEmitentPolicy = 300
	SameEmitentPolicy_AllowOnlyOne SameEmitentPolicy = 400
)

var SameEmitentPolicies Domain[SameEmitentPolicy] = Domain[SameEmitentPolicy]{}

func init() {

	SameEmitentPolicies.init([]DomainItem[SameEmitentPolicy]{
		{Id: SameEmitentPolicy_Allow, Descr: "allow both"},
		{Id: SameEmitentPolicy_PreferPrefs, Descr: "choose preferred"},
		{Id: SameEmitentPolicy_PreferOrd, Descr: "choose common"},
		{Id: SameEmitentPolicy_AllowOnlyOne, Descr: "choose best"},
	})

}

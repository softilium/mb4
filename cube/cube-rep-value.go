package cube

func IIF[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func Growth(b, p float64, d int) float64 {
	return IIF(p == 0, 0, RoundX((b-p)/p*100, d))
}

type ChangeValulation int

const (
	VeryBad  ChangeValulation = -50
	Bad      ChangeValulation = -15
	Neutral  ChangeValulation = 0
	Good     ChangeValulation = 15
	VeryGood ChangeValulation = 50
)

type RepV struct {
	Src                   float64 //Ytd from rep src
	YtdAdj                float64
	Ltm                   float64
	AGYtdAdj              float64
	AGLtm                 float64
	IndustryUpside_Ltm    float64
	IndustryUpside_YtdAdj float64
	FromR2                func(*Report2) RepV `json:"-"`
}

func (p *RepV) SetFromPnl(newV, newLtm float64, src *Report2) {

	p.Src = newV
	p.Ltm = newLtm

	p.YtdAdj = p.Src / float64(src.ReportQuarter) * 4
	if src.prevQuarter == nil || src.prevYear == nil {
		p.Ltm = p.YtdAdj
	} else {
		p.Ltm = p.FromR2(src.prevYear).Src - p.FromR2(src.prevQuarter).Src + p.Src
	}

	p.CalcAG(src.prevYear)

}

func (p *RepV) CalcAG(prevY *Report2) {

	if prevY == nil {
		p.AGYtdAdj = 0
		p.AGLtm = 0
	} else {
		p.AGYtdAdj = IIF(p.FromR2(prevY).YtdAdj == 0, 0, RoundX((p.YtdAdj-p.FromR2(prevY).YtdAdj)/p.FromR2(prevY).YtdAdj*100, 1))
		p.AGLtm = IIF(p.FromR2(prevY).Ltm == 0, 0, RoundX((p.Ltm-p.FromR2(prevY).Ltm)/p.FromR2(prevY).Ltm*100, 1))
	}

}

type RepS struct {
	S              float64
	AG             float64
	IndustryUpside float64
	FromR2         func(*Report2) RepS `json:"-"`
}

func (p *RepS) CalcAG(prevY *Report2) {

	if prevY == nil {
		p.AG = 0
		return
	}
	p.AG = Growth(p.S, p.FromR2(prevY).S, 1)

}

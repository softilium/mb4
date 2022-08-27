package cube

type RepV struct {
	InverseGrowth bool
	V             float64 //Ytd/Sld
	YtdAdj        float64
	Ltm           float64
	AG            float64
	AGLtm         float64
	FromR2        func(*Report2) RepV `json:"-"`
}

func (p *RepV) SetFromPnl(newV, newLtm float64, src *Report2) {

	p.V = newV
	p.Ltm = newLtm

	p.YtdAdj = p.V / float64(src.ReportQuarter) * 4
	if p.Ltm == 0 { // skip when we assign it before
		if src.prevQuarter == nil || src.prevYear == nil {
			p.Ltm = p.YtdAdj
		} else {
			p.Ltm = p.FromR2(src.prevYear).V - p.FromR2(src.prevQuarter).V + p.V
		}
	}
	if src.prevYear == nil {
		p.AG = 0
		p.AGLtm = 0
	} else {
		p.AG = RoundX(p.V/p.FromR2(src.prevYear).V*100, 1) - 100
		p.AGLtm = RoundX(p.Ltm/p.FromR2(src.prevYear).Ltm*100, 1) - 100
	}

}

func (p *RepV) CalcCashflowAnnualGrowth(prevY *Report2) {

	if prevY == nil {
		p.AG = 0
	} else {
		p.AG = RoundX(p.V/p.FromR2(prevY).V*100, 1) - 100
	}

}

func (p *RepV) CalcIndUpside_V(iv *RepV) float64 {
	if p.InverseGrowth {
		return RoundX((iv.V-p.V)/iv.V*100, 1)
	}
	return RoundX((p.V-iv.V)/iv.V*100, 1)
}

func (p *RepV) CalcIndUpside_YtdAdj(iv *RepV) float64 {
	if p.InverseGrowth {
		return RoundX((iv.YtdAdj-p.YtdAdj)/iv.YtdAdj*100, 1)
	}
	return RoundX((p.YtdAdj-iv.YtdAdj)/iv.YtdAdj*100, 1)
}

func (p *RepV) CalcIndUpside_Ltm(iv *RepV) float64 {
	if p.InverseGrowth {
		return RoundX((iv.Ltm-p.Ltm)/iv.Ltm*100, 1)
	}
	return RoundX((p.Ltm-iv.Ltm)/iv.Ltm*100, 1)
}

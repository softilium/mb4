package pages

import (
	"context"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/db"
	d "github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/emitent"
)

type ri struct {
	V float64
	C string
}

type RV struct {
	Ytd       ri
	Ltm       ri
	AG_YtdAdj ri
	AG_Ltm    ri
}

func (t *RV) LoadV(r *cube.Cell, rv d.ReportValue) {
	q := r.GetRepV(rv)
	if q == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepV\n", rv)
		return
	}

	t.Ytd.V = q.Src
	t.Ytd.C = r.EvalChangeAsClass(rv, d.RVT_Src)

	t.Ltm.V = q.Ltm
	t.Ltm.C = r.EvalChangeAsClass(rv, d.RVT_Ltm)

	t.AG_YtdAdj.V = q.AGYtdAdj
	t.AG_YtdAdj.C = r.EvalChangeAsClass(rv, d.RVT_AG_YtdAdj)

	t.AG_Ltm.V = q.AGLtm
	t.AG_Ltm.C = r.EvalChangeAsClass(rv, d.RVT_AG_Ltm)

}

type RS struct {
	S  ri
	AG ri
}

func (t *RS) LoadS(r *cube.Cell, rv d.ReportValue) {
	q := r.GetRepS(rv)
	if q == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepS\n", rv)
		return
	}
	t.S.V, t.S.C = q.S, r.EvalChangeAsClass(rv, d.RVT_S)
	t.AG.V, t.AG.C = q.AG, r.EvalChangeAsClass(rv, d.RVT_AG)
}

type MS struct {
	S      ri
	AG     ri
	IND_S  ri
	Upside ri
}

func (t *MS) LoadMS(r *cube.Cell, rv d.ReportValue) {
	q := r.GetRepS(rv)
	if q == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepS\n", rv)
		return
	}

	qi := r.IndustryCell.GetRepS(rv)
	if qi == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepS\n", rv)
		return
	}

	t.S.V = q.S
	t.S.C = r.EvalChangeAsClass(rv, d.RVT_S)

	t.AG.V = q.AG
	t.AG.C = r.EvalChangeAsClass(rv, d.RVT_AG)

	t.IND_S.V = qi.S
	t.IND_S.C = r.IndustryCell.EvalChangeAsClass(rv, d.RVT_S)

	t.Upside.V = q.IndustryUpside
	t.Upside.C = r.EvalChangeAsClass(rv, d.RVT_IndUpside)

}

type MV struct {
	YtdAdj        ri
	AG_YtdAdj     ri
	IND_YtdAdj    ri
	Upside_YtdAdj ri
	Ltm           ri
	AG_Ltm        ri
	IND_Ltm       ri
	Upside_Ltm    ri
}

func (t *MV) LoadMV(r *cube.Cell, rv d.ReportValue) {

	q := r.GetRepV(rv)
	if q == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepS\n", rv)
		return
	}

	qi := r.IndustryCell.GetRepV(rv)
	if qi == nil {
		log.Panicf("ReportValue %v does not return value from cube.Cell.GetRepS\n", rv)
		return
	}

	t.YtdAdj.V = q.YtdAdj
	t.YtdAdj.C = r.EvalChangeAsClass(rv, d.RVT_Src)

	t.AG_YtdAdj.V = q.AGYtdAdj
	t.AG_YtdAdj.C = r.EvalChangeAsClass(rv, d.RVT_AG_YtdAdj)

	t.IND_YtdAdj.V = qi.YtdAdj
	t.IND_YtdAdj.C = ""

	t.Upside_YtdAdj.V = q.IndustryUpside_YtdAdj
	t.Upside_YtdAdj.C = r.EvalChangeAsClass(rv, d.RVT_IndUpside_YtdAdj)

	t.Ltm.V = q.Ltm
	t.Ltm.C = r.EvalChangeAsClass(rv, d.RVT_Ltm)

	t.AG_Ltm.V = q.AGLtm
	t.AG_Ltm.C = r.EvalChangeAsClass(rv, d.RVT_AG_Ltm)

	t.IND_Ltm.V = qi.Ltm
	t.IND_Ltm.C = ""

	t.Upside_Ltm.V = q.IndustryUpside_Ltm
	t.Upside_Ltm.C = r.EvalChangeAsClass(rv, d.RVT_IndUpside_Ltm)

}

func Report(w http.ResponseWriter, r *http.Request) {

	tickerid := r.URL.Query().Get("tickerid")
	if tickerid == "" {
		http.Error(w, "No ticker is specified", http.StatusNotFound)
		return
	}

	ystr := r.URL.Query().Get("y")
	y, err := strconv.ParseInt(ystr, 10, 0)
	HandleErr(err, w)

	qstr := r.URL.Query().Get("q")
	q, err := strconv.ParseInt(qstr, 10, 0)
	HandleErr(err, w)

	reps := cube.Market.GetReports2(tickerid)
	for _, r2 := range reps {

		if r2.ReportYear == int(y) && r2.ReportQuarter == int(q) {

			r3 := cube.Market.CellsByTickerByDate(tickerid, r2.ReportDate, cube.LookAhead)

			indRep := cube.Market.GetIndustryCell(r3.Quote.Edges.Ticker.Edges.Emitent.Edges.Industry.ID, r2.ReportDate)

			allreps := cube.Market.GetReports2(tickerid)
			sort.Slice(allreps, func(i, j int) bool {
				if allreps[i].ReportYear == allreps[j].ReportYear {
					return allreps[i].ReportQuarter > allreps[j].ReportQuarter
				} else {
					return allreps[i].ReportYear > allreps[j].ReportYear
				}
			})

			em, err := db.DB.Emitent.Query().
				Where(emitent.IDEQ(r3.Quote.Edges.Ticker.Edges.Emitent.ID)).
				WithTickers().
				Only(context.Background())
			HandleErr(err, w)

			si := LoadSessionStruct(r)
			pd := struct {
				SessionStruct

				R2 *cube.Report2

				//Pnl
				Revenue          RV
				Amortization     RV
				OperatingIncome  RV
				InterestIncome   RV
				InterestExpenses RV
				IncomeTax        RV
				NetIncome        RV
				OIBDA            RV
				EBITDA           RV

				//Cf
				Cash                  RS
				CurrentLiabilities    RS
				NonCurrentLiabilities RS
				NonControlling        RS
				Equity                RS
				Total                 RS
				NetDebt               RS
				Cap                   RS
				EV                    RS
				BookValue             RS

				//MultV
				EV_on_EBITDA      MV
				Debt_on_EBITDA    MV
				P_on_E            MV
				P_on_S            MV
				ROE               MV
				OIBDAMargin       MV
				EBITDAMargin      MV
				OperationalMargin MV
				NetMargin         MV

				//MultS
				P_on_BV MS
				DSI     MS

				Emitent  *ent.Emitent
				IR       *cube.Cell
				AllReps  []*cube.Report2
				TickerId string
			}{
				SessionStruct: si,
				R2:            r3.R2,
				Emitent:       em,
				IR:            indRep,
				AllReps:       allreps,
				TickerId:      tickerid,
			}

			pd.Revenue.LoadV(r3, d.RK_Revenue)
			pd.Amortization.LoadV(r3, d.RK_Amortization)
			pd.OperatingIncome.LoadV(r3, d.RK_OperatingIncome)
			pd.InterestIncome.LoadV(r3, d.RK_InterestIncome)
			pd.InterestExpenses.LoadV(r3, d.RK_InterestExpenses)
			pd.IncomeTax.LoadV(r3, d.RK_IncomeTax)
			pd.NetIncome.LoadV(r3, d.RK_NetIncome)
			pd.OIBDA.LoadV(r3, d.RK_OIBDA)
			pd.EBITDA.LoadV(r3, d.RK_EBITDA)

			pd.Cash.LoadS(r3, d.RK_Cash)
			pd.CurrentLiabilities.LoadS(r3, d.RK_CurrentLiabilities)
			pd.NonCurrentLiabilities.LoadS(r3, d.RK_NonCurrentLiabilities)
			pd.NonControlling.LoadS(r3, d.RK_NonCurrentLiabilities)
			pd.Equity.LoadS(r3, d.RK_Equity)
			pd.Total.LoadS(r3, d.RK_Total)
			pd.NetDebt.LoadS(r3, d.RK_NetDebt)
			pd.Cap.LoadS(r3, d.RK_Cap)
			pd.EV.LoadS(r3, d.RK_EV)
			pd.BookValue.LoadS(r3, d.RK_BookValue)

			pd.EV_on_EBITDA.LoadMV(r3, d.RK_EV_on_EBITDA)
			pd.Debt_on_EBITDA.LoadMV(r3, d.RK_Debt_on_EBITDA)
			pd.P_on_E.LoadMV(r3, d.RK_P_on_E)
			pd.P_on_S.LoadMV(r3, d.RK_P_on_S)
			pd.ROE.LoadMV(r3, d.RK_ROE)
			pd.OIBDAMargin.LoadMV(r3, d.RK_OIBDAMargin)
			pd.EBITDAMargin.LoadMV(r3, d.RK_EBITDAMargin)
			pd.OperationalMargin.LoadMV(r3, d.RK_OperationalMargin)
			pd.NetMargin.LoadMV(r3, d.RK_NetMargin)

			pd.P_on_BV.LoadMS(r3, d.RK_P_on_BV)
			pd.DSI.LoadMS(r3, d.RK_DSI)

			pd.Vue = false
			pd.Echarts = false

			tmpl, err := pongo2.FromCache("pages/report.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)
			HandleErr(err, w)
			return

		}
	}

}

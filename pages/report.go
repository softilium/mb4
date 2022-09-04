package pages

import (
	"context"
	"net/http"
	"sort"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/emitent"
)

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
			pageData := struct {
				SessionStruct
				R2       *cube.Report2
				R3       *cube.Cell
				Emitent  *ent.Emitent
				IR       *cube.Cell
				AllReps  []*cube.Report2
				TickerId string
			}{SessionStruct: si, R2: r2, R3: r3, Emitent: em, IR: indRep, AllReps: allreps, TickerId: tickerid}
			pageData.Vue = false
			pageData.Echarts = false

			tmpl, err := pongo2.FromCache("pages/report.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteWriter(pongo2.Context{"pd": pageData}, w)
			HandleErr(err, w)
			return

		}
	}

}

package pages

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
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

			si := LoadSessionStruct(r)
			pageData := struct {
				SessionStruct
				R2 *cube.Report2
				R3 *cube.Cell
			}{SessionStruct: si, R2: r2, R3: r3}
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

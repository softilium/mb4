package pages

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
)

func Ticker(w http.ResponseWriter, r *http.Request) {

	si := LoadSessionStruct(r)

	tickerId := r.URL.Query().Get("id")
	mode := r.URL.Query().Get("mode")

	if mode == "" {

		pageData := struct {
			SessionStruct
			TRI *cube.TickerRenderInfo
		}{SessionStruct: si, TRI: cube.Market.GetTickerRenderInfo(tickerId, false)}
		pageData.Vue = true
		pageData.Echarts = true

		tmpl, err := pongo2.FromCache("pages/ticker.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteWriter(pongo2.Context{"pd": pageData}, w)

		return
	}

	if mode == "candles" {
		tri := cube.Market.GetTickerRenderInfo(tickerId, true)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tri)
		return
	}

	if mode == "pnl" {

		pnlres := struct {
			Dates            []string
			Revenues         []float64
			InterestIncomes  []float64
			Ebitdas          []float64
			Amortizations    []float64
			InterestExpenses []float64
			Taxes            []float64
			Incomes          []float64
		}{}

		reps := GetYearReports(tickerId)
		pnlres.Dates = make([]string, len(reps))
		pnlres.Revenues = make([]float64, len(reps))
		pnlres.InterestIncomes = make([]float64, len(reps))
		pnlres.Ebitdas = make([]float64, len(reps))
		pnlres.Amortizations = make([]float64, len(reps))
		pnlres.InterestExpenses = make([]float64, len(reps))
		pnlres.Taxes = make([]float64, len(reps))
		pnlres.Incomes = make([]float64, len(reps))
		for idx, rep := range reps {

			if rep.ReportQuarter == 4 {
				pnlres.Dates[idx] = fmt.Sprintf("%v.Q%v", rep.ReportDate.Year(), rep.ReportQuarter)
			} else {
				pnlres.Dates[idx] = "LTM"
			}
			pnlres.Revenues[idx] = rep.YV[cube.RK2Revenue].V
			pnlres.InterestIncomes[idx] = rep.YV[cube.RK2InterestIncome].V
			pnlres.Ebitdas[idx] = rep.YV[cube.RK2EBITDA].V
			pnlres.Amortizations[idx] = rep.YV[cube.RK2Amortization].V
			pnlres.InterestExpenses[idx] = rep.YV[cube.RK2InterestExpenses].V
			pnlres.Taxes[idx] = rep.YV[cube.RK2IncomeTax].V
			pnlres.Incomes[idx] = rep.YV[cube.RK2NetIncome].V
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pnlres)
		return

	}

	if mode == "cf" {

		cfres := struct {
			Dates     []string
			Cash      []float64
			Debt      []float64
			Equity    []float64
			MCap      []float64
			BookValue []float64
		}{}

		reps := GetYearReports(tickerId)
		cfres.Dates = make([]string, len(reps))
		cfres.Cash = make([]float64, len(reps))
		cfres.Debt = make([]float64, len(reps))
		cfres.Equity = make([]float64, len(reps))
		cfres.MCap = make([]float64, len(reps))
		cfres.BookValue = make([]float64, len(reps))
		for idx, rep := range reps {

			if rep.ReportQuarter == 4 {
				cfres.Dates[idx] = fmt.Sprintf("%v.Q%v", rep.ReportDate.Year(), rep.ReportQuarter)
			} else {
				cfres.Dates[idx] = "LTM"
			}
			cfres.Cash[idx] = rep.SV[cube.RK2Cash].V
			cfres.Debt[idx] = rep.SV[cube.RK2NetDebt].V
			cfres.Equity[idx] = rep.SV[cube.RK2Equity].V
			cfres.MCap[idx] = cube.Market.CellsByTickerByDate(tickerId, rep.ReportDate).R3[cube.RK3Cap].V
			cfres.BookValue[idx] = cube.Market.CellsByTickerByDate(tickerId, rep.ReportDate).R3[cube.RK3BookValue].V
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cfres)
		return

	}

	http.Error(w, "", http.StatusBadRequest)

}

func GetYearReports(tickerId string) []*cube.Report2 {

	allreps := cube.Market.GetReports2(tickerId)
	reps := make([]*cube.Report2, 0, len(allreps)/3)
	for _, rep := range allreps {
		if rep.ReportQuarter == 4 {
			reps = append(reps, rep)
		}
	}
	if allreps[len(allreps)-1].ReportQuarter != 4 {
		reps = append(reps, allreps[len(allreps)-1])
	}
	return reps
}

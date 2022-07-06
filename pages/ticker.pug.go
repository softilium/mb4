package pages

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		templates["ticker"].Execute(w, pageData)
		return
	}

	if mode == "candles" {
		tri := cube.Market.GetTickerRenderInfo(tickerId, true)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tri)
		return
	}

	if mode == "pnl" {

		pnlResult := struct {
			Dates            []string
			Revenues         []float64
			InterestIncomes  []float64
			Ebitdas          []float64
			Amortizations    []float64
			InterestExpenses []float64
			Taxes            []float64
			Incomes          []float64
		}{}
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
		pnlResult.Dates = make([]string, len(reps))
		pnlResult.Revenues = make([]float64, len(reps))
		pnlResult.InterestIncomes = make([]float64, len(reps))
		pnlResult.Ebitdas = make([]float64, len(reps))
		pnlResult.Amortizations = make([]float64, len(reps))
		pnlResult.InterestExpenses = make([]float64, len(reps))
		pnlResult.Taxes = make([]float64, len(reps))
		pnlResult.Incomes = make([]float64, len(reps))
		for idx, rep := range reps {

			if rep.ReportQuarter == 4 {
				pnlResult.Dates[idx] = fmt.Sprintf("%v.Q%v", rep.ReportDate.Year(), rep.ReportQuarter)
			} else {
				pnlResult.Dates[idx] = "LTM"
			}
			pnlResult.Revenues[idx] = rep.YV[cube.RK2Revenue].Ytd
			pnlResult.InterestIncomes[idx] = rep.YV[cube.RK2InterestIncome].Ytd
			pnlResult.Ebitdas[idx] = rep.YV[cube.RK2EBITDA].Ytd
			pnlResult.Amortizations[idx] = rep.YV[cube.RK2Amortization].Ytd
			pnlResult.InterestExpenses[idx] = rep.YV[cube.RK2InterestExpenses].Ytd
			pnlResult.Taxes[idx] = rep.YV[cube.RK2IncomeTax].Ytd
			pnlResult.Incomes[idx] = rep.YV[cube.RK2NetIncome].Ytd
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pnlResult)
		return

	}

	http.Error(w, "", http.StatusBadRequest)

}

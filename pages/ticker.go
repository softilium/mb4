package pages

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
)

func Ticker(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		{
			tickerId := r.URL.Query().Get("id")
			mode := r.URL.Query().Get("mode")

			switch mode {
			case "": //get page
				{
					si := LoadSessionStruct(r)
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

				}
			case "candles":
				{
					tri := cube.Market.GetTickerRenderInfo(tickerId, true)

					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(tri)
					if err != nil {
						HandleErr(err, w)
					}

				}
			case "pnl":
				{
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
							pnlres.Dates[idx] = fmt.Sprintf("%v", rep.ReportDate.Year())
							pnlres.Revenues[idx] = rep.Revenue.V
							pnlres.InterestIncomes[idx] = rep.InterestIncome.V
							pnlres.Ebitdas[idx] = rep.EBITDA.V
							pnlres.Amortizations[idx] = rep.Amortization.V
							pnlres.InterestExpenses[idx] = rep.InterestExpenses.V
							pnlres.Taxes[idx] = rep.IncomeTax.V
							pnlres.Incomes[idx] = rep.NetIncome.V
						} else {
							pnlres.Dates[idx] = "LTM"
							pnlres.Revenues[idx] = rep.Revenue.Ltm
							pnlres.InterestIncomes[idx] = rep.InterestIncome.Ltm
							pnlres.Ebitdas[idx] = rep.EBITDA.Ltm
							pnlres.Amortizations[idx] = rep.Amortization.Ltm
							pnlres.InterestExpenses[idx] = rep.InterestExpenses.Ltm
							pnlres.Taxes[idx] = rep.IncomeTax.Ltm
							pnlres.Incomes[idx] = rep.NetIncome.Ltm
						}
					}

					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(pnlres)
					if err != nil {
						HandleErr(err, w)
					}
				}
			case "cf":
				{
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
							cfres.Dates[idx] = fmt.Sprintf("%v", rep.ReportDate.Year())
						} else {
							cfres.Dates[idx] = "LTM"
						}
						cfres.Cash[idx] = rep.Cash.V
						cfres.Debt[idx] = rep.NetDebt.V
						cfres.Equity[idx] = rep.Equity.V
						cfres.MCap[idx] = cube.Market.CellsByTickerByDate(tickerId, rep.ReportDate).Cap.V
						cfres.BookValue[idx] = cube.Market.CellsByTickerByDate(tickerId, rep.ReportDate).BookValue.V
					}

					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(cfres)
					if err != nil {
						HandleErr(err, w)
					}
				}
			case "mult":
				{

					multres := struct {
						Dates           []string
						NetMargin       []float64
						NetMarginInd    []float64
						EBITDAMargin    []float64
						EBITDAMarginInd []float64
					}{}

					reps := GetYearReports(tickerId)
					multres.Dates = make([]string, len(reps))
					multres.NetMargin = make([]float64, len(reps))
					multres.NetMarginInd = make([]float64, len(reps))
					multres.EBITDAMargin = make([]float64, len(reps))
					multres.EBITDAMarginInd = make([]float64, len(reps))
					for idx, rep := range reps {

						cell := cube.Market.CellsByTickerByDate(tickerId, rep.ReportDate)
						indCell := cube.Market.GetIndustryCell(cell.Industry.ID, cell.D)

						if rep.ReportQuarter == 4 {
							multres.Dates[idx] = fmt.Sprintf("%v", rep.ReportDate.Year())
							multres.NetMargin[idx] = rep.NetMargin.V
							multres.NetMarginInd[idx] = indCell.R2.NetMargin.V
							multres.EBITDAMargin[idx] = rep.EBITDAMargin.V
							multres.EBITDAMarginInd[idx] = indCell.R2.EBITDAMargin.V
						} else {
							multres.Dates[idx] = "LTM"
							multres.EBITDAMargin[idx] = rep.EBITDAMargin.Ltm
							multres.EBITDAMarginInd[idx] = indCell.R2.EBITDAMargin.Ltm
						}
					}

					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(multres)
					if err != nil {
						HandleErr(err, w)
					}

				}
			default:
				http.Error(w, "Mode not allowed", http.StatusBadRequest)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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

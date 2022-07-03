package pages

import (
	"fmt"
	"net/http"

	"github.com/softilium/mb4/cube"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	topYield5Y := cube.Market.TopDivYields5Y(20)
	topDSI := cube.Market.TopDSI(20)
	topFallen := cube.Market.TopFallenRaise(20, false)
	topRaise := cube.Market.TopFallenRaise(20, true)

	type TopItem struct {
		Ticker string
		Descr  string
		V      string
		V1     string
	}

	type pageDataStruct struct {
		SessionStruct
		TopY      []TopItem
		TopDSI    []TopItem
		TopFallen []TopItem
		TopRaise  []TopItem
	}

	pd := pageDataStruct{}
	pd.SessionStruct = LoadSessionStruct(r)
	pd.TopY = make([]TopItem, len(topYield5Y))
	for k, v := range topYield5Y {
		pd.TopY[k].Ticker = "/ticker/" + v.Quote.Edges.Ticker.ID
		pd.TopY[k].Descr = v.Quote.Edges.Ticker.Descr
		pd.TopY[k].V = fmt.Sprintf("%.1f", v.DivYield5Y)
		pd.TopY[k].V1 = fmt.Sprintf("%.1f", v.DSI)
	}
	pd.TopDSI = make([]TopItem, len(topDSI))
	for k, v := range topDSI {
		pd.TopDSI[k].Ticker = "/ticker/" + v.Quote.Edges.Ticker.ID
		pd.TopDSI[k].Descr = v.Quote.Edges.Ticker.Descr
		pd.TopDSI[k].V = fmt.Sprintf("%.1f", v.DivYield5Y)
		pd.TopDSI[k].V1 = fmt.Sprintf("%.1f", v.DSI)
	}
	pd.TopFallen = make([]TopItem, len(topFallen))
	for k, v := range topFallen {
		pd.TopFallen[k].Ticker = "/ticker/" + v.Cell.Quote.Edges.Ticker.ID
		pd.TopFallen[k].Descr = v.Cell.Quote.Edges.Ticker.Descr
		pd.TopFallen[k].V = fmt.Sprintf("%.1f", -v.PercentPriceChange)
	}
	pd.TopRaise = make([]TopItem, len(topRaise))
	for k, v := range topRaise {
		pd.TopRaise[k].Ticker = "/ticker/" + v.Cell.Quote.Edges.Ticker.ID
		pd.TopRaise[k].Descr = v.Cell.Quote.Edges.Ticker.Descr
		pd.TopRaise[k].V = fmt.Sprintf("%.1f", v.PercentPriceChange)
	}

	templates["index"].Execute(w, pd)

}

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

	topYield5Y, err := cube.Market.TopDivYields5Y(20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	topDSI, err := cube.Market.TopDSI(20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	type TopItem struct {
		Ticker string
		Descr  string
		Yield  string
		DSI    string
	}

	type pageDataStruct struct {
		SessionStruct
		TopY   []TopItem
		TopDSI []TopItem
	}

	pd := pageDataStruct{}
	pd.SessionStruct = LoadSessionStruct(r)
	pd.TopY = make([]TopItem, len(topYield5Y))
	for k, v := range topYield5Y {
		pd.TopY[k].Ticker = "/ticker/" + v.Quote.Edges.Ticker.ID
		pd.TopY[k].Descr = v.Quote.Edges.Ticker.Descr
		pd.TopY[k].Yield = fmt.Sprintf("%.1f", v.DivYield5Y)
		pd.TopY[k].DSI = fmt.Sprintf("%.1f", v.DSI)
	}
	pd.TopDSI = make([]TopItem, len(topDSI))
	for k, v := range topDSI {
		pd.TopDSI[k].Ticker = "/ticker/" + v.Quote.Edges.Ticker.ID
		pd.TopDSI[k].Descr = v.Quote.Edges.Ticker.Descr
		pd.TopDSI[k].Yield = fmt.Sprintf("%.1f", v.DivYield5Y)
		pd.TopDSI[k].DSI = fmt.Sprintf("%.1f", v.DSI)
	}

	templates["index"].Execute(w, pd)

}

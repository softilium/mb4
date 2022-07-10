package pages

import (
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	topxx := 25

	topYield5Y := cube.Market.TopDivYields5Y(topxx)
	topDSI := cube.Market.TopDSI(topxx)
	topFallen := cube.Market.TopFallenRaise(topxx, false)
	topRaise := cube.Market.TopFallenRaise(topxx, true)

	type TopItem struct {
		Ticker string
		Star   bool
		Descr  string
		V      string
		V1     string
	}

	type pageDataStruct struct {
		SessionStruct
		TopY      []*TopItem
		TopDSI    []*TopItem
		TopFallen []*TopItem
		TopRaise  []*TopItem
	}

	pd := pageDataStruct{SessionStruct: LoadSessionStruct(r)}
	pd.TopY = make([]*TopItem, len(topYield5Y))
	for k, v := range topYield5Y {
		pd.TopY[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Quote.Edges.Ticker.ID,
			Descr:  v.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", v.DivYield5Y),
			V1:     fmt.Sprintf("%.1f", v.DSI),
		}
	}
	pd.TopDSI = make([]*TopItem, len(topDSI))
	for k, v := range topDSI {
		pd.TopDSI[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Quote.Edges.Ticker.ID,
			Descr:  v.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", v.DivYield5Y),
			V1:     fmt.Sprintf("%.1f", v.DSI),
		}
	}
	pd.TopFallen = make([]*TopItem, len(topFallen))
	for k, v := range topFallen {
		pd.TopFallen[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Cell.Quote.Edges.Ticker.ID,
			Descr:  v.Cell.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", -v.PercentPriceChange),
		}
	}
	pd.TopRaise = make([]*TopItem, len(topRaise))
	for k, v := range topRaise {
		pd.TopRaise[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Cell.Quote.Edges.Ticker.ID,
			Descr:  v.Cell.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", v.PercentPriceChange),
		}
	}
	// count stars
	starmap := make(map[string]int)
	for _, ti := range pd.TopY {
		starmap[ti.Ticker] = starmap[ti.Ticker] + 1
	}
	for _, ti := range pd.TopDSI {
		starmap[ti.Ticker] = starmap[ti.Ticker] + 1
	}
	for _, ti := range pd.TopFallen {
		starmap[ti.Ticker] = starmap[ti.Ticker] + 1
	}

	// set stars
	for _, ti := range pd.TopY {
		if starmap[ti.Ticker] > 2 {
			ti.Star = true
		}
	}
	for _, ti := range pd.TopDSI {
		if starmap[ti.Ticker] > 2 {
			ti.Star = true
		}
	}
	for _, ti := range pd.TopFallen {
		if starmap[ti.Ticker] > 2 {
			ti.Star = true
		}
	}

	tmpl, err := pongo2.FromCache("pages/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)

}

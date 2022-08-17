package pages

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/db"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	auth := LoadSessionStruct(r)

	topxx := 20
	if auth.Authenticated {
		topxx = auth.user.HowManyTickersOnHomepage
	}

	qhm := r.URL.Query().Get("hm")
	if qhm != "" {
		t, err := strconv.ParseInt(qhm, 10, 64)
		if err == nil {
			topxx = int(t)

			if auth.Authenticated {
				_, err := db.DB.User.UpdateOneID(auth.UserID).SetHowManyTickersOnHomepage(topxx).Save(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}

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
		HM        int
	}

	pd := pageDataStruct{SessionStruct: auth, HM: topxx}
	pd.TopY = make([]*TopItem, len(topYield5Y))
	for k, v := range topYield5Y {
		pd.TopY[k] = &TopItem{
			Ticker: "/ticker?id=" + v.TickerId(),
			Descr:  v.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", v.DivYield5Y.V),
			V1:     fmt.Sprintf("%.1f", v.DSI.V),
		}
	}
	pd.TopDSI = make([]*TopItem, len(topDSI))
	for k, v := range topDSI {
		pd.TopDSI[k] = &TopItem{
			Ticker: "/ticker?id=" + v.TickerId(),
			Descr:  v.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", v.DivYield5Y.V),
			V1:     fmt.Sprintf("%.1f", v.DSI.V),
		}
	}
	pd.TopFallen = make([]*TopItem, len(topFallen))
	for k, v := range topFallen {
		pd.TopFallen[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Cell.TickerId(),
			Descr:  v.Cell.Quote.Edges.Ticker.Descr,
			V:      fmt.Sprintf("%.1f", -v.PercentPriceChange),
		}
	}
	pd.TopRaise = make([]*TopItem, len(topRaise))
	for k, v := range topRaise {
		pd.TopRaise[k] = &TopItem{
			Ticker: "/ticker?id=" + v.Cell.TickerId(),
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

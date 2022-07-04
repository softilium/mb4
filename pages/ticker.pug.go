package pages

import (
	"encoding/json"
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

	if mode == "json" {
		tri := cube.Market.GetTickerRenderInfo(tickerId, true)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tri)
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

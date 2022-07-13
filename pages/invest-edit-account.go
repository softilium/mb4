package pages

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
)

func InvestAccount(w http.ResponseWriter, r *http.Request) {

	s := LoadSessionStruct(r)
	if !s.Authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	pd := struct {
		SessionStruct
		AccId string
	}{SessionStruct: s, AccId: r.URL.Query().Get("id")}
	pd.Vue = true
	pd.Echarts = true

	tmpl, err := pongo2.FromCache("pages/invest-edit-account.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)

}

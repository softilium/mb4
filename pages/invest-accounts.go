package pages

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
)

func InvestAccounts(w http.ResponseWriter, r *http.Request) {

	pd := LoadSessionStruct(r)
	if !pd.Authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	pd.Vue = true
	pd.Echarts = true

	tmpl, err := pongo2.FromCache("pages/invest-accounts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)

}

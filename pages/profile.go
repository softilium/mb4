package pages

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	pd := LoadSessionStruct(r)
	pd.Vue = true
	if !pd.Authenticated {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	tmpl, err := pongo2.FromCache("pages/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)
	HandleErr(err, w)

}

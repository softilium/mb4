package pages

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
)

func Login(w http.ResponseWriter, r *http.Request) {

	tmpl, err := pongo2.FromCache("pages/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteWriter(pongo2.Context{}, w)

}
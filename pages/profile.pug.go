package pages

import (
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	si := loadSessionStruct(r)
	if !si.Authenticated {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	templates["profile"].Execute(w, si)

}

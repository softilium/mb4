package pages

import (
	"net/http"
)

func InvestAccount(w http.ResponseWriter, r *http.Request) {

	s := LoadSessionStruct(r)
	if !s.Authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	pagedata := struct {
		SessionStruct
		AccId string
	}{}
	pagedata.SessionStruct = s
	pagedata.AccId = r.URL.Query().Get("id")

	templates["invest-edit-account"].Execute(w, pagedata)

}

package pages

import (
	"net/http"
)

func InvestAccounts(w http.ResponseWriter, r *http.Request) {

	s := LoadSessionStruct(r)
	if !s.Authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	templates["invest-accounts"].Execute(w, LoadSessionStruct(r))

}

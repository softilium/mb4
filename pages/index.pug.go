package pages

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	templates["index"].Execute(w, loadSessionStruct(r))

}

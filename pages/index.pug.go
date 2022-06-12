package pages

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	templates["index"].Execute(w, loadAuthInfo(r))

}

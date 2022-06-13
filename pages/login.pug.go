package pages

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	templates["login"].Execute(w, nil)

}

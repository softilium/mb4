package api

import (
	"net/http"

	"log"

	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/pages"
)

func Cube(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)

	if !session.UserIsAdmin {
		http.Error(w, "", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		err := cube.Market.LoadCube()
		if err != nil {
			pages.HandleErr(err, w)
		}
		log.Printf("Market cube reloaded\n")
	}

}

package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/industry"
	"github.com/softilium/mb4/pages"
)

func Industries(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)

	if r.Method == http.MethodGet {

		id := r.URL.Query().Get("id")
		if len(id) == 0 {

			res, err := db.DB.Industry.Query().All(context.Background())
			handleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			handleErr(err, w)
			return

		} else {

			res, err := db.DB.Industry.Query().Where(industry.IDEQ(id)).All(context.Background())
			handleErr(err, w)

			if len(res) == 0 {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res[0])
			handleErr(err, w)
			return

		}

	}

	if r.Method == http.MethodPost {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		buf := ent.Industry{}
		err := json.NewDecoder(r.Body).Decode(&buf)
		handleErr(err, w)

		//deleteIndustry(buf.ID, w)

		_, err = db.DB.Industry.Create().SetID(buf.ID).SetDescr(buf.Descr).Save(context.Background())
		handleErr(err, w)
		return

	}

	if r.Method == http.MethodDelete {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		deleteIndustry(id, w)
		return

	}

}

func deleteIndustry(id string, w http.ResponseWriter) {

	_, err := db.DB.Industry.Delete().Where(industry.IDEQ(id)).Exec(context.Background())
	handleErr(err, w)

}
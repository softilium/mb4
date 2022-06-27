package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/industry"
)

func Industries(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		id := r.URL.Query().Get("id")
		if len(id) == 0 {

			res, err := db.DB.Industry.Query().All(context.Background())
			handleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			handleErr(err, w)

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

		}

	}

	if r.Method == http.MethodPost {

		buf := ent.Industry{}
		err := json.NewDecoder(r.Body).Decode(&buf)
		handleErr(err, w)

		_, err = db.DB.Industry.Create().SetID(buf.ID).SetDescr(buf.Descr).Save(context.Background())
		handleErr(err, w)

	}

	if r.Method == http.MethodDelete {

		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := db.DB.Industry.DeleteOneID(id).Exec(context.Background())
		handleErr(err, w)

	}

}

package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/ticker"
	"github.com/softilium/mb4/pages"
)

func Tickers(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)

	deleteOne := func(id string, w http.ResponseWriter) {
		_, err := db.DB.Ticker.Delete().Where(ticker.IDEQ(id)).Exec(context.Background())
		handleErr(err, w)
	}

	id := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {

		if len(id) == 0 {

			res, err := db.DB.Ticker.Query().All(context.Background())
			handleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			handleErr(err, w)
			return

		} else {

			res, err := db.DB.Ticker.Query().Where(ticker.IDEQ(id)).All(context.Background())
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

		buf := ent.Ticker{}
		err := json.NewDecoder(r.Body).Decode(&buf)
		handleErr(err, w)

		//deleteOne(buf.ID, w)

		_, err = db.DB.Ticker.Create().
			SetID(buf.ID).
			SetDescr(buf.Descr).
			SetKind(buf.Kind).
			SetEmitentID(buf.Edges.Emitent.ID).
			Save(context.Background())
		handleErr(err, w)
		return

	}

	if r.Method == http.MethodDelete {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		deleteOne(id, w)
		return

	}

}

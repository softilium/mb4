package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/quote"
	"github.com/softilium/mb4/pages"
)

func Quotes(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)

	deleteOne := func(id xid.ID, w http.ResponseWriter) {
		_, err := db.DB.Quote.Delete().Where(quote.IDEQ(id)).Exec(context.Background())
		pages.HandleErr(err, w)
	}

	id := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {

		if len(id) == 0 {

			res, err := db.DB.Quote.Query().All(context.Background())
			pages.HandleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			pages.HandleErr(err, w)
			return

		} else {

			xid, err := xid.FromString(id)
			pages.HandleErr(err, w)
			res, err := db.DB.Quote.Query().Where(quote.IDEQ(xid)).All(context.Background())
			pages.HandleErr(err, w)

			if len(res) == 0 {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res[0])
			pages.HandleErr(err, w)
			return

		}

	}

	if r.Method == http.MethodPost {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		buf := []ent.Quote{}
		err := json.NewDecoder(r.Body).Decode(&buf)
		pages.HandleErr(err, w)

		for _, v := range buf {
			v.ID = xid.New()
			//deleteOne(v.ID, w)
			_, err = db.DB.Quote.Create().
				SetTickerID(v.Edges.Ticker.ID).
				SetID(v.ID).
				SetD(v.D).
				SetO(v.O).
				SetC(v.C).
				SetH(v.H).
				SetL(v.L).
				SetV(v.V).
				Save(context.Background())
			pages.HandleErr(err, w)
		}

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
		xid, err := xid.FromString(id)
		pages.HandleErr(err, w)
		deleteOne(xid, w)
		return

	}

}

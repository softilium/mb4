package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/pages"
)

func Emitents(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)

	deleteOne := func(id xid.ID, w http.ResponseWriter) {

		_, err := db.DB.Emitent.Delete().Where(emitent.IDEQ(id)).Exec(context.Background())
		handleErr(err, w)
	}

	id := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {

		if len(id) == 0 {

			res, err := db.DB.Emitent.Query().All(context.Background())
			handleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			handleErr(err, w)
			return

		} else {

			xid, err := xid.FromString(id)
			handleErr(err, w)
			res, err := db.DB.Emitent.Query().Where(emitent.IDEQ(xid)).All(context.Background())
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

		buf := ent.Emitent{ID: xid.New()}
		err := json.NewDecoder(r.Body).Decode(&buf)
		handleErr(err, w)

		//deleteOne(buf.ID, w)

		newdata, err := db.DB.Emitent.Create().
			SetID(buf.ID).
			SetDescr(buf.Descr).
			SetIndustryID(buf.Edges.Industry.ID).
			Save(context.Background())
		handleErr(err, w)

		w.Write([]byte(newdata.ID.String()))
		return

	}

	if r.Method == http.MethodDelete {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		xid, err := xid.FromString(id)
		handleErr(err, w)
		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		deleteOne(xid, w)
		return

	}

}
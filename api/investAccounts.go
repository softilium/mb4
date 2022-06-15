package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/investaccountcashflow"
	"github.com/softilium/mb4/ent/investaccountvaluation"
	"github.com/softilium/mb4/ent/user"
	"github.com/softilium/mb4/pages"
)

func ApiInvestAccounts(w http.ResponseWriter, r *http.Request) {

	currentSession := pages.LoadSessionStruct(r)
	if !currentSession.Authenticated || !currentSession.User.Admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if r.Method == http.MethodGet {

	// 	//TODO: get all/one invest accounts for user

	// }

	if r.Method == http.MethodDelete {

		id, err := xid.FromString(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		db.DB.InvestAccountValuation.Delete().
			Where(investaccountvaluation.HasOwnerWith(investaccount.IDEQ(id))).
			Exec(context.Background())

		db.DB.InvestAccountCashflow.Delete().
			Where(investaccountcashflow.HasOwnerWith(investaccount.IDEQ(id))).
			Exec(context.Background())

		db.DB.InvestAccount.Delete().
			Where(investaccount.And(investaccount.ID(id), investaccount.HasOwnerWith(user.IDEQ(currentSession.User.ID)))).
			Exec(context.Background())
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {

		var p []*ent.InvestAccount

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, pone := range p {
			existings, err := db.DB.InvestAccount.Query().
				Where(investaccount.And(
					investaccount.DescrEQ(pone.Descr),
					investaccount.HasOwnerWith(user.IDEQ(currentSession.User.ID)))).
				All(context.Background())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, existing := range existings {

				_, err = db.DB.InvestAccountValuation.Delete().
					Where(investaccountvaluation.HasOwnerWith(investaccount.IDEQ(existing.ID))).
					Exec(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				_, err = db.DB.InvestAccountCashflow.Delete().
					Where(investaccountcashflow.HasOwnerWith(investaccount.IDEQ(existing.ID))).
					Exec(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = db.DB.InvestAccount.DeleteOneID(existing.ID).Exec(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

		}

		for _, v := range p {

			ins, err := db.DB.InvestAccount.Create().
				SetDescr(v.Descr).
				SetOwnerID(currentSession.User.ID).
				Save(context.Background())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for _, vv := range v.Edges.Valuations {
				_, err = db.DB.InvestAccountValuation.Create().
					SetOwnerID(ins.ID).
					SetRecDate(vv.RecDate).
					SetValue(vv.Value).
					Save(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			}

			for _, vv := range v.Edges.Cashflows {
				_, err = db.DB.InvestAccountCashflow.Create().
					SetOwnerID(ins.ID).
					SetRecDate(vv.RecDate).
					SetQty(vv.Qty).
					Save(context.Background())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			}

		}

		w.WriteHeader(http.StatusOK)
		return
	}

}

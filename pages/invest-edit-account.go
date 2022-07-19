package pages

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
)

func InvestAccount(w http.ResponseWriter, r *http.Request) {

	s := LoadSessionStruct(r)
	if !s.Authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	mode := r.URL.Query().Get("mode")
	switch r.Method {
	case http.MethodGet:
		{
			pd := struct {
				SessionStruct
				AccId string
			}{SessionStruct: s, AccId: r.URL.Query().Get("id")}
			pd.Vue = true
			pd.Echarts = true

			tmpl, err := pongo2.FromCache("pages/invest-edit-account.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.ExecuteWriter(pongo2.Context{"pd": pd}, w)
		}
	case http.MethodPut:
		{
			xid, err := xid.FromString(r.URL.Query().Get("id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			switch mode {
			case "val":
				{
					par := ent.InvestAccountValuation{}
					err := json.NewDecoder(r.Body).Decode(&par)
					HandleErr(err, w)

					err = db.DB.InvestAccountValuation.UpdateOneID(xid).
						SetRecDate(par.RecDate).
						SetValue(par.Value).
						Exec(context.Background())
					HandleErr(err, w)
				}
			case "cf":
				{
					par := ent.InvestAccountCashflow{}
					err := json.NewDecoder(r.Body).Decode(&par)
					HandleErr(err, w)

					err = db.DB.InvestAccountCashflow.UpdateOneID(xid).
						SetRecDate(par.RecDate).
						SetQty(par.Qty).
						Exec(context.Background())
					HandleErr(err, w)
				}
			default:
				http.Error(w, "mode not allowed", http.StatusInternalServerError)
			}
		}
	case http.MethodDelete:
		{
			xid, err := xid.FromString(r.URL.Query().Get("id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			switch mode {
			case "cf":
				{
					err = db.DB.InvestAccountCashflow.DeleteOneID(xid).Exec(context.Background())
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			case "val":
				{
					err = db.DB.InvestAccountValuation.DeleteOneID(xid).Exec(context.Background())
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			default:
				http.Error(w, "mode not allowed", http.StatusInternalServerError)
			}
		}
	case http.MethodPost:
		{
			switch mode {
			case "cf":
				{
					owner, err := xid.FromString(r.URL.Query().Get("owner"))
					HandleErr(err, w)

					par := ent.InvestAccountCashflow{}
					err = json.NewDecoder(r.Body).Decode(&par)
					HandleErr(err, w)

					newObj, err := db.DB.InvestAccountCashflow.Create().
						SetID(xid.New()).
						SetRecDate(par.RecDate).
						SetQty(par.Qty).
						SetOwner(db.DB.InvestAccount.GetX(context.Background(), owner)).
						Save(context.Background())
					HandleErr(err, w)
					w.Write([]byte(newObj.ID.String()))
				}
			case "val":
				{
					owner, err := xid.FromString(r.URL.Query().Get("owner"))
					HandleErr(err, w)

					par := ent.InvestAccountValuation{}
					err = json.NewDecoder(r.Body).Decode(&par)
					HandleErr(err, w)

					newObj, err := db.DB.InvestAccountValuation.Create().
						SetID(xid.New()).
						SetRecDate(par.RecDate).
						SetValue(par.Value).
						SetOwner(db.DB.InvestAccount.GetX(context.Background(), owner)).
						Save(context.Background())
					HandleErr(err, w)
					w.Write([]byte(newObj.ID.String()))
				}
			default:
				http.Error(w, "mode not allowed", http.StatusInternalServerError)
			}

		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

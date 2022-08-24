package pages

import (
	"context"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/strategy"
	"github.com/softilium/mb4/ent/user"
)

func Strategies(w http.ResponseWriter, r *http.Request) {

	si := LoadSessionStruct(r)
	pageData := struct {
		SessionStruct
		Items []*ent.Strategy
	}{SessionStruct: si}
	pageData.Vue = false
	pageData.Echarts = false

	switch r.Method {
	case http.MethodGet:
		{
			var err error

			pageData.Items, err = db.DB.Strategy.
				Query().
				Where(strategy.HasUserWith(user.IDEQ(si.user.ID))).
				All(context.Background())
			HandleErr(err, w)

			tmpl, err := pongo2.FromCache("pages/strategies.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.ExecuteWriter(pongo2.Context{"pd": pageData}, w)
			HandleErr(err, w)
		}
	case http.MethodPost:
		{
			descr := r.FormValue("descr")
			_, err := db.DB.Strategy.Create().
				SetDescr(descr).
				SetUser(si.user).
				SetWeekRefillAmount(1000).
				SetStartAmount(1000).
				Save(context.Background())
			HandleErr(err, w)
			http.Redirect(w, r, "/strategies", http.StatusSeeOther)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

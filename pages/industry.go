package pages

import (
	"context"
	"net/http"
	"sort"

	"github.com/flosch/pongo2/v6"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Industry(w http.ResponseWriter, r *http.Request) {

	iid := r.URL.Query().Get("id")
	if iid == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	_, err := db.DB.Industry.Get(context.Background(), iid)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	ld := cube.Market.LastDate()

	type emitentRow struct {
		Ticker     *ent.Ticker
		Cap        float64
		EV         float64
		EBITDA     float64
		EBITDAGrow float64
	}

	rows := make([]*emitentRow, 0)
	for _, t := range cube.Market.GetAllTickers() {
		if t.Edges.Emitent.Edges.Industry.ID != iid {
			continue
		}

		q := cube.Market.CellsByTickerByDate(t.ID, ld, true)
		if q.R2 == nil {
			rows = append(rows, &emitentRow{Ticker: t, Cap: q.Cap.Ltm})
		} else {
			rows = append(rows, &emitentRow{Ticker: t, Cap: q.Cap.V, EV: q.R2.EV.V, EBITDA: q.R2.EBITDA.V, EBITDAGrow: q.R2.EBITDA.AG})
		}

	}

	p := message.NewPrinter(language.Russian)
	p.Printf("%d\n", 1000)

	sort.Slice(rows, func(i, j int) bool { return rows[i].Cap > rows[j].Cap })

	si := LoadSessionStruct(r)
	pageData := struct {
		SessionStruct
		Emitents []*emitentRow
	}{SessionStruct: si, Emitents: rows}
	pageData.Vue = true
	pageData.Echarts = true

	tmpl, err := pongo2.FromCache("pages/industry.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteWriter(pongo2.Context{"pd": pageData}, w)

}

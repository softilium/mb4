package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/investaccountcashflow"
	"github.com/softilium/mb4/ent/investaccountvaluation"
	"github.com/softilium/mb4/ent/user"
	"github.com/softilium/mb4/pages"
)

//InvestAccounts handles the /api/invest-accounts.* endpoint.
func InvestAccounts(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)
	if !session.Authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	parMode := r.URL.Query().Get("mode")
	parID := r.URL.Query().Get("id")

	if r.Method == http.MethodGet {
		if parMode == "weekflow" {
			// /api/invest-accounts?mode=weekmode
			handleAccsWeekflow(w, r, session)
		} else if parID != "" {
			// /api/invent-acconts?id=123123
			handleAccsGetOne(w, r, session, parID)
		} else {
			// api/invent-accounts
			handleAccsGetList(session, w)
		}
	}

	if !session.User.Admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// methods is for exchange with backoffice (for admin only)
	if r.Method == http.MethodDelete {
		handleAccsDelete(r, w, session)
	}

	if r.Method == http.MethodPost {
		handleAccsPost(r, w, session)
	}

	if r.Method == http.MethodPut {
		handleAccsPut(r, w, session)
	}

}

func InvestAccountValuations(w http.ResponseWriter, r *http.Request) {

	var id xid.ID
	var err error

	session := pages.LoadSessionStruct(r)
	if !session.Authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodDelete && r.Method != http.MethodPut && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodDelete && r.Method != http.MethodPut {
		id, err = xid.FromString(r.URL.Query().Get("id"))
		handleErr(err, w)
	}

	accXids, err := session.GetInvestAccountXids()
	handleErr(err, w)

	if r.Method == http.MethodDelete {

		cnt, err := db.DB.InvestAccountValuation.Delete().Where(investaccountvaluation.And(
			investaccountvaluation.ID(id), investaccountvaluation.HasOwnerWith(investaccount.IDIn(accXids...)))).
			Exec(context.Background())
		handleErr(err, w)

		if cnt == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	if r.Method == http.MethodPut {

		par := ent.InvestAccountValuation{}
		err = json.NewDecoder(r.Body).Decode(&par)
		handleErr(err, w)

		err := db.DB.InvestAccountValuation.UpdateOneID(id).SetRecDate(par.RecDate).SetValue(par.Value).Exec(context.Background())
		handleErr(err, w)

	}

	if r.Method == http.MethodPost {

		owner, err := xid.FromString(r.URL.Query().Get("owner"))
		handleErr(err, w)

		par := ent.InvestAccountValuation{}
		err = json.NewDecoder(r.Body).Decode(&par)
		handleErr(err, w)

		newObj, err := db.DB.InvestAccountValuation.Create().
			SetID(xid.New()).
			SetRecDate(par.RecDate).
			SetValue(par.Value).
			SetOwner(db.DB.InvestAccount.GetX(context.Background(), owner)).
			Save(context.Background())
		handleErr(err, w)
		w.Write([]byte(newObj.ID.String()))

	}

}

func handleAccsPut(r *http.Request, w http.ResponseWriter, session pages.SessionStruct) {

	id, err := xid.FromString(r.URL.Query().Get("id"))
	handleErr(err, w)

	updater := db.DB.InvestAccount.UpdateOneID(id)
	updCnt := 0

	newDescr := r.URL.Query().Get("newdescr")
	if newDescr != "" {
		updater = updater.SetDescr(newDescr)
		updCnt++
	}

	if updCnt > 0 {
		err = updater.Exec(context.Background())
		handleErr(err, w)
	}

}

func handleAccsPost(r *http.Request, w http.ResponseWriter, curSes pages.SessionStruct) {

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.

	var p []*ent.InvestAccount

	err := json.NewDecoder(r.Body).Decode(&p)
	handleErr(err, w)

	for _, pone := range p {
		existings, err := db.DB.InvestAccount.Query().
			Where(investaccount.And(
				investaccount.DescrEQ(pone.Descr),
				investaccount.HasOwnerWith(user.IDEQ(curSes.User.ID)))).
			All(context.Background())
		handleErr(err, w)

		for _, existing := range existings {

			_, err = db.DB.InvestAccountValuation.Delete().
				Where(investaccountvaluation.HasOwnerWith(investaccount.IDEQ(existing.ID))).
				Exec(context.Background())
			handleErr(err, w)

			_, err = db.DB.InvestAccountCashflow.Delete().
				Where(investaccountcashflow.HasOwnerWith(investaccount.IDEQ(existing.ID))).
				Exec(context.Background())
			handleErr(err, w)

			err = db.DB.InvestAccount.DeleteOneID(existing.ID).Exec(context.Background())
			handleErr(err, w)
		}

	}

	for _, v := range p {

		ins, err := db.DB.InvestAccount.Create().
			SetDescr(v.Descr).
			SetOwnerID(curSes.User.ID).
			Save(context.Background())
		handleErr(err, w)

		for _, vv := range v.Edges.Valuations {
			_, err = db.DB.InvestAccountValuation.Create().
				SetOwnerID(ins.ID).
				SetRecDate(vv.RecDate).
				SetValue(vv.Value).
				Save(context.Background())
			handleErr(err, w)

		}

		for _, vv := range v.Edges.Cashflows {
			_, err = db.DB.InvestAccountCashflow.Create().
				SetOwnerID(ins.ID).
				SetRecDate(vv.RecDate).
				SetQty(vv.Qty).
				Save(context.Background())
			handleErr(err, w)

		}

	}

}

func handleAccsDelete(r *http.Request, w http.ResponseWriter, session pages.SessionStruct) {

	id, err := xid.FromString(r.URL.Query().Get("id"))
	handleErr(err, w)

	tx, err := db.DB.Tx(context.Background())
	defer tx.Rollback()
	handleErr(err, w)

	cnt, err := tx.InvestAccount.Delete().
		Where(investaccount.And(investaccount.ID(id), investaccount.HasOwnerWith(user.IDEQ(session.User.ID)))).
		Exec(context.Background())
	handleErr(err, w)
	if cnt == 1 {
		_, err = tx.InvestAccountValuation.Delete().
			Where(investaccountvaluation.HasOwnerWith(investaccount.IDEQ(id))).
			Exec(context.Background())
		handleErr(err, w)

		_, err = tx.InvestAccountCashflow.Delete().
			Where(investaccountcashflow.HasOwnerWith(investaccount.IDEQ(id))).
			Exec(context.Background())
		handleErr(err, w)

		tx.Commit()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func handleAccsGetList(curSes pages.SessionStruct, w http.ResponseWriter) {
	data, err := db.DB.InvestAccount.Query().
		WithValuations(
			func(q *ent.InvestAccountValuationQuery) {
				q.Order(ent.Desc(investaccountvaluation.FieldRecDate))
			}).
		Where(investaccount.HasOwnerWith(user.IDEQ(curSes.User.ID))).
		All(context.Background())
	handleErr(err, w)

	type AccountListResult struct {
		ID      string    `json:"id"`
		Descr   string    `json:"descr"`
		RecDate time.Time `json:"rec_date"`
		Value   float64   `json:"value"`
	}

	data2 := make([]AccountListResult, len(data))
	for i, v := range data {

		data2[i] = AccountListResult{ID: v.ID.String(), Descr: v.Descr}
		if len(v.Edges.Valuations) > 0 {
			data2[i].RecDate = v.Edges.Valuations[0].RecDate
			data2[i].Value = v.Edges.Valuations[0].Value
		}
	}

	sort.Slice(data2,
		func(i, j int) bool { return data2[i].Value > data2[j].Value })

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data2)
	handleErr(err, w)

}

func handleAccsWeekflow(w http.ResponseWriter, r *http.Request, session pages.SessionStruct) {

	startDate := session.User.StartInvestAccountsFlow

	var err error

	accsMap, err := session.GetInvestAccountXidsMap()
	handleErr(err, w)

	ids := strings.Split(r.URL.Query().Get("ids"), ",")
	if len(ids) == 0 {
		log.Println("No ids provided")
		http.Error(w, "No ids provided", http.StatusBadRequest)
		return
	}

	xids := make([]xid.ID, 0)
	for _, v := range ids {
		newXid, err := xid.FromString(v)
		handleErr(err, w)
		if _, hasValue := accsMap[newXid]; hasValue {
			xids = append(xids, newXid)
		}
	}

	evals, err := db.DB.InvestAccountValuation.Query().WithOwner().
		Where(investaccountvaluation.And(
			investaccountvaluation.HasOwnerWith(investaccount.IDIn(xids...)),
			investaccountvaluation.RecDateGTE(startDate.Add(-time.Hour*24*7)))).
		Order(ent.Asc(investaccountvaluation.FieldRecDate)).
		All(context.Background())
	handleErr(err, w)
	if len(evals) == 0 {
		log.Println("No evaluations found")
		return
	}

	cf, err := db.DB.InvestAccountCashflow.Query().WithOwner().
		Where(investaccountcashflow.And(
			investaccountcashflow.HasOwnerWith(investaccount.IDIn(xids...)),
			investaccountcashflow.RecDateGTE(startDate.Add(-time.Hour*24*7)))).
		Order(ent.Asc(investaccountcashflow.FieldRecDate)).
		All(context.Background())
	handleErr(err, w)

	// raw items tree by acc, eow
	type weekRec struct {
		eow time.Time
		ev  float64
		cf  float64
	}
	allraws := make(map[xid.ID]map[time.Time]*weekRec)
	alleows := make(map[time.Time]bool)
	noevMarker := -9999999.0
	for _, accid := range xids {
		allraws[accid] = make(map[time.Time]*weekRec)
		raws := allraws[accid]
		for _, v := range evals {
			if v.Edges.Owner.ID != accid {
				continue
			}
			eow := endOfWeek(v.RecDate)
			alleows[eow] = true
			if rec, ok := raws[eow]; !ok {
				raws[eow] = &weekRec{eow: eow, ev: v.Value}
			} else {
				rec.ev = v.Value
			}
		}
		for _, v := range cf {
			if v.Edges.Owner.ID != accid {
				continue
			}
			eow := endOfWeek(v.RecDate)
			alleows[eow] = true
			if rec, ok := raws[eow]; !ok {
				raws[eow] = &weekRec{eow: eow, ev: noevMarker, cf: v.Qty}
			} else {
				rec.cf += v.Qty
			}

		}
	}

	// add missing weeks to allraws
	for eow := range alleows {
		for _, accid := range xids {
			if _, ok := allraws[accid][eow]; !ok {
				allraws[accid][eow] = &weekRec{eow: eow, ev: noevMarker}
			}
		}
	}

	//flat and sort records for each accs SEPARATED, fill noevs
	alllines := make(map[xid.ID][]*weekRec)
	for _, accid := range xids {
		lines := make([]*weekRec, 0)
		if raws, ok := allraws[accid]; ok {
			for _, v := range raws {
				lines = append(lines, v)
			}
		}
		sort.Slice(lines, func(i, j int) bool { return lines[i].eow.Before(lines[j].eow) })
		lastEv := 0.00
		for _, v := range lines {
			if v.ev == noevMarker {
				v.ev = lastEv
			}
			lastEv = v.ev
		}
		alllines[accid] = lines
	}

	//merge all accs into one
	merged := make([]*weekRec, 0)
	for _, accslice := range alllines {
		for _, rec := range accslice {
			found := false
			for _, mc := range merged {
				if mc.eow.Equal(rec.eow) {
					found = true
					mc.ev += rec.ev
					mc.cf += rec.cf
					break
				}
			}
			if !found {
				merged = append(merged, rec)
			}
		}
	}
	sort.Slice(merged, func(i, j int) bool { return merged[i].eow.Before(merged[j].eow) })

	// trim merged for startDate
	for len(merged) > 1 && merged[0].eow.Before(startDate) {
		merged = merged[1:]
	}
	merged[0].cf = merged[0].ev

	type WeekLine struct {
		WNum          int       `json:"wNum"`
		Eow           time.Time `json:"eow"`
		Eval          float64   `json:"eval"`
		WeekCashflow  float64   `json:"weekCashflow"`
		TotalCashflow float64   `json:"totalCashflow"`
		TotalProfit   float64   `json:"totalProfit"`
		WeekProfit    float64   `json:"weekProfit"`
		TotalYield    float64   `json:"totalYield"`
		Yield         float64   `json:"yield"`
		YearYield     float64   `json:"yearYield"`
	}

	// enrich merged with totals
	res := make([]*WeekLine, len(merged))
	wnum := 0
	totalcf := 0.0
	prevProfit := 0.0
	for i, v := range merged {
		wnum++
		res[i] = &WeekLine{
			WNum:         wnum,
			Eow:          v.eow,
			WeekCashflow: v.cf,
			Eval:         v.ev,
		}
		totalcf += res[i].WeekCashflow
		res[i].TotalCashflow = totalcf
		res[i].TotalProfit = res[i].Eval - res[i].TotalCashflow
		res[i].WeekProfit = res[i].TotalProfit - prevProfit
		prevProfit += res[i].WeekProfit
		if res[i].TotalCashflow == 0 {
			res[i].TotalYield = 0
		} else {
			res[i].TotalYield = res[i].TotalProfit / res[i].TotalCashflow * 100
		}

		if res[i].TotalCashflow == 0 || (res[i].Eow.Sub(merged[0].eow).Hours() < 24) {
			res[i].YearYield = 0
		} else {
			days := res[i].Eow.Sub(merged[0].eow).Hours() / 24.0
			res[i].YearYield = res[i].TotalProfit / res[i].TotalCashflow / days * 360.0 * 100.0
		}

	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	handleErr(err, w)

}

func handleAccsGetOne(w http.ResponseWriter, r *http.Request, session pages.SessionStruct, parID string) {

	xid, err := xid.FromString(parID)
	handleErr(err, w)

	accsMap, err := session.GetInvestAccountXidsMap()
	handleErr(err, w)

	if _, has := accsMap[xid]; !has {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	data, err := db.DB.InvestAccount.Query().
		WithCashflows(
			func(q *ent.InvestAccountCashflowQuery) {
				q.Order(ent.Asc(investaccountcashflow.FieldRecDate))
			}).
		WithValuations(
			func(q *ent.InvestAccountValuationQuery) {
				q.Order(ent.Asc(investaccountvaluation.FieldRecDate))
			}).
		Where(investaccount.IDEQ(xid)).All(context.Background())
	handleErr(err, w)
	if len(data) != 1 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data[0])
	handleErr(err, w)

}

func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func endOfWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, 6-int(t.Weekday()))
}

package api

import (
	"context"
	"encoding/json"
	"log"
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
		pages.HandleErr(err, w)
	}

	id := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {

		if len(id) == 0 {

			res, err := db.DB.Emitent.Query().All(context.Background())
			pages.HandleErr(err, w)

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			pages.HandleErr(err, w)
			return

		} else {

			xid, err := xid.FromString(id)
			pages.HandleErr(err, w)
			res, err := db.DB.Emitent.Query().Where(emitent.IDEQ(xid)).All(context.Background())
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

		buf := ent.Emitent{ID: xid.New()}
		err := json.NewDecoder(r.Body).Decode(&buf)
		pages.HandleErr(err, w)

		//deleteOne(buf.ID, w)

		tx, err := db.DB.Tx(context.Background())
		pages.HandleErr(err, w)
		defer func() {
			err := tx.Rollback()
			if err != nil {
				log.Printf("Error when rollback in post method on emitents page: %s\n", err.Error())
			}
		}()

		newdata, err := tx.Emitent.Create().
			SetID(buf.ID).
			SetDescr(buf.Descr).
			SetIndustryID(buf.Edges.Industry.ID).
			Save(context.Background())
		pages.HandleErr(err, w)

		for _, v := range buf.Edges.Reports {
			_, err := tx.Report.Create().
				SetEmitentID(newdata.ID).
				SetReportYear(v.ReportYear).
				SetReportQuarter(v.ReportQuarter).
				SetReportDate(v.ReportDate).
				SetPnlRevenueYtd(v.PnlRevenueYtd).
				SetPnlAmortizationYtd(v.PnlAmortizationYtd).
				SetPnlOperatingIncomeYtd(v.PnlOperatingIncomeYtd).
				SetPnlInterestIncomeYtd(v.PnlInterestIncomeYtd).
				SetPnlInterestExpensesYtd(v.PnlInterestExpensesYtd).
				SetPnlIncomeTaxYtd(v.PnlIncomeTaxYtd).
				SetPnlNetIncomeYtd(v.PnlNetIncomeYtd).
				SetCfCashSld(v.CfCashSld).
				SetCfNonCurrentLiabilitiesSld(v.CfNonCurrentLiabilitiesSld).
				SetCfCurrentLiabilitesSld(v.CfCurrentLiabilitesSld).
				SetCfNonControllingSld(v.CfNonControllingSld).
				SetCfEquitySld(v.CfEquitySld).
				SetCfTotalSld(v.CfTotalSld).
				Save(context.Background())
			pages.HandleErr(err, w)

		}

		err = tx.Commit()
		pages.HandleErr(err, w)

		_, err = w.Write([]byte(newdata.ID.String()))
		pages.HandleErr(err, w)
		return

	}

	if r.Method == http.MethodDelete {

		if !session.UserIsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		xid, err := xid.FromString(id)
		pages.HandleErr(err, w)
		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		deleteOne(xid, w)
		return

	}

}

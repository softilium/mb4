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

		tx, err := db.DB.Tx(context.Background())
		handleErr(err, w)
		defer tx.Rollback()

		newdata, err := tx.Emitent.Create().
			SetID(buf.ID).
			SetDescr(buf.Descr).
			SetIndustryID(buf.Edges.Industry.ID).
			Save(context.Background())
		handleErr(err, w)

		for _, v := range buf.Edges.Reports {
			v.Edges.Emitent.ID = newdata.ID
			_, err := tx.Report.Create().
				SetEmitentID(newdata.ID).
				SetYear(v.Year).
				SetQuarter(v.Quarter).
				SetReportDate(v.ReportDate).
				SetPnlRevenueYtd(v.PnlRevenueYtd).
				SetPnlAmortizationYtd(v.PnlAmortizationYtd).
				SetPnlOperationIncomeYtd(v.PnlOperationIncomeYtd).
				SetPnlInterestIncomeYtd(v.PnlInterestIncomeYtd).
				SetPnlInterestExpensesYtd(v.PnlInterestExpensesYtd).
				SetPnlIncomeTaxYtd(v.PnlIncomeTaxYtd).
				SetPnlNetIncomeYtd(v.PnlNetIncomeYtd).
				SetCfCashSld(v.CfCashSld).
				SetCfNonCurrentLiabilitiesSld(v.CfNonCurrentLiabilitiesSld).
				SetCfCurrentLiabilitesSld(v.CfCurrentLiabilitesSld).
				SetCfNonControlledSld(v.CfNonControlledSld).
				SetCfEquitySld(v.CfEquitySld).
				SetCfTotalSld(v.CfTotalSld).
				Save(context.Background())
			handleErr(err, w)

		}

		tx.Commit()

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

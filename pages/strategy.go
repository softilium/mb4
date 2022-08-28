package pages

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/rs/xid"
	"github.com/softilium/mb4/backtest"
	"github.com/softilium/mb4/cube"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/strategy"
	"github.com/softilium/mb4/ent/strategyfactor"
	"github.com/softilium/mb4/ent/strategyfilter"
	"github.com/softilium/mb4/ent/strategyfixedticker"
	"github.com/softilium/mb4/ent/ticker"
)

func Strategy(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		{

			mode := r.URL.Query().Get("mode")

			switch mode {
			case "obj":
				{
					id := r.URL.Query().Get("id")
					xid, err := xid.FromString(id)
					HandleErr(err, w)

					obj := getStrategyObj(xid, w)

					w.Header().Set("Content-Type", "application/json")
					err = json.NewEncoder(w).Encode(obj)
					HandleErr(err, w)

				}
			case "results":
				{

					id := r.URL.Query().Get("id")
					xid, err := xid.FromString(id)
					HandleErr(err, w)

					strategy := getStrategyObj(xid, w)

					res := backtest.Simulate(strategy, cube.Market, (*time.Time)(strategy.StartSimulation), strategy.StartAmount, strategy.WeekRefillAmount, nil)

					w.Header().Set("Content-Type", "application/json")
					err = json.NewEncoder(w).Encode(res)
					HandleErr(err, w)
				}
			case "renderinfo":
				{
					type RenderInfoType struct {
						FilterOps           []domains.DomainItem[domains.FilterOp]
						SameEmitentPolicies []domains.DomainItem[domains.SameEmitentPolicy]
						Tickers             []*domains.DomainItem[string]
						Industries          []*domains.DomainItem[string]
						ReportValues        []domains.DomainItem[domains.ReportValue]
						ReportValueTypes    []domains.DomainItem[domains.ReportValueType]
						FilterValueKinds    []domains.DomainItem[domains.FilterValueKind]
						FilterOpsShort      []domains.DomainItem[domains.FilterOp]
					}
					ri := RenderInfoType{
						FilterOps:           domains.FilterOps.Slice(),
						SameEmitentPolicies: domains.SameEmitentPolicies.Slice(),
						ReportValues:        domains.ReportValues.Slice(),
						ReportValueTypes:    domains.ReportValueTypes.Slice(),
						FilterValueKinds:    domains.FilterValueKinds.Slice(),
					}

					ri.FilterOpsShort = []domains.DomainItem[domains.FilterOp]{
						{Id: domains.FilterOp_Eq, Descr: domains.FilterOps.ById(domains.FilterOp_Eq).Descr},
						{Id: domains.FilterOp_Ne, Descr: domains.FilterOps.ById(domains.FilterOp_Ne).Descr},
					}

					data, err := db.DB.Ticker.Query().Order(ent.Asc(ticker.FieldDescr)).All(context.Background())
					HandleErr(err, w)
					ri.Tickers = make([]*domains.DomainItem[string], len(data))
					for idx, v := range data {
						ri.Tickers[idx] = &domains.DomainItem[string]{Id: v.ID, Descr: v.Descr}
					}

					di, err := db.DB.Industry.Query().Order(ent.Asc(ticker.FieldDescr)).All(context.Background())
					HandleErr(err, w)
					ri.Industries = make([]*domains.DomainItem[string], len(di))
					for idx, v := range di {
						ri.Industries[idx] = &domains.DomainItem[string]{Id: v.ID, Descr: v.Descr}
					}

					w.Header().Set("Content-Type", "application/json")
					err = json.NewEncoder(w).Encode(ri)
					HandleErr(err, w)

				}
			default: //page
				{
					si := LoadSessionStruct(r)
					pageData := struct {
						SessionStruct
						StrategyId string
					}{SessionStruct: si, StrategyId: r.URL.Query().Get("id")}
					pageData.Vue = true
					pageData.Echarts = true

					tmpl, err := pongo2.FromCache("pages/strategy.html")
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					err = tmpl.ExecuteWriter(pongo2.Context{"pd": pageData}, w)
					HandleErr(err, w)
				}
			}
		}
	case http.MethodPut:
		{
			obj := ent.Strategy{}
			err := json.NewDecoder(r.Body).Decode(&obj)
			HandleErr(err, w)

			tx, err := db.DB.BeginTx(context.Background(), nil)
			HandleErr(err, w)
			defer func() {
				_ = tx.Rollback() // we call rollback in any case. In most cases rollback will throw error "transaction has already been commited"
			}()

			err = tx.Strategy.UpdateOne(&obj).
				SetDescr(obj.Descr).
				SetMaxTickers(obj.MaxTickers).
				SetMaxTickersPerIndustry(obj.MaxTickersPerIndustry).
				SetBaseIndex(obj.BaseIndex).
				SetWeekRefillAmount(obj.WeekRefillAmount).
				SetStartAmount(obj.StartAmount).
				SetStartSimulation(obj.StartSimulation).
				SetBuyOnlyLowPrice(obj.BuyOnlyLowPrice).
				SetAllowLossWhenSell(obj.AllowLossWhenSell).
				SetAllowSellToFit(obj.AllowSellToFit).
				SetSameEmitent(obj.SameEmitent).
				Exec(context.Background())
			HandleErr(err, w)

			_, err = tx.StrategyFixedTicker.Delete().Where(strategyfixedticker.HasStrategyWith(strategy.ID(obj.ID))).Exec(context.Background())
			HandleErr(err, w)

			for idx, v := range obj.Edges.FixedTickers {
				_, err = tx.StrategyFixedTicker.Create().
					SetStrategyID(obj.ID).
					SetLineNum(idx + 1).
					SetIsUsed(v.IsUsed).
					SetTicker(v.Ticker).
					SetShare(v.Share).
					Save(context.Background())
				HandleErr(err, w)
			}

			_, err = tx.StrategyFilter.Delete().Where(strategyfilter.HasStrategyWith(strategy.ID(obj.ID))).Exec(context.Background())
			HandleErr(err, w)
			for idx, v := range obj.Edges.Filters {
				_, err = tx.StrategyFilter.Create().
					SetStrategyID(obj.ID).
					SetLineNum(idx + 1).
					SetIsUsed(v.IsUsed).
					SetLeftValueKind(v.LeftValueKind).
					SetLeftReportValue(v.LeftReportValue).
					SetLeftReportValueType(v.LeftReportValueType).
					SetOperation(v.Operation).
					SetRightValueStr(v.RightValueStr).
					SetRightValueFloat(v.RightValueFloat).
					Save(context.Background())
				HandleErr(err, w)
			}

			_, err = tx.StrategyFactor.Delete().Where(strategyfactor.HasStrategyWith(strategy.ID(obj.ID))).Exec(context.Background())
			HandleErr(err, w)
			for idx, v := range obj.Edges.Factors {
				_, err = tx.StrategyFactor.Create().
					SetStrategyID(obj.ID).
					SetLineNum(idx + 1).
					SetIsUsed(v.IsUsed).
					SetRK(v.RK).
					SetRVT(v.RVT).
					SetMinAcceptable(v.MinAcceptable).
					SetMaxAcceptable(v.MaxAcceptable).
					SetInverse(v.Inverse).
					SetK(v.K).
					SetGist(v.Gist).
					Save(context.Background())
				HandleErr(err, w)
			}
			err = tx.Commit()
			HandleErr(err, w)

		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func getStrategyObj(xid xid.ID, w http.ResponseWriter) *ent.Strategy {

	obj, err := db.DB.Strategy.
		Query().
		WithFactors(
			func(q *ent.StrategyFactorQuery) { q.Order(ent.Asc(strategyfactor.FieldLineNum)) }).
		WithFilters(
			func(q *ent.StrategyFilterQuery) { q.Order(ent.Asc(strategyfactor.FieldLineNum)) }).
		WithFixedTickers(
			func(q *ent.StrategyFixedTickerQuery) { q.Order(ent.Asc(strategyfixedticker.FieldLineNum)) }).
		Where(strategy.ID(xid)).
		Only(context.Background())
	HandleErr(err, w)
	return obj

}

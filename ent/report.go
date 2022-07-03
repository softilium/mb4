// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/ent/report"
)

// Report is the model entity for the Report schema.
type Report struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// ReportYear holds the value of the "ReportYear" field.
	ReportYear int `json:"ReportYear,omitempty"`
	// ReportQuarter holds the value of the "ReportQuarter" field.
	ReportQuarter int `json:"ReportQuarter,omitempty"`
	// ReportDate holds the value of the "ReportDate" field.
	ReportDate time.Time `json:"ReportDate,omitempty"`
	// PnlRevenueYtd holds the value of the "PnlRevenueYtd" field.
	PnlRevenueYtd float64 `json:"PnlRevenueYtd,omitempty"`
	// PnlAmortizationYtd holds the value of the "PnlAmortizationYtd" field.
	PnlAmortizationYtd float64 `json:"PnlAmortizationYtd,omitempty"`
	// PnlOperatingIncomeYtd holds the value of the "PnlOperatingIncomeYtd" field.
	PnlOperatingIncomeYtd float64 `json:"PnlOperatingIncomeYtd,omitempty"`
	// PnlInterestIncomeYtd holds the value of the "PnlInterestIncomeYtd" field.
	PnlInterestIncomeYtd float64 `json:"PnlInterestIncomeYtd,omitempty"`
	// PnlInterestExpensesYtd holds the value of the "PnlInterestExpensesYtd" field.
	PnlInterestExpensesYtd float64 `json:"PnlInterestExpensesYtd,omitempty"`
	// PnlIncomeTaxYtd holds the value of the "PnlIncomeTaxYtd" field.
	PnlIncomeTaxYtd float64 `json:"PnlIncomeTaxYtd,omitempty"`
	// PnlNetIncomeYtd holds the value of the "PnlNetIncomeYtd" field.
	PnlNetIncomeYtd float64 `json:"PnlNetIncomeYtd,omitempty"`
	// CfCashSld holds the value of the "CfCashSld" field.
	CfCashSld float64 `json:"CfCashSld,omitempty"`
	// CfNonCurrentLiabilitiesSld holds the value of the "CfNonCurrentLiabilitiesSld" field.
	CfNonCurrentLiabilitiesSld float64 `json:"CfNonCurrentLiabilitiesSld,omitempty"`
	// CfCurrentLiabilitesSld holds the value of the "CfCurrentLiabilitesSld" field.
	CfCurrentLiabilitesSld float64 `json:"CfCurrentLiabilitesSld,omitempty"`
	// CfNonControllingSld holds the value of the "CfNonControllingSld" field.
	CfNonControllingSld float64 `json:"CfNonControllingSld,omitempty"`
	// CfEquitySld holds the value of the "CfEquitySld" field.
	CfEquitySld float64 `json:"CfEquitySld,omitempty"`
	// CfTotalSld holds the value of the "CfTotalSld" field.
	CfTotalSld float64 `json:"CfTotalSld,omitempty"`
	// URL holds the value of the "Url" field.
	URL string `json:"Url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReportQuery when eager-loading is set.
	Edges           ReportEdges `json:"edges"`
	emitent_reports *xid.ID
}

// ReportEdges holds the relations/edges for other nodes in the graph.
type ReportEdges struct {
	// Emitent holds the value of the Emitent edge.
	Emitent *Emitent `json:"Emitent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmitentOrErr returns the Emitent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReportEdges) EmitentOrErr() (*Emitent, error) {
	if e.loadedTypes[0] {
		if e.Emitent == nil {
			// The edge Emitent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: emitent.Label}
		}
		return e.Emitent, nil
	}
	return nil, &NotLoadedError{edge: "Emitent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Report) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case report.FieldPnlRevenueYtd, report.FieldPnlAmortizationYtd, report.FieldPnlOperatingIncomeYtd, report.FieldPnlInterestIncomeYtd, report.FieldPnlInterestExpensesYtd, report.FieldPnlIncomeTaxYtd, report.FieldPnlNetIncomeYtd, report.FieldCfCashSld, report.FieldCfNonCurrentLiabilitiesSld, report.FieldCfCurrentLiabilitesSld, report.FieldCfNonControllingSld, report.FieldCfEquitySld, report.FieldCfTotalSld:
			values[i] = new(sql.NullFloat64)
		case report.FieldReportYear, report.FieldReportQuarter:
			values[i] = new(sql.NullInt64)
		case report.FieldURL:
			values[i] = new(sql.NullString)
		case report.FieldReportDate:
			values[i] = new(sql.NullTime)
		case report.FieldID:
			values[i] = new(xid.ID)
		case report.ForeignKeys[0]: // emitent_reports
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Report", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Report fields.
func (r *Report) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case report.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case report.FieldReportYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ReportYear", values[i])
			} else if value.Valid {
				r.ReportYear = int(value.Int64)
			}
		case report.FieldReportQuarter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ReportQuarter", values[i])
			} else if value.Valid {
				r.ReportQuarter = int(value.Int64)
			}
		case report.FieldReportDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ReportDate", values[i])
			} else if value.Valid {
				r.ReportDate = value.Time
			}
		case report.FieldPnlRevenueYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlRevenueYtd", values[i])
			} else if value.Valid {
				r.PnlRevenueYtd = value.Float64
			}
		case report.FieldPnlAmortizationYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlAmortizationYtd", values[i])
			} else if value.Valid {
				r.PnlAmortizationYtd = value.Float64
			}
		case report.FieldPnlOperatingIncomeYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlOperatingIncomeYtd", values[i])
			} else if value.Valid {
				r.PnlOperatingIncomeYtd = value.Float64
			}
		case report.FieldPnlInterestIncomeYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlInterestIncomeYtd", values[i])
			} else if value.Valid {
				r.PnlInterestIncomeYtd = value.Float64
			}
		case report.FieldPnlInterestExpensesYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlInterestExpensesYtd", values[i])
			} else if value.Valid {
				r.PnlInterestExpensesYtd = value.Float64
			}
		case report.FieldPnlIncomeTaxYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlIncomeTaxYtd", values[i])
			} else if value.Valid {
				r.PnlIncomeTaxYtd = value.Float64
			}
		case report.FieldPnlNetIncomeYtd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field PnlNetIncomeYtd", values[i])
			} else if value.Valid {
				r.PnlNetIncomeYtd = value.Float64
			}
		case report.FieldCfCashSld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfCashSld", values[i])
			} else if value.Valid {
				r.CfCashSld = value.Float64
			}
		case report.FieldCfNonCurrentLiabilitiesSld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfNonCurrentLiabilitiesSld", values[i])
			} else if value.Valid {
				r.CfNonCurrentLiabilitiesSld = value.Float64
			}
		case report.FieldCfCurrentLiabilitesSld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfCurrentLiabilitesSld", values[i])
			} else if value.Valid {
				r.CfCurrentLiabilitesSld = value.Float64
			}
		case report.FieldCfNonControllingSld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfNonControllingSld", values[i])
			} else if value.Valid {
				r.CfNonControllingSld = value.Float64
			}
		case report.FieldCfEquitySld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfEquitySld", values[i])
			} else if value.Valid {
				r.CfEquitySld = value.Float64
			}
		case report.FieldCfTotalSld:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field CfTotalSld", values[i])
			} else if value.Valid {
				r.CfTotalSld = value.Float64
			}
		case report.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Url", values[i])
			} else if value.Valid {
				r.URL = value.String
			}
		case report.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field emitent_reports", values[i])
			} else if value.Valid {
				r.emitent_reports = new(xid.ID)
				*r.emitent_reports = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryEmitent queries the "Emitent" edge of the Report entity.
func (r *Report) QueryEmitent() *EmitentQuery {
	return (&ReportClient{config: r.config}).QueryEmitent(r)
}

// Update returns a builder for updating this Report.
// Note that you need to call Report.Unwrap() before calling this method if this Report
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Report) Update() *ReportUpdateOne {
	return (&ReportClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Report entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Report) Unwrap() *Report {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Report is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Report) String() string {
	var builder strings.Builder
	builder.WriteString("Report(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", ReportYear=")
	builder.WriteString(fmt.Sprintf("%v", r.ReportYear))
	builder.WriteString(", ReportQuarter=")
	builder.WriteString(fmt.Sprintf("%v", r.ReportQuarter))
	builder.WriteString(", ReportDate=")
	builder.WriteString(r.ReportDate.Format(time.ANSIC))
	builder.WriteString(", PnlRevenueYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlRevenueYtd))
	builder.WriteString(", PnlAmortizationYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlAmortizationYtd))
	builder.WriteString(", PnlOperatingIncomeYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlOperatingIncomeYtd))
	builder.WriteString(", PnlInterestIncomeYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlInterestIncomeYtd))
	builder.WriteString(", PnlInterestExpensesYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlInterestExpensesYtd))
	builder.WriteString(", PnlIncomeTaxYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlIncomeTaxYtd))
	builder.WriteString(", PnlNetIncomeYtd=")
	builder.WriteString(fmt.Sprintf("%v", r.PnlNetIncomeYtd))
	builder.WriteString(", CfCashSld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfCashSld))
	builder.WriteString(", CfNonCurrentLiabilitiesSld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfNonCurrentLiabilitiesSld))
	builder.WriteString(", CfCurrentLiabilitesSld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfCurrentLiabilitesSld))
	builder.WriteString(", CfNonControllingSld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfNonControllingSld))
	builder.WriteString(", CfEquitySld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfEquitySld))
	builder.WriteString(", CfTotalSld=")
	builder.WriteString(fmt.Sprintf("%v", r.CfTotalSld))
	builder.WriteString(", Url=")
	builder.WriteString(r.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Reports is a parsable slice of Report.
type Reports []*Report

func (r Reports) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

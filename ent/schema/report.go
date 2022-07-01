package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/rs/xid"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("Year").Range(1999, 2999),
		field.Int("Quarter").Range(1, 4),
		field.Time("ReportDate"),
		field.Float("PnlRevenueYtd"),
		field.Float("PnlAmortizationYtd"),
		field.Float("PnlOperationIncomeYtd"),
		field.Float("PnlInterestIncomeYtd"),
		field.Float("PnlInterestExpensesYtd"),
		field.Float("PnlIncomeTaxYtd"),
		field.Float("PnlNetIncomeYtd"),
		field.Float("CfCashSld"),
		field.Float("CfNonCurrentLiabilitiesSld"),
		field.Float("CfCurrentLiabilitesSld"),
		field.Float("CfNonControlledSld"),
		field.Float("CfEquitySld"),
		field.Float("CfTotalSld"),
		field.String("Url").Optional().MaxLen(255),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Emitent", Emitent.Type).Ref("Reports").Required().Unique(),
	}
}

func (Report) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("Emitent").Fields("Year").Fields("Quarter").Unique(),
	}
}

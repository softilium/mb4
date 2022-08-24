package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// InvestAccountCashflow holds the schema definition for the InvestAccountCashflow entity.
type InvestAccountCashflow struct {
	ent.Schema
}

// Fields of the InvestAccountCashflow.
func (InvestAccountCashflow) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Time("RecDate").SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Float("Qty"),
	}
}

// Edges of the InvestAccountCashflow.
func (InvestAccountCashflow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Owner", InvestAccount.Type).Ref("Cashflows").Required().Unique(),
	}
}

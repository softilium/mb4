package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// InvestAccount holds the schema definition for the InvestAccount entity.
type InvestAccount struct {
	ent.Schema
}

// Fields of the InvestAccount.
func (InvestAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New),
		field.String("Descr").NotEmpty().MinLen(3).MaxLen(100),
	}
}

// Edges of the InvestAccount.
func (InvestAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Owner", User.Type).Ref("InvestAccounts").Required().Unique(),
		edge.To("Cashflows", InvestAccountCashflow.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("Valuations", InvestAccountValuation.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

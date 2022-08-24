package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// InvestAccountValuation holds the schema definition for the InvestAccountValuation entity.
type InvestAccountValuation struct {
	ent.Schema
}

// Fields of the InvestAccountValuation.
func (InvestAccountValuation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Time("RecDate").SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Float("Value"),
	}
}

// Edges of the InvestAccountValuation.
func (InvestAccountValuation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Owner", InvestAccount.Type).Ref("Valuations").Required().Unique(),
	}
}

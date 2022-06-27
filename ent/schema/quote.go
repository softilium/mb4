package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Quote holds the schema definition for the Quote entity.
type Quote struct {
	ent.Schema
}

// Fields of the Quote.
func (Quote) Fields() []ent.Field {
	return []ent.Field{
		field.Time("D").Immutable(),
		field.Float("O").Immutable(),
		field.Float("C").Immutable(),
		field.Float("H").Immutable(),
		field.Float("L").Immutable(),
		field.Float("V").Immutable(),
		field.Float("Cap").Immutable(),
		field.Float("DivSum_5Y").Immutable(),
		field.Float("DivYield_5Y").Immutable(),
		field.Int("LotSize").Positive(),
		field.Int("ListLevel").Positive(),
	}
}

// Edges of the Quote.
func (Quote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Ticker", Ticker.Type).Ref("Quotes").Unique(),
	}
}

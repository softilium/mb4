package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Emitent holds the schema definition for the Emitent entity.
type Emitent struct {
	ent.Schema
}

// Fields of the Emitent.
func (Emitent) Fields() []ent.Field {
	return []ent.Field{
		field.String("Descr").NotEmpty().MinLen(3).MaxLen(100).Unique(),
	}
}

// Edges of the Emitent.
func (Emitent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Industry", Industry.Type).Ref("Emitents").Unique(),
		edge.To("Tickers", Ticker.Type),
	}
}

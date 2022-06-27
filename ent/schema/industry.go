package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Industry holds the schema definition for the Industry entity.
type Industry struct {
	ent.Schema
}

// Fields of the Industry.
func (Industry) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().MinLen(3).MaxLen(20).Unique().Immutable(), //use natural keys instead of XID here
		field.String("Descr").NotEmpty().MinLen(3).MaxLen(100).Unique(),
	}
}

// Edges of the Industry.
func (Industry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Emitents", Emitent.Type),
	}
}

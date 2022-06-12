package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Emitent holds the schema definition for the Emitent entity.
type Emitent struct {
	ent.Schema
}

// Fields of the Emitent.
func (Emitent) Fields() []ent.Field {
	return []ent.Field{
		field.String("Descr").NotEmpty().MinLen(1).MaxLen(100),
	}
}

// Edges of the Emitent.
func (Emitent) Edges() []ent.Edge {
	return nil
}

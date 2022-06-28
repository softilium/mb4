package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// Emitent holds the schema definition for the Emitent entity.
type Emitent struct {
	ent.Schema
}

// Fields of the Emitent.
func (Emitent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New),
		field.String("Descr").NotEmpty().MinLen(1).MaxLen(100).Unique(),
	}
}

// Edges of the Emitent.
func (Emitent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Industry", Industry.Type).Ref("Emitents").Required().Unique(),
		edge.To("Tickers", Ticker.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

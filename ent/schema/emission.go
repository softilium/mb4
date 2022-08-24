package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// Emission holds the schema definition for the Emission entity.
type Emission struct {
	ent.Schema
}

// Fields of the Emission.
func (Emission) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Time("RecDate").SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Int64("Size"),
		field.Int("FreeFloat").Optional(),
		field.Int("LotSize").Optional(),
		field.Int("ListingLevel"),
	}
}

// Edges of the Emission.
func (Emission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Ticker", Ticker.Type).Ref("Emissions").Required().Unique(),
	}
}

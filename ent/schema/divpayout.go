package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	DivPayoutStatus_CompanyCharter = 10
	DivPayoutStatus_DivPolicy      = 20
	DivPayoutStatus_Planned        = 30
	DivPayoutStatus_Fact           = 40
)

// DivPayout holds the schema definition for the DivPayout entity.
type DivPayout struct {
	ent.Schema
}

// Fields of the DivPayout.
func (DivPayout) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ForYear").Range(1900, 2999),
		field.Int("ForQuarter").Range(1, 4),
		field.Time("CloseDate").SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Int("Status").Range(DivPayoutStatus_CompanyCharter, DivPayoutStatus_Fact),
		field.Float("DPS"),
	}
}

// Edges of the DivPayout.
func (DivPayout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Tickers", Ticker.Type).Ref("DivPayouts").Required().Unique(),
	}
}

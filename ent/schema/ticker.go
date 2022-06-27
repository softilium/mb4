package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	TickerKind_Stock             = 100
	TickerKind_DepositaryReceipt = 150
	TickerKind_StockPref         = 200
	TickerKind_Bond              = 300
	TickerKind_Etf               = 400
	TickerKind_Index             = 500
)

// Ticker holds the schema definition for the Ticker entity.
type Ticker struct {
	ent.Schema
}

// Fields of the Ticker.
func (Ticker) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().MinLen(1).MaxLen(20).Unique().Immutable(),
		field.String("Descr").NotEmpty().MinLen(3).MaxLen(50).Unique(),
		field.Int32("Kind").Default(TickerKind_Stock),
	}
}

// Edges of the Ticker.
func (Ticker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Emitent", Emitent.Type).Ref("Tickers").Unique(),
		edge.To("Quotes", Quote.Type),
		edge.To("DivPayouts", DivPayout.Type),
	}
}

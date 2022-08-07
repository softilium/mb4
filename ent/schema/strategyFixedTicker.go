package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type StrategyFixedTicker struct {
	ent.Schema
}

func (StrategyFixedTicker) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("LineNum").Range(1, 10000).Default(1),
		field.Bool("IsUsed").Default(true),
	}
}

func (StrategyFixedTicker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Strategy", Strategy.Type).Ref("FixedTickers").Unique().Required(),
	}
}

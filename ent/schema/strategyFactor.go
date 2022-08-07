package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type StrategyFactor struct {
	ent.Schema
}

func (StrategyFactor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("LineNum").Range(1, 10000).Default(1),
		field.Bool("IsUsed").Default(true),
		field.Int("RK"),  // cube.RK_* report value (Revenue, EBITDA, etc.)
		field.Int("RVT"), // cube.RVT_* - report value type (YTD, LTM, ...)
		field.Float("MinAcceptabe"),
		field.Float("MaxAcceptable"),
		field.Bool("Inverse").Default(false),
		field.Float("K").Default(1),
		field.Float("Gist").Default(1),
	}
}

func (StrategyFactor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Strategy", Strategy.Type).Ref("Factors").Unique().Required(),
	}
}

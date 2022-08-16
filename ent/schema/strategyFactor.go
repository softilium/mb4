package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/domains"
)

type StrategyFactor struct {
	ent.Schema
}

func (StrategyFactor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("LineNum").Range(1, 10000).Default(1),
		field.Bool("IsUsed").Default(true),
		field.Int("RK").GoType(domains.ReportValue(0)),
		field.Int("RVT").GoType(domains.ReportValueType(0)),
		field.Float("MinAcceptable"),
		field.Float("MaxAcceptable"),
		field.Bool("Inverse").Default(false),
		field.Float("K").Default(1.0),
		field.Float("Gist").Default(1),
	}
}

func (StrategyFactor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Strategy", Strategy.Type).Ref("Factors").Unique().Required(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/domains"
)

type Strategy struct {
	ent.Schema
}

func (Strategy) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.String("Descr").MaxLen(100).NotEmpty(),
		field.Int("MaxTickers").Default(10),
		field.Int("MaxTickersPerIndustry").Default(5),
		field.String("BaseIndex").MaxLen(20).Optional(),
		field.Float("LastYearInventResult").Default(0.0),
		field.Float("LastYearYield").Default(0.0),
		field.Float("Last3YearsInvertResult").Default(0.0),
		field.Float("Last3YearsYield").Default(0.0),
		field.Float("WeekRefillAmount").Positive(),
		field.Float("StartAmount").Positive(),
		field.Time("StartSimulation").GoType(&domains.JSDateOnly{}),
		field.Bool("BuyOnlyLowPrice").Default(false),
		field.Bool("AllowLossWhenSell").Default(true),
		field.Bool("AllowSellToFit").Default(true),
		field.Int("SameEmitent").GoType(domains.SameEmitentPolicy(0)),
	}
}

func (Strategy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).Ref("Strategies").Unique().Required(),
		edge.To("Factors", StrategyFactor.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("Filters", StrategyFilter.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("FixedTickers", StrategyFixedTicker.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

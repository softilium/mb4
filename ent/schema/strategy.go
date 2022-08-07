package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Strategy struct {
	ent.Schema
}

const (
	SameEmitentPolicy_Allow        = 100
	SameEmitentPolicy_PreferPrefs  = 200
	SameEmitentPolicy_PreferOrd    = 300
	SameEmitentPolicy_AllowOnlyOne = 400
)

func (Strategy) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.String("Descr").MaxLen(100).NotEmpty(),
		field.Int("MaxTickers").Range(1, 100).Default(10),
		field.Int("MaxTickersPerIndustry").Range(1, 100).Default(5),
		field.String("BaseIndex").MaxLen(20).Optional(),
		field.Float("LastYearInventResult").Default(0.0),
		field.Float("LastYearYield").Default(0.0),
		field.Float("Last3YearsInvertResult").Default(0.0),
		field.Float("Last3YearsYield").Default(0.0),
		field.Float("WeekRefillAmount").Positive(),
		field.Float("StartAmount").Positive(),
		field.Time("StartSimulation").Default(time.Now()),
		field.Bool("BuyOnlyLowPrice").Default(false),
		field.Bool("AllowLossWhenSell").Default(true),
		field.Int("SameEmitent").Range(SameEmitentPolicy_Allow, SameEmitentPolicy_AllowOnlyOne).Default(SameEmitentPolicy_Allow),
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

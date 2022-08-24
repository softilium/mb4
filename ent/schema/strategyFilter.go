package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/domains"
)

type StrategyFilter struct {
	ent.Schema
}

func (StrategyFilter) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("LineNum").Range(1, 10000).Default(1),
		field.Bool("IsUsed").Default(true),
		field.Int("LeftValueKind").GoType(domains.FilterValueKind(0)),
		field.Int("LeftReportValue").GoType(domains.ReportValue(0)).Default(int(domains.RK_Revenue)),         //applicable when LeftValueKind == FVK_ReportValue
		field.Int("LeftReportValueType").GoType(domains.ReportValueType(0)).Default(int(domains.RVT_AG_Ltm)), //applicable when LeftValueKind == FVK_ReportValue
		field.Int("Operation").GoType(domains.FilterOp(0)),
		field.String("RightValueStr").MaxLen(100).Default(""),
		field.Float("RightValueFloat").Default(0.0),
	}
}

func (StrategyFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Strategy", Strategy.Type).Ref("Filters").Unique().Required(),
	}
}

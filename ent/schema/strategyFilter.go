package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// value kinds
const (
	VK_Ticker      = 100
	VK_Industry    = 200
	VK_ReportValue = 300
	VK_OrGroup     = 400
	VK_AndGroup    = 500
)

// filter operations
const (
	FO_Eq = 10
	FO_Lt = 20
	FO_Le = 30
	FO_Gt = 40
	FO_Ge = 50
	FO_Ne = 60
)

type StrategyFilter struct {
	ent.Schema
}

func (StrategyFilter) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.Int("LineNum").Range(1, 10000).Default(1),
		field.Bool("IsUsed").Default(true),
		field.Int("LeftValueKind").Default(VK_Ticker).Range(VK_Ticker, VK_AndGroup),
		field.String("LeftValue").MaxLen(100).NotEmpty(),
		field.Int("RVT"), // // cube.RVT_* - report value type (YTD, LTM, ...) WHEN FLVK = FLVK_ReportValue
		field.Int("Operation").Default(FO_Eq).Range(FO_Eq, FO_Ne), // Filter Operation
		field.String("RightValue").MaxLen(100).NotEmpty(),
	}
}

func (StrategyFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Strategy", Strategy.Type).Ref("Filters").Unique().Required(),
	}
}

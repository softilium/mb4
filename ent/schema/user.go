package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).MaxLen(20).Immutable().NotEmpty(),
		field.String("UserName").NotEmpty().MinLen(3).MaxLen(50).Unique(),
		field.String("PasswordHash"),
		field.Bool("Admin").Default(false),
		field.Time("StartInvestAccountsFlow").Optional().SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Int("HowManyTickersOnHomepage").Range(10, 100).Default(20),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("InvestAccounts", InvestAccount.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("Strategies", Strategy.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

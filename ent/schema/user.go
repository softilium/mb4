package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

const Auth_Type_email = 1
const Auth_Type_Telegram = 2

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New),
		field.String("UserName").NotEmpty().MinLen(3).MaxLen(50).Unique(),
		field.Int32("AuthType").Immutable().Default(Auth_Type_email).NonNegative(),
		field.String("PasswordHash"),
		field.Bool("Admin").Default(false),
	}
}

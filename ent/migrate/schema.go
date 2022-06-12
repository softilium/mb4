// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EmitentsColumns holds the columns for the "emitents" table.
	EmitentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "descr", Type: field.TypeString, Size: 100},
	}
	// EmitentsTable holds the schema information for the "emitents" table.
	EmitentsTable = &schema.Table{
		Name:       "emitents",
		Columns:    EmitentsColumns,
		PrimaryKey: []*schema.Column{EmitentsColumns[0]},
	}
	// IndustriesColumns holds the columns for the "industries" table.
	IndustriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "descr", Type: field.TypeString, Unique: true, Size: 100},
	}
	// IndustriesTable holds the schema information for the "industries" table.
	IndustriesTable = &schema.Table{
		Name:       "industries",
		Columns:    IndustriesColumns,
		PrimaryKey: []*schema.Column{IndustriesColumns[0]},
	}
	// TickersColumns holds the columns for the "tickers" table.
	TickersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TickersTable holds the schema information for the "tickers" table.
	TickersTable = &schema.Table{
		Name:       "tickers",
		Columns:    TickersColumns,
		PrimaryKey: []*schema.Column{TickersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_name", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "auth_type", Type: field.TypeInt32, Default: 1},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EmitentsTable,
		IndustriesTable,
		TickersTable,
		UsersTable,
	}
)

func init() {
}

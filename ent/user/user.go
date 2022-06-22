// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldPasswordHash holds the string denoting the passwordhash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldStartInvestAccountsFlow holds the string denoting the startinvestaccountsflow field in the database.
	FieldStartInvestAccountsFlow = "start_invest_accounts_flow"
	// EdgeInvestAccounts holds the string denoting the investaccounts edge name in mutations.
	EdgeInvestAccounts = "InvestAccounts"
	// Table holds the table name of the user in the database.
	Table = "users"
	// InvestAccountsTable is the table that holds the InvestAccounts relation/edge.
	InvestAccountsTable = "invest_accounts"
	// InvestAccountsInverseTable is the table name for the InvestAccount entity.
	// It exists in this package in order to avoid circular dependency with the "investaccount" package.
	InvestAccountsInverseTable = "invest_accounts"
	// InvestAccountsColumn is the table column denoting the InvestAccounts relation/edge.
	InvestAccountsColumn = "user_invest_accounts"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserName,
	FieldPasswordHash,
	FieldAdmin,
	FieldStartInvestAccountsFlow,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserNameValidator is a validator for the "UserName" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// DefaultAdmin holds the default value on creation for the "Admin" field.
	DefaultAdmin bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)

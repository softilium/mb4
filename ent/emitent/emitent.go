// Code generated by entc, DO NOT EDIT.

package emitent

import (
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the emitent type in the database.
	Label = "emitent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescr holds the string denoting the descr field in the database.
	FieldDescr = "descr"
	// EdgeIndustry holds the string denoting the industry edge name in mutations.
	EdgeIndustry = "Industry"
	// EdgeTickers holds the string denoting the tickers edge name in mutations.
	EdgeTickers = "Tickers"
	// EdgeReports holds the string denoting the reports edge name in mutations.
	EdgeReports = "Reports"
	// Table holds the table name of the emitent in the database.
	Table = "emitents"
	// IndustryTable is the table that holds the Industry relation/edge.
	IndustryTable = "emitents"
	// IndustryInverseTable is the table name for the Industry entity.
	// It exists in this package in order to avoid circular dependency with the "industry" package.
	IndustryInverseTable = "industries"
	// IndustryColumn is the table column denoting the Industry relation/edge.
	IndustryColumn = "industry_emitents"
	// TickersTable is the table that holds the Tickers relation/edge.
	TickersTable = "tickers"
	// TickersInverseTable is the table name for the Ticker entity.
	// It exists in this package in order to avoid circular dependency with the "ticker" package.
	TickersInverseTable = "tickers"
	// TickersColumn is the table column denoting the Tickers relation/edge.
	TickersColumn = "emitent_tickers"
	// ReportsTable is the table that holds the Reports relation/edge.
	ReportsTable = "reports"
	// ReportsInverseTable is the table name for the Report entity.
	// It exists in this package in order to avoid circular dependency with the "report" package.
	ReportsInverseTable = "reports"
	// ReportsColumn is the table column denoting the Reports relation/edge.
	ReportsColumn = "emitent_reports"
)

// Columns holds all SQL columns for emitent fields.
var Columns = []string{
	FieldID,
	FieldDescr,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "emitents"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"industry_emitents",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	DescrValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

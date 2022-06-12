// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/softilium/mb4/ent/ticker"
)

// Ticker is the model entity for the Ticker schema.
type Ticker struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ticker) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticker.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ticker", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ticker fields.
func (t *Ticker) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticker.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Ticker.
// Note that you need to call Ticker.Unwrap() before calling this method if this Ticker
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Ticker) Update() *TickerUpdateOne {
	return (&TickerClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Ticker entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Ticker) Unwrap() *Ticker {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ticker is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Ticker) String() string {
	var builder strings.Builder
	builder.WriteString("Ticker(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Tickers is a parsable slice of Ticker.
type Tickers []*Ticker

func (t Tickers) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}

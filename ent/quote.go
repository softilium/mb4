// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/quote"
	"github.com/softilium/mb4/ent/ticker"
)

// Quote is the model entity for the Quote schema.
type Quote struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// D holds the value of the "D" field.
	D time.Time `json:"D,omitempty"`
	// O holds the value of the "O" field.
	O float64 `json:"O,omitempty"`
	// C holds the value of the "C" field.
	C float64 `json:"C,omitempty"`
	// H holds the value of the "H" field.
	H float64 `json:"H,omitempty"`
	// L holds the value of the "L" field.
	L float64 `json:"L,omitempty"`
	// V holds the value of the "V" field.
	V float64 `json:"V,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuoteQuery when eager-loading is set.
	Edges         QuoteEdges `json:"edges"`
	ticker_quotes *string
}

// QuoteEdges holds the relations/edges for other nodes in the graph.
type QuoteEdges struct {
	// Ticker holds the value of the Ticker edge.
	Ticker *Ticker `json:"Ticker,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TickerOrErr returns the Ticker value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuoteEdges) TickerOrErr() (*Ticker, error) {
	if e.loadedTypes[0] {
		if e.Ticker == nil {
			// The edge Ticker was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ticker.Label}
		}
		return e.Ticker, nil
	}
	return nil, &NotLoadedError{edge: "Ticker"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Quote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case quote.FieldO, quote.FieldC, quote.FieldH, quote.FieldL, quote.FieldV:
			values[i] = new(sql.NullFloat64)
		case quote.FieldD:
			values[i] = new(sql.NullTime)
		case quote.FieldID:
			values[i] = new(xid.ID)
		case quote.ForeignKeys[0]: // ticker_quotes
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Quote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Quote fields.
func (q *Quote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case quote.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				q.ID = *value
			}
		case quote.FieldD:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field D", values[i])
			} else if value.Valid {
				q.D = value.Time
			}
		case quote.FieldO:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field O", values[i])
			} else if value.Valid {
				q.O = value.Float64
			}
		case quote.FieldC:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field C", values[i])
			} else if value.Valid {
				q.C = value.Float64
			}
		case quote.FieldH:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field H", values[i])
			} else if value.Valid {
				q.H = value.Float64
			}
		case quote.FieldL:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field L", values[i])
			} else if value.Valid {
				q.L = value.Float64
			}
		case quote.FieldV:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field V", values[i])
			} else if value.Valid {
				q.V = value.Float64
			}
		case quote.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ticker_quotes", values[i])
			} else if value.Valid {
				q.ticker_quotes = new(string)
				*q.ticker_quotes = value.String
			}
		}
	}
	return nil
}

// QueryTicker queries the "Ticker" edge of the Quote entity.
func (q *Quote) QueryTicker() *TickerQuery {
	return (&QuoteClient{config: q.config}).QueryTicker(q)
}

// Update returns a builder for updating this Quote.
// Note that you need to call Quote.Unwrap() before calling this method if this Quote
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Quote) Update() *QuoteUpdateOne {
	return (&QuoteClient{config: q.config}).UpdateOne(q)
}

// Unwrap unwraps the Quote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Quote) Unwrap() *Quote {
	tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Quote is not a transactional entity")
	}
	q.config.driver = tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Quote) String() string {
	var builder strings.Builder
	builder.WriteString("Quote(")
	builder.WriteString(fmt.Sprintf("id=%v", q.ID))
	builder.WriteString(", D=")
	builder.WriteString(q.D.Format(time.ANSIC))
	builder.WriteString(", O=")
	builder.WriteString(fmt.Sprintf("%v", q.O))
	builder.WriteString(", C=")
	builder.WriteString(fmt.Sprintf("%v", q.C))
	builder.WriteString(", H=")
	builder.WriteString(fmt.Sprintf("%v", q.H))
	builder.WriteString(", L=")
	builder.WriteString(fmt.Sprintf("%v", q.L))
	builder.WriteString(", V=")
	builder.WriteString(fmt.Sprintf("%v", q.V))
	builder.WriteByte(')')
	return builder.String()
}

// Quotes is a parsable slice of Quote.
type Quotes []*Quote

func (q Quotes) config(cfg config) {
	for _i := range q {
		q[_i].config = cfg
	}
}

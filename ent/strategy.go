// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
	"github.com/softilium/mb4/domains"
	"github.com/softilium/mb4/ent/strategy"
	"github.com/softilium/mb4/ent/user"
)

// Strategy is the model entity for the Strategy schema.
type Strategy struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Descr holds the value of the "Descr" field.
	Descr string `json:"Descr,omitempty"`
	// MaxTickers holds the value of the "MaxTickers" field.
	MaxTickers int `json:"MaxTickers,omitempty"`
	// MaxTickersPerIndustry holds the value of the "MaxTickersPerIndustry" field.
	MaxTickersPerIndustry int `json:"MaxTickersPerIndustry,omitempty"`
	// BaseIndex holds the value of the "BaseIndex" field.
	BaseIndex string `json:"BaseIndex,omitempty"`
	// LastYearInventResult holds the value of the "LastYearInventResult" field.
	LastYearInventResult float64 `json:"LastYearInventResult,omitempty"`
	// LastYearYield holds the value of the "LastYearYield" field.
	LastYearYield float64 `json:"LastYearYield,omitempty"`
	// Last3YearsInvertResult holds the value of the "Last3YearsInvertResult" field.
	Last3YearsInvertResult float64 `json:"Last3YearsInvertResult,omitempty"`
	// Last3YearsYield holds the value of the "Last3YearsYield" field.
	Last3YearsYield float64 `json:"Last3YearsYield,omitempty"`
	// WeekRefillAmount holds the value of the "WeekRefillAmount" field.
	WeekRefillAmount float64 `json:"WeekRefillAmount,omitempty"`
	// StartAmount holds the value of the "StartAmount" field.
	StartAmount float64 `json:"StartAmount,omitempty"`
	// StartSimulation holds the value of the "StartSimulation" field.
	StartSimulation *domains.JSDateOnly `json:"StartSimulation,omitempty"`
	// BuyOnlyLowPrice holds the value of the "BuyOnlyLowPrice" field.
	BuyOnlyLowPrice bool `json:"BuyOnlyLowPrice,omitempty"`
	// AllowLossWhenSell holds the value of the "AllowLossWhenSell" field.
	AllowLossWhenSell bool `json:"AllowLossWhenSell,omitempty"`
	// AllowSellToFit holds the value of the "AllowSellToFit" field.
	AllowSellToFit bool `json:"AllowSellToFit,omitempty"`
	// SameEmitent holds the value of the "SameEmitent" field.
	SameEmitent int `json:"SameEmitent,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StrategyQuery when eager-loading is set.
	Edges           StrategyEdges `json:"edges"`
	user_strategies *xid.ID
}

// StrategyEdges holds the relations/edges for other nodes in the graph.
type StrategyEdges struct {
	// User holds the value of the User edge.
	User *User `json:"User,omitempty"`
	// Factors holds the value of the Factors edge.
	Factors []*StrategyFactor `json:"Factors,omitempty"`
	// Filters holds the value of the Filters edge.
	Filters []*StrategyFilter `json:"Filters,omitempty"`
	// FixedTickers holds the value of the FixedTickers edge.
	FixedTickers []*StrategyFixedTicker `json:"FixedTickers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StrategyEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge User was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "User"}
}

// FactorsOrErr returns the Factors value or an error if the edge
// was not loaded in eager-loading.
func (e StrategyEdges) FactorsOrErr() ([]*StrategyFactor, error) {
	if e.loadedTypes[1] {
		return e.Factors, nil
	}
	return nil, &NotLoadedError{edge: "Factors"}
}

// FiltersOrErr returns the Filters value or an error if the edge
// was not loaded in eager-loading.
func (e StrategyEdges) FiltersOrErr() ([]*StrategyFilter, error) {
	if e.loadedTypes[2] {
		return e.Filters, nil
	}
	return nil, &NotLoadedError{edge: "Filters"}
}

// FixedTickersOrErr returns the FixedTickers value or an error if the edge
// was not loaded in eager-loading.
func (e StrategyEdges) FixedTickersOrErr() ([]*StrategyFixedTicker, error) {
	if e.loadedTypes[3] {
		return e.FixedTickers, nil
	}
	return nil, &NotLoadedError{edge: "FixedTickers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Strategy) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case strategy.FieldStartSimulation:
			values[i] = new(domains.JSDateOnly)
		case strategy.FieldBuyOnlyLowPrice, strategy.FieldAllowLossWhenSell, strategy.FieldAllowSellToFit:
			values[i] = new(sql.NullBool)
		case strategy.FieldLastYearInventResult, strategy.FieldLastYearYield, strategy.FieldLast3YearsInvertResult, strategy.FieldLast3YearsYield, strategy.FieldWeekRefillAmount, strategy.FieldStartAmount:
			values[i] = new(sql.NullFloat64)
		case strategy.FieldMaxTickers, strategy.FieldMaxTickersPerIndustry, strategy.FieldSameEmitent:
			values[i] = new(sql.NullInt64)
		case strategy.FieldDescr, strategy.FieldBaseIndex:
			values[i] = new(sql.NullString)
		case strategy.FieldID:
			values[i] = new(xid.ID)
		case strategy.ForeignKeys[0]: // user_strategies
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Strategy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Strategy fields.
func (s *Strategy) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case strategy.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case strategy.FieldDescr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Descr", values[i])
			} else if value.Valid {
				s.Descr = value.String
			}
		case strategy.FieldMaxTickers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field MaxTickers", values[i])
			} else if value.Valid {
				s.MaxTickers = int(value.Int64)
			}
		case strategy.FieldMaxTickersPerIndustry:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field MaxTickersPerIndustry", values[i])
			} else if value.Valid {
				s.MaxTickersPerIndustry = int(value.Int64)
			}
		case strategy.FieldBaseIndex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BaseIndex", values[i])
			} else if value.Valid {
				s.BaseIndex = value.String
			}
		case strategy.FieldLastYearInventResult:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field LastYearInventResult", values[i])
			} else if value.Valid {
				s.LastYearInventResult = value.Float64
			}
		case strategy.FieldLastYearYield:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field LastYearYield", values[i])
			} else if value.Valid {
				s.LastYearYield = value.Float64
			}
		case strategy.FieldLast3YearsInvertResult:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Last3YearsInvertResult", values[i])
			} else if value.Valid {
				s.Last3YearsInvertResult = value.Float64
			}
		case strategy.FieldLast3YearsYield:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Last3YearsYield", values[i])
			} else if value.Valid {
				s.Last3YearsYield = value.Float64
			}
		case strategy.FieldWeekRefillAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field WeekRefillAmount", values[i])
			} else if value.Valid {
				s.WeekRefillAmount = value.Float64
			}
		case strategy.FieldStartAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field StartAmount", values[i])
			} else if value.Valid {
				s.StartAmount = value.Float64
			}
		case strategy.FieldStartSimulation:
			if value, ok := values[i].(*domains.JSDateOnly); !ok {
				return fmt.Errorf("unexpected type %T for field StartSimulation", values[i])
			} else if value != nil {
				s.StartSimulation = value
			}
		case strategy.FieldBuyOnlyLowPrice:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field BuyOnlyLowPrice", values[i])
			} else if value.Valid {
				s.BuyOnlyLowPrice = value.Bool
			}
		case strategy.FieldAllowLossWhenSell:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field AllowLossWhenSell", values[i])
			} else if value.Valid {
				s.AllowLossWhenSell = value.Bool
			}
		case strategy.FieldAllowSellToFit:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field AllowSellToFit", values[i])
			} else if value.Valid {
				s.AllowSellToFit = value.Bool
			}
		case strategy.FieldSameEmitent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field SameEmitent", values[i])
			} else if value.Valid {
				s.SameEmitent = int(value.Int64)
			}
		case strategy.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_strategies", values[i])
			} else if value.Valid {
				s.user_strategies = new(xid.ID)
				*s.user_strategies = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryUser queries the "User" edge of the Strategy entity.
func (s *Strategy) QueryUser() *UserQuery {
	return (&StrategyClient{config: s.config}).QueryUser(s)
}

// QueryFactors queries the "Factors" edge of the Strategy entity.
func (s *Strategy) QueryFactors() *StrategyFactorQuery {
	return (&StrategyClient{config: s.config}).QueryFactors(s)
}

// QueryFilters queries the "Filters" edge of the Strategy entity.
func (s *Strategy) QueryFilters() *StrategyFilterQuery {
	return (&StrategyClient{config: s.config}).QueryFilters(s)
}

// QueryFixedTickers queries the "FixedTickers" edge of the Strategy entity.
func (s *Strategy) QueryFixedTickers() *StrategyFixedTickerQuery {
	return (&StrategyClient{config: s.config}).QueryFixedTickers(s)
}

// Update returns a builder for updating this Strategy.
// Note that you need to call Strategy.Unwrap() before calling this method if this Strategy
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Strategy) Update() *StrategyUpdateOne {
	return (&StrategyClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Strategy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Strategy) Unwrap() *Strategy {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Strategy is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Strategy) String() string {
	var builder strings.Builder
	builder.WriteString("Strategy(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Descr=")
	builder.WriteString(s.Descr)
	builder.WriteString(", MaxTickers=")
	builder.WriteString(fmt.Sprintf("%v", s.MaxTickers))
	builder.WriteString(", MaxTickersPerIndustry=")
	builder.WriteString(fmt.Sprintf("%v", s.MaxTickersPerIndustry))
	builder.WriteString(", BaseIndex=")
	builder.WriteString(s.BaseIndex)
	builder.WriteString(", LastYearInventResult=")
	builder.WriteString(fmt.Sprintf("%v", s.LastYearInventResult))
	builder.WriteString(", LastYearYield=")
	builder.WriteString(fmt.Sprintf("%v", s.LastYearYield))
	builder.WriteString(", Last3YearsInvertResult=")
	builder.WriteString(fmt.Sprintf("%v", s.Last3YearsInvertResult))
	builder.WriteString(", Last3YearsYield=")
	builder.WriteString(fmt.Sprintf("%v", s.Last3YearsYield))
	builder.WriteString(", WeekRefillAmount=")
	builder.WriteString(fmt.Sprintf("%v", s.WeekRefillAmount))
	builder.WriteString(", StartAmount=")
	builder.WriteString(fmt.Sprintf("%v", s.StartAmount))
	builder.WriteString(", StartSimulation=")
	builder.WriteString(fmt.Sprintf("%v", s.StartSimulation))
	builder.WriteString(", BuyOnlyLowPrice=")
	builder.WriteString(fmt.Sprintf("%v", s.BuyOnlyLowPrice))
	builder.WriteString(", AllowLossWhenSell=")
	builder.WriteString(fmt.Sprintf("%v", s.AllowLossWhenSell))
	builder.WriteString(", AllowSellToFit=")
	builder.WriteString(fmt.Sprintf("%v", s.AllowSellToFit))
	builder.WriteString(", SameEmitent=")
	builder.WriteString(fmt.Sprintf("%v", s.SameEmitent))
	builder.WriteByte(')')
	return builder.String()
}

// Strategies is a parsable slice of Strategy.
type Strategies []*Strategy

func (s Strategies) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}

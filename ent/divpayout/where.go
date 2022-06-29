// Code generated by entc, DO NOT EDIT.

package divpayout

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/softilium/mb4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ForYear applies equality check predicate on the "ForYear" field. It's identical to ForYearEQ.
func ForYear(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForYear), v))
	})
}

// ForQuarter applies equality check predicate on the "ForQuarter" field. It's identical to ForQuarterEQ.
func ForQuarter(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForQuarter), v))
	})
}

// CloseDate applies equality check predicate on the "CloseDate" field. It's identical to CloseDateEQ.
func CloseDate(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCloseDate), v))
	})
}

// Status applies equality check predicate on the "Status" field. It's identical to StatusEQ.
func Status(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// DPS applies equality check predicate on the "DPS" field. It's identical to DPSEQ.
func DPS(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDPS), v))
	})
}

// ForYearEQ applies the EQ predicate on the "ForYear" field.
func ForYearEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForYear), v))
	})
}

// ForYearNEQ applies the NEQ predicate on the "ForYear" field.
func ForYearNEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForYear), v))
	})
}

// ForYearIn applies the In predicate on the "ForYear" field.
func ForYearIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldForYear), v...))
	})
}

// ForYearNotIn applies the NotIn predicate on the "ForYear" field.
func ForYearNotIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldForYear), v...))
	})
}

// ForYearGT applies the GT predicate on the "ForYear" field.
func ForYearGT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldForYear), v))
	})
}

// ForYearGTE applies the GTE predicate on the "ForYear" field.
func ForYearGTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldForYear), v))
	})
}

// ForYearLT applies the LT predicate on the "ForYear" field.
func ForYearLT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldForYear), v))
	})
}

// ForYearLTE applies the LTE predicate on the "ForYear" field.
func ForYearLTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldForYear), v))
	})
}

// ForQuarterEQ applies the EQ predicate on the "ForQuarter" field.
func ForQuarterEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldForQuarter), v))
	})
}

// ForQuarterNEQ applies the NEQ predicate on the "ForQuarter" field.
func ForQuarterNEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldForQuarter), v))
	})
}

// ForQuarterIn applies the In predicate on the "ForQuarter" field.
func ForQuarterIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldForQuarter), v...))
	})
}

// ForQuarterNotIn applies the NotIn predicate on the "ForQuarter" field.
func ForQuarterNotIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldForQuarter), v...))
	})
}

// ForQuarterGT applies the GT predicate on the "ForQuarter" field.
func ForQuarterGT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldForQuarter), v))
	})
}

// ForQuarterGTE applies the GTE predicate on the "ForQuarter" field.
func ForQuarterGTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldForQuarter), v))
	})
}

// ForQuarterLT applies the LT predicate on the "ForQuarter" field.
func ForQuarterLT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldForQuarter), v))
	})
}

// ForQuarterLTE applies the LTE predicate on the "ForQuarter" field.
func ForQuarterLTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldForQuarter), v))
	})
}

// CloseDateEQ applies the EQ predicate on the "CloseDate" field.
func CloseDateEQ(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCloseDate), v))
	})
}

// CloseDateNEQ applies the NEQ predicate on the "CloseDate" field.
func CloseDateNEQ(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCloseDate), v))
	})
}

// CloseDateIn applies the In predicate on the "CloseDate" field.
func CloseDateIn(vs ...time.Time) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCloseDate), v...))
	})
}

// CloseDateNotIn applies the NotIn predicate on the "CloseDate" field.
func CloseDateNotIn(vs ...time.Time) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCloseDate), v...))
	})
}

// CloseDateGT applies the GT predicate on the "CloseDate" field.
func CloseDateGT(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCloseDate), v))
	})
}

// CloseDateGTE applies the GTE predicate on the "CloseDate" field.
func CloseDateGTE(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCloseDate), v))
	})
}

// CloseDateLT applies the LT predicate on the "CloseDate" field.
func CloseDateLT(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCloseDate), v))
	})
}

// CloseDateLTE applies the LTE predicate on the "CloseDate" field.
func CloseDateLTE(v time.Time) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCloseDate), v))
	})
}

// StatusEQ applies the EQ predicate on the "Status" field.
func StatusEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "Status" field.
func StatusNEQ(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "Status" field.
func StatusIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "Status" field.
func StatusNotIn(vs ...int) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "Status" field.
func StatusGT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "Status" field.
func StatusGTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "Status" field.
func StatusLT(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "Status" field.
func StatusLTE(v int) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// DPSEQ applies the EQ predicate on the "DPS" field.
func DPSEQ(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDPS), v))
	})
}

// DPSNEQ applies the NEQ predicate on the "DPS" field.
func DPSNEQ(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDPS), v))
	})
}

// DPSIn applies the In predicate on the "DPS" field.
func DPSIn(vs ...float64) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDPS), v...))
	})
}

// DPSNotIn applies the NotIn predicate on the "DPS" field.
func DPSNotIn(vs ...float64) predicate.DivPayout {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DivPayout(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDPS), v...))
	})
}

// DPSGT applies the GT predicate on the "DPS" field.
func DPSGT(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDPS), v))
	})
}

// DPSGTE applies the GTE predicate on the "DPS" field.
func DPSGTE(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDPS), v))
	})
}

// DPSLT applies the LT predicate on the "DPS" field.
func DPSLT(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDPS), v))
	})
}

// DPSLTE applies the LTE predicate on the "DPS" field.
func DPSLTE(v float64) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDPS), v))
	})
}

// HasTickers applies the HasEdge predicate on the "Tickers" edge.
func HasTickers() predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TickersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TickersTable, TickersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTickersWith applies the HasEdge predicate on the "Tickers" edge with a given conditions (other predicates).
func HasTickersWith(preds ...predicate.Ticker) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TickersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TickersTable, TickersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DivPayout) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DivPayout) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DivPayout) predicate.DivPayout {
	return predicate.DivPayout(func(s *sql.Selector) {
		p(s.Not())
	})
}
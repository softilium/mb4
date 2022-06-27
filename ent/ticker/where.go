// Code generated by entc, DO NOT EDIT.

package ticker

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/softilium/mb4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Descr applies equality check predicate on the "Descr" field. It's identical to DescrEQ.
func Descr(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescr), v))
	})
}

// Kind applies equality check predicate on the "Kind" field. It's identical to KindEQ.
func Kind(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// DescrEQ applies the EQ predicate on the "Descr" field.
func DescrEQ(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescr), v))
	})
}

// DescrNEQ applies the NEQ predicate on the "Descr" field.
func DescrNEQ(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescr), v))
	})
}

// DescrIn applies the In predicate on the "Descr" field.
func DescrIn(vs ...string) predicate.Ticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescr), v...))
	})
}

// DescrNotIn applies the NotIn predicate on the "Descr" field.
func DescrNotIn(vs ...string) predicate.Ticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescr), v...))
	})
}

// DescrGT applies the GT predicate on the "Descr" field.
func DescrGT(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescr), v))
	})
}

// DescrGTE applies the GTE predicate on the "Descr" field.
func DescrGTE(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescr), v))
	})
}

// DescrLT applies the LT predicate on the "Descr" field.
func DescrLT(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescr), v))
	})
}

// DescrLTE applies the LTE predicate on the "Descr" field.
func DescrLTE(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescr), v))
	})
}

// DescrContains applies the Contains predicate on the "Descr" field.
func DescrContains(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescr), v))
	})
}

// DescrHasPrefix applies the HasPrefix predicate on the "Descr" field.
func DescrHasPrefix(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescr), v))
	})
}

// DescrHasSuffix applies the HasSuffix predicate on the "Descr" field.
func DescrHasSuffix(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescr), v))
	})
}

// DescrEqualFold applies the EqualFold predicate on the "Descr" field.
func DescrEqualFold(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescr), v))
	})
}

// DescrContainsFold applies the ContainsFold predicate on the "Descr" field.
func DescrContainsFold(v string) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescr), v))
	})
}

// KindEQ applies the EQ predicate on the "Kind" field.
func KindEQ(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// KindNEQ applies the NEQ predicate on the "Kind" field.
func KindNEQ(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKind), v))
	})
}

// KindIn applies the In predicate on the "Kind" field.
func KindIn(vs ...int32) predicate.Ticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKind), v...))
	})
}

// KindNotIn applies the NotIn predicate on the "Kind" field.
func KindNotIn(vs ...int32) predicate.Ticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKind), v...))
	})
}

// KindGT applies the GT predicate on the "Kind" field.
func KindGT(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKind), v))
	})
}

// KindGTE applies the GTE predicate on the "Kind" field.
func KindGTE(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKind), v))
	})
}

// KindLT applies the LT predicate on the "Kind" field.
func KindLT(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKind), v))
	})
}

// KindLTE applies the LTE predicate on the "Kind" field.
func KindLTE(v int32) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKind), v))
	})
}

// HasEmitent applies the HasEdge predicate on the "Emitent" edge.
func HasEmitent() predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmitentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmitentTable, EmitentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmitentWith applies the HasEdge predicate on the "Emitent" edge with a given conditions (other predicates).
func HasEmitentWith(preds ...predicate.Emitent) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmitentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmitentTable, EmitentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuotes applies the HasEdge predicate on the "Quotes" edge.
func HasQuotes() predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuotesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuotesTable, QuotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuotesWith applies the HasEdge predicate on the "Quotes" edge with a given conditions (other predicates).
func HasQuotesWith(preds ...predicate.Quote) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuotesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuotesTable, QuotesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDivPayouts applies the HasEdge predicate on the "DivPayouts" edge.
func HasDivPayouts() predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DivPayoutsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DivPayoutsTable, DivPayoutsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDivPayoutsWith applies the HasEdge predicate on the "DivPayouts" edge with a given conditions (other predicates).
func HasDivPayoutsWith(preds ...predicate.DivPayout) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DivPayoutsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DivPayoutsTable, DivPayoutsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
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
func Not(p predicate.Ticker) predicate.Ticker {
	return predicate.Ticker(func(s *sql.Selector) {
		p(s.Not())
	})
}

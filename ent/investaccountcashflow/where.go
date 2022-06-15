// Code generated by entc, DO NOT EDIT.

package investaccountcashflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
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
func IDNotIn(ids ...xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
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
func IDGT(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RecDate applies equality check predicate on the "RecDate" field. It's identical to RecDateEQ.
func RecDate(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecDate), v))
	})
}

// Qty applies equality check predicate on the "Qty" field. It's identical to QtyEQ.
func Qty(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQty), v))
	})
}

// RecDateEQ applies the EQ predicate on the "RecDate" field.
func RecDateEQ(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecDate), v))
	})
}

// RecDateNEQ applies the NEQ predicate on the "RecDate" field.
func RecDateNEQ(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecDate), v))
	})
}

// RecDateIn applies the In predicate on the "RecDate" field.
func RecDateIn(vs ...time.Time) predicate.InvestAccountCashflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecDate), v...))
	})
}

// RecDateNotIn applies the NotIn predicate on the "RecDate" field.
func RecDateNotIn(vs ...time.Time) predicate.InvestAccountCashflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecDate), v...))
	})
}

// RecDateGT applies the GT predicate on the "RecDate" field.
func RecDateGT(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecDate), v))
	})
}

// RecDateGTE applies the GTE predicate on the "RecDate" field.
func RecDateGTE(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecDate), v))
	})
}

// RecDateLT applies the LT predicate on the "RecDate" field.
func RecDateLT(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecDate), v))
	})
}

// RecDateLTE applies the LTE predicate on the "RecDate" field.
func RecDateLTE(v time.Time) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecDate), v))
	})
}

// QtyEQ applies the EQ predicate on the "Qty" field.
func QtyEQ(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQty), v))
	})
}

// QtyNEQ applies the NEQ predicate on the "Qty" field.
func QtyNEQ(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQty), v))
	})
}

// QtyIn applies the In predicate on the "Qty" field.
func QtyIn(vs ...float64) predicate.InvestAccountCashflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQty), v...))
	})
}

// QtyNotIn applies the NotIn predicate on the "Qty" field.
func QtyNotIn(vs ...float64) predicate.InvestAccountCashflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQty), v...))
	})
}

// QtyGT applies the GT predicate on the "Qty" field.
func QtyGT(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQty), v))
	})
}

// QtyGTE applies the GTE predicate on the "Qty" field.
func QtyGTE(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQty), v))
	})
}

// QtyLT applies the LT predicate on the "Qty" field.
func QtyLT(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQty), v))
	})
}

// QtyLTE applies the LTE predicate on the "Qty" field.
func QtyLTE(v float64) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQty), v))
	})
}

// HasOwner applies the HasEdge predicate on the "Owner" edge.
func HasOwner() predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "Owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.InvestAccount) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InvestAccountCashflow) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InvestAccountCashflow) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
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
func Not(p predicate.InvestAccountCashflow) predicate.InvestAccountCashflow {
	return predicate.InvestAccountCashflow(func(s *sql.Selector) {
		p(s.Not())
	})
}

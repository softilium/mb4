// Code generated by entc, DO NOT EDIT.

package strategyfactor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
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
func IDNotIn(ids ...xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
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
func IDGT(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LineNum applies equality check predicate on the "LineNum" field. It's identical to LineNumEQ.
func LineNum(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLineNum), v))
	})
}

// IsUsed applies equality check predicate on the "IsUsed" field. It's identical to IsUsedEQ.
func IsUsed(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsUsed), v))
	})
}

// RK applies equality check predicate on the "RK" field. It's identical to RKEQ.
func RK(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRK), v))
	})
}

// RVT applies equality check predicate on the "RVT" field. It's identical to RVTEQ.
func RVT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRVT), v))
	})
}

// MinAcceptabe applies equality check predicate on the "MinAcceptabe" field. It's identical to MinAcceptabeEQ.
func MinAcceptabe(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinAcceptabe), v))
	})
}

// MaxAcceptable applies equality check predicate on the "MaxAcceptable" field. It's identical to MaxAcceptableEQ.
func MaxAcceptable(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxAcceptable), v))
	})
}

// Inverse applies equality check predicate on the "Inverse" field. It's identical to InverseEQ.
func Inverse(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInverse), v))
	})
}

// K applies equality check predicate on the "K" field. It's identical to KEQ.
func K(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK), v))
	})
}

// Gist applies equality check predicate on the "Gist" field. It's identical to GistEQ.
func Gist(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGist), v))
	})
}

// LineNumEQ applies the EQ predicate on the "LineNum" field.
func LineNumEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLineNum), v))
	})
}

// LineNumNEQ applies the NEQ predicate on the "LineNum" field.
func LineNumNEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLineNum), v))
	})
}

// LineNumIn applies the In predicate on the "LineNum" field.
func LineNumIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLineNum), v...))
	})
}

// LineNumNotIn applies the NotIn predicate on the "LineNum" field.
func LineNumNotIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLineNum), v...))
	})
}

// LineNumGT applies the GT predicate on the "LineNum" field.
func LineNumGT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLineNum), v))
	})
}

// LineNumGTE applies the GTE predicate on the "LineNum" field.
func LineNumGTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLineNum), v))
	})
}

// LineNumLT applies the LT predicate on the "LineNum" field.
func LineNumLT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLineNum), v))
	})
}

// LineNumLTE applies the LTE predicate on the "LineNum" field.
func LineNumLTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLineNum), v))
	})
}

// IsUsedEQ applies the EQ predicate on the "IsUsed" field.
func IsUsedEQ(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsUsed), v))
	})
}

// IsUsedNEQ applies the NEQ predicate on the "IsUsed" field.
func IsUsedNEQ(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsUsed), v))
	})
}

// RKEQ applies the EQ predicate on the "RK" field.
func RKEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRK), v))
	})
}

// RKNEQ applies the NEQ predicate on the "RK" field.
func RKNEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRK), v))
	})
}

// RKIn applies the In predicate on the "RK" field.
func RKIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRK), v...))
	})
}

// RKNotIn applies the NotIn predicate on the "RK" field.
func RKNotIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRK), v...))
	})
}

// RKGT applies the GT predicate on the "RK" field.
func RKGT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRK), v))
	})
}

// RKGTE applies the GTE predicate on the "RK" field.
func RKGTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRK), v))
	})
}

// RKLT applies the LT predicate on the "RK" field.
func RKLT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRK), v))
	})
}

// RKLTE applies the LTE predicate on the "RK" field.
func RKLTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRK), v))
	})
}

// RVTEQ applies the EQ predicate on the "RVT" field.
func RVTEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRVT), v))
	})
}

// RVTNEQ applies the NEQ predicate on the "RVT" field.
func RVTNEQ(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRVT), v))
	})
}

// RVTIn applies the In predicate on the "RVT" field.
func RVTIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRVT), v...))
	})
}

// RVTNotIn applies the NotIn predicate on the "RVT" field.
func RVTNotIn(vs ...int) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRVT), v...))
	})
}

// RVTGT applies the GT predicate on the "RVT" field.
func RVTGT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRVT), v))
	})
}

// RVTGTE applies the GTE predicate on the "RVT" field.
func RVTGTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRVT), v))
	})
}

// RVTLT applies the LT predicate on the "RVT" field.
func RVTLT(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRVT), v))
	})
}

// RVTLTE applies the LTE predicate on the "RVT" field.
func RVTLTE(v int) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRVT), v))
	})
}

// MinAcceptabeEQ applies the EQ predicate on the "MinAcceptabe" field.
func MinAcceptabeEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinAcceptabe), v))
	})
}

// MinAcceptabeNEQ applies the NEQ predicate on the "MinAcceptabe" field.
func MinAcceptabeNEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinAcceptabe), v))
	})
}

// MinAcceptabeIn applies the In predicate on the "MinAcceptabe" field.
func MinAcceptabeIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMinAcceptabe), v...))
	})
}

// MinAcceptabeNotIn applies the NotIn predicate on the "MinAcceptabe" field.
func MinAcceptabeNotIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMinAcceptabe), v...))
	})
}

// MinAcceptabeGT applies the GT predicate on the "MinAcceptabe" field.
func MinAcceptabeGT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinAcceptabe), v))
	})
}

// MinAcceptabeGTE applies the GTE predicate on the "MinAcceptabe" field.
func MinAcceptabeGTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinAcceptabe), v))
	})
}

// MinAcceptabeLT applies the LT predicate on the "MinAcceptabe" field.
func MinAcceptabeLT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinAcceptabe), v))
	})
}

// MinAcceptabeLTE applies the LTE predicate on the "MinAcceptabe" field.
func MinAcceptabeLTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinAcceptabe), v))
	})
}

// MaxAcceptableEQ applies the EQ predicate on the "MaxAcceptable" field.
func MaxAcceptableEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxAcceptable), v))
	})
}

// MaxAcceptableNEQ applies the NEQ predicate on the "MaxAcceptable" field.
func MaxAcceptableNEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxAcceptable), v))
	})
}

// MaxAcceptableIn applies the In predicate on the "MaxAcceptable" field.
func MaxAcceptableIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxAcceptable), v...))
	})
}

// MaxAcceptableNotIn applies the NotIn predicate on the "MaxAcceptable" field.
func MaxAcceptableNotIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxAcceptable), v...))
	})
}

// MaxAcceptableGT applies the GT predicate on the "MaxAcceptable" field.
func MaxAcceptableGT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxAcceptable), v))
	})
}

// MaxAcceptableGTE applies the GTE predicate on the "MaxAcceptable" field.
func MaxAcceptableGTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxAcceptable), v))
	})
}

// MaxAcceptableLT applies the LT predicate on the "MaxAcceptable" field.
func MaxAcceptableLT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxAcceptable), v))
	})
}

// MaxAcceptableLTE applies the LTE predicate on the "MaxAcceptable" field.
func MaxAcceptableLTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxAcceptable), v))
	})
}

// InverseEQ applies the EQ predicate on the "Inverse" field.
func InverseEQ(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInverse), v))
	})
}

// InverseNEQ applies the NEQ predicate on the "Inverse" field.
func InverseNEQ(v bool) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInverse), v))
	})
}

// KEQ applies the EQ predicate on the "K" field.
func KEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldK), v))
	})
}

// KNEQ applies the NEQ predicate on the "K" field.
func KNEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldK), v))
	})
}

// KIn applies the In predicate on the "K" field.
func KIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldK), v...))
	})
}

// KNotIn applies the NotIn predicate on the "K" field.
func KNotIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldK), v...))
	})
}

// KGT applies the GT predicate on the "K" field.
func KGT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldK), v))
	})
}

// KGTE applies the GTE predicate on the "K" field.
func KGTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldK), v))
	})
}

// KLT applies the LT predicate on the "K" field.
func KLT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldK), v))
	})
}

// KLTE applies the LTE predicate on the "K" field.
func KLTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldK), v))
	})
}

// GistEQ applies the EQ predicate on the "Gist" field.
func GistEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGist), v))
	})
}

// GistNEQ applies the NEQ predicate on the "Gist" field.
func GistNEQ(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGist), v))
	})
}

// GistIn applies the In predicate on the "Gist" field.
func GistIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGist), v...))
	})
}

// GistNotIn applies the NotIn predicate on the "Gist" field.
func GistNotIn(vs ...float64) predicate.StrategyFactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StrategyFactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGist), v...))
	})
}

// GistGT applies the GT predicate on the "Gist" field.
func GistGT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGist), v))
	})
}

// GistGTE applies the GTE predicate on the "Gist" field.
func GistGTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGist), v))
	})
}

// GistLT applies the LT predicate on the "Gist" field.
func GistLT(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGist), v))
	})
}

// GistLTE applies the LTE predicate on the "Gist" field.
func GistLTE(v float64) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGist), v))
	})
}

// HasStrategy applies the HasEdge predicate on the "Strategy" edge.
func HasStrategy() predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StrategyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StrategyTable, StrategyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStrategyWith applies the HasEdge predicate on the "Strategy" edge with a given conditions (other predicates).
func HasStrategyWith(preds ...predicate.Strategy) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StrategyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StrategyTable, StrategyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StrategyFactor) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StrategyFactor) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
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
func Not(p predicate.StrategyFactor) predicate.StrategyFactor {
	return predicate.StrategyFactor(func(s *sql.Selector) {
		p(s.Not())
	})
}

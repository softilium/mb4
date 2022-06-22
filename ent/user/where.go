// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserName applies equality check predicate on the "UserName" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// PasswordHash applies equality check predicate on the "PasswordHash" field. It's identical to PasswordHashEQ.
func PasswordHash(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordHash), v))
	})
}

// Admin applies equality check predicate on the "Admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// StartInvestAccountsFlow applies equality check predicate on the "StartInvestAccountsFlow" field. It's identical to StartInvestAccountsFlowEQ.
func StartInvestAccountsFlow(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// UserNameEQ applies the EQ predicate on the "UserName" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "UserName" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "UserName" field.
func UserNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "UserName" field.
func UserNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "UserName" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "UserName" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "UserName" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "UserName" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "UserName" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "UserName" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "UserName" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "UserName" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "UserName" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// PasswordHashEQ applies the EQ predicate on the "PasswordHash" field.
func PasswordHashEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashNEQ applies the NEQ predicate on the "PasswordHash" field.
func PasswordHashNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashIn applies the In predicate on the "PasswordHash" field.
func PasswordHashIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPasswordHash), v...))
	})
}

// PasswordHashNotIn applies the NotIn predicate on the "PasswordHash" field.
func PasswordHashNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPasswordHash), v...))
	})
}

// PasswordHashGT applies the GT predicate on the "PasswordHash" field.
func PasswordHashGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashGTE applies the GTE predicate on the "PasswordHash" field.
func PasswordHashGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashLT applies the LT predicate on the "PasswordHash" field.
func PasswordHashLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashLTE applies the LTE predicate on the "PasswordHash" field.
func PasswordHashLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashContains applies the Contains predicate on the "PasswordHash" field.
func PasswordHashContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashHasPrefix applies the HasPrefix predicate on the "PasswordHash" field.
func PasswordHashHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashHasSuffix applies the HasSuffix predicate on the "PasswordHash" field.
func PasswordHashHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashEqualFold applies the EqualFold predicate on the "PasswordHash" field.
func PasswordHashEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPasswordHash), v))
	})
}

// PasswordHashContainsFold applies the ContainsFold predicate on the "PasswordHash" field.
func PasswordHashContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPasswordHash), v))
	})
}

// AdminEQ applies the EQ predicate on the "Admin" field.
func AdminEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdmin), v))
	})
}

// AdminNEQ applies the NEQ predicate on the "Admin" field.
func AdminNEQ(v bool) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdmin), v))
	})
}

// StartInvestAccountsFlowEQ applies the EQ predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowNEQ applies the NEQ predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowIn applies the In predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartInvestAccountsFlow), v...))
	})
}

// StartInvestAccountsFlowNotIn applies the NotIn predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartInvestAccountsFlow), v...))
	})
}

// StartInvestAccountsFlowGT applies the GT predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowGTE applies the GTE predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowLT applies the LT predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowLTE applies the LTE predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartInvestAccountsFlow), v))
	})
}

// StartInvestAccountsFlowIsNil applies the IsNil predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartInvestAccountsFlow)))
	})
}

// StartInvestAccountsFlowNotNil applies the NotNil predicate on the "StartInvestAccountsFlow" field.
func StartInvestAccountsFlowNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartInvestAccountsFlow)))
	})
}

// HasInvestAccounts applies the HasEdge predicate on the "InvestAccounts" edge.
func HasInvestAccounts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InvestAccountsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InvestAccountsTable, InvestAccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvestAccountsWith applies the HasEdge predicate on the "InvestAccounts" edge with a given conditions (other predicates).
func HasInvestAccountsWith(preds ...predicate.InvestAccount) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InvestAccountsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InvestAccountsTable, InvestAccountsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}

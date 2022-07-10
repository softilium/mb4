// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/predicate"
	"github.com/softilium/mb4/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserName sets the "UserName" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetPasswordHash sets the "PasswordHash" field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetAdmin sets the "Admin" field.
func (uu *UserUpdate) SetAdmin(b bool) *UserUpdate {
	uu.mutation.SetAdmin(b)
	return uu
}

// SetNillableAdmin sets the "Admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetAdmin(*b)
	}
	return uu
}

// SetStartInvestAccountsFlow sets the "StartInvestAccountsFlow" field.
func (uu *UserUpdate) SetStartInvestAccountsFlow(t time.Time) *UserUpdate {
	uu.mutation.SetStartInvestAccountsFlow(t)
	return uu
}

// SetNillableStartInvestAccountsFlow sets the "StartInvestAccountsFlow" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStartInvestAccountsFlow(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetStartInvestAccountsFlow(*t)
	}
	return uu
}

// ClearStartInvestAccountsFlow clears the value of the "StartInvestAccountsFlow" field.
func (uu *UserUpdate) ClearStartInvestAccountsFlow() *UserUpdate {
	uu.mutation.ClearStartInvestAccountsFlow()
	return uu
}

// SetHowManyTickersOnHomepage sets the "HowManyTickersOnHomepage" field.
func (uu *UserUpdate) SetHowManyTickersOnHomepage(i int) *UserUpdate {
	uu.mutation.ResetHowManyTickersOnHomepage()
	uu.mutation.SetHowManyTickersOnHomepage(i)
	return uu
}

// SetNillableHowManyTickersOnHomepage sets the "HowManyTickersOnHomepage" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHowManyTickersOnHomepage(i *int) *UserUpdate {
	if i != nil {
		uu.SetHowManyTickersOnHomepage(*i)
	}
	return uu
}

// AddHowManyTickersOnHomepage adds i to the "HowManyTickersOnHomepage" field.
func (uu *UserUpdate) AddHowManyTickersOnHomepage(i int) *UserUpdate {
	uu.mutation.AddHowManyTickersOnHomepage(i)
	return uu
}

// AddInvestAccountIDs adds the "InvestAccounts" edge to the InvestAccount entity by IDs.
func (uu *UserUpdate) AddInvestAccountIDs(ids ...xid.ID) *UserUpdate {
	uu.mutation.AddInvestAccountIDs(ids...)
	return uu
}

// AddInvestAccounts adds the "InvestAccounts" edges to the InvestAccount entity.
func (uu *UserUpdate) AddInvestAccounts(i ...*InvestAccount) *UserUpdate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.AddInvestAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearInvestAccounts clears all "InvestAccounts" edges to the InvestAccount entity.
func (uu *UserUpdate) ClearInvestAccounts() *UserUpdate {
	uu.mutation.ClearInvestAccounts()
	return uu
}

// RemoveInvestAccountIDs removes the "InvestAccounts" edge to InvestAccount entities by IDs.
func (uu *UserUpdate) RemoveInvestAccountIDs(ids ...xid.ID) *UserUpdate {
	uu.mutation.RemoveInvestAccountIDs(ids...)
	return uu
}

// RemoveInvestAccounts removes "InvestAccounts" edges to InvestAccount entities.
func (uu *UserUpdate) RemoveInvestAccounts(i ...*InvestAccount) *UserUpdate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.RemoveInvestAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`ent: validator failed for field "User.UserName": %w`, err)}
		}
	}
	if v, ok := uu.mutation.HowManyTickersOnHomepage(); ok {
		if err := user.HowManyTickersOnHomepageValidator(v); err != nil {
			return &ValidationError{Name: "HowManyTickersOnHomepage", err: fmt.Errorf(`ent: validator failed for field "User.HowManyTickersOnHomepage": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
	}
	if value, ok := uu.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uu.mutation.StartInvestAccountsFlow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldStartInvestAccountsFlow,
		})
	}
	if uu.mutation.StartInvestAccountsFlowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldStartInvestAccountsFlow,
		})
	}
	if value, ok := uu.mutation.HowManyTickersOnHomepage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHowManyTickersOnHomepage,
		})
	}
	if value, ok := uu.mutation.AddedHowManyTickersOnHomepage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHowManyTickersOnHomepage,
		})
	}
	if uu.mutation.InvestAccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedInvestAccountsIDs(); len(nodes) > 0 && !uu.mutation.InvestAccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.InvestAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "UserName" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetPasswordHash sets the "PasswordHash" field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetAdmin sets the "Admin" field.
func (uuo *UserUpdateOne) SetAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetAdmin(b)
	return uuo
}

// SetNillableAdmin sets the "Admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetAdmin(*b)
	}
	return uuo
}

// SetStartInvestAccountsFlow sets the "StartInvestAccountsFlow" field.
func (uuo *UserUpdateOne) SetStartInvestAccountsFlow(t time.Time) *UserUpdateOne {
	uuo.mutation.SetStartInvestAccountsFlow(t)
	return uuo
}

// SetNillableStartInvestAccountsFlow sets the "StartInvestAccountsFlow" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStartInvestAccountsFlow(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetStartInvestAccountsFlow(*t)
	}
	return uuo
}

// ClearStartInvestAccountsFlow clears the value of the "StartInvestAccountsFlow" field.
func (uuo *UserUpdateOne) ClearStartInvestAccountsFlow() *UserUpdateOne {
	uuo.mutation.ClearStartInvestAccountsFlow()
	return uuo
}

// SetHowManyTickersOnHomepage sets the "HowManyTickersOnHomepage" field.
func (uuo *UserUpdateOne) SetHowManyTickersOnHomepage(i int) *UserUpdateOne {
	uuo.mutation.ResetHowManyTickersOnHomepage()
	uuo.mutation.SetHowManyTickersOnHomepage(i)
	return uuo
}

// SetNillableHowManyTickersOnHomepage sets the "HowManyTickersOnHomepage" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHowManyTickersOnHomepage(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetHowManyTickersOnHomepage(*i)
	}
	return uuo
}

// AddHowManyTickersOnHomepage adds i to the "HowManyTickersOnHomepage" field.
func (uuo *UserUpdateOne) AddHowManyTickersOnHomepage(i int) *UserUpdateOne {
	uuo.mutation.AddHowManyTickersOnHomepage(i)
	return uuo
}

// AddInvestAccountIDs adds the "InvestAccounts" edge to the InvestAccount entity by IDs.
func (uuo *UserUpdateOne) AddInvestAccountIDs(ids ...xid.ID) *UserUpdateOne {
	uuo.mutation.AddInvestAccountIDs(ids...)
	return uuo
}

// AddInvestAccounts adds the "InvestAccounts" edges to the InvestAccount entity.
func (uuo *UserUpdateOne) AddInvestAccounts(i ...*InvestAccount) *UserUpdateOne {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.AddInvestAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearInvestAccounts clears all "InvestAccounts" edges to the InvestAccount entity.
func (uuo *UserUpdateOne) ClearInvestAccounts() *UserUpdateOne {
	uuo.mutation.ClearInvestAccounts()
	return uuo
}

// RemoveInvestAccountIDs removes the "InvestAccounts" edge to InvestAccount entities by IDs.
func (uuo *UserUpdateOne) RemoveInvestAccountIDs(ids ...xid.ID) *UserUpdateOne {
	uuo.mutation.RemoveInvestAccountIDs(ids...)
	return uuo
}

// RemoveInvestAccounts removes "InvestAccounts" edges to InvestAccount entities.
func (uuo *UserUpdateOne) RemoveInvestAccounts(i ...*InvestAccount) *UserUpdateOne {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.RemoveInvestAccountIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`ent: validator failed for field "User.UserName": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.HowManyTickersOnHomepage(); ok {
		if err := user.HowManyTickersOnHomepageValidator(v); err != nil {
			return &ValidationError{Name: "HowManyTickersOnHomepage", err: fmt.Errorf(`ent: validator failed for field "User.HowManyTickersOnHomepage": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
	}
	if value, ok := uuo.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uuo.mutation.StartInvestAccountsFlow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldStartInvestAccountsFlow,
		})
	}
	if uuo.mutation.StartInvestAccountsFlowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldStartInvestAccountsFlow,
		})
	}
	if value, ok := uuo.mutation.HowManyTickersOnHomepage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHowManyTickersOnHomepage,
		})
	}
	if value, ok := uuo.mutation.AddedHowManyTickersOnHomepage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldHowManyTickersOnHomepage,
		})
	}
	if uuo.mutation.InvestAccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedInvestAccountsIDs(); len(nodes) > 0 && !uuo.mutation.InvestAccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.InvestAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.InvestAccountsTable,
			Columns: []string{user.InvestAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: investaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/ent/industry"
	"github.com/softilium/mb4/ent/predicate"
)

// IndustryUpdate is the builder for updating Industry entities.
type IndustryUpdate struct {
	config
	hooks    []Hook
	mutation *IndustryMutation
}

// Where appends a list predicates to the IndustryUpdate builder.
func (iu *IndustryUpdate) Where(ps ...predicate.Industry) *IndustryUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetDescr sets the "Descr" field.
func (iu *IndustryUpdate) SetDescr(s string) *IndustryUpdate {
	iu.mutation.SetDescr(s)
	return iu
}

// AddEmitentIDs adds the "Emitents" edge to the Emitent entity by IDs.
func (iu *IndustryUpdate) AddEmitentIDs(ids ...xid.ID) *IndustryUpdate {
	iu.mutation.AddEmitentIDs(ids...)
	return iu
}

// AddEmitents adds the "Emitents" edges to the Emitent entity.
func (iu *IndustryUpdate) AddEmitents(e ...*Emitent) *IndustryUpdate {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.AddEmitentIDs(ids...)
}

// Mutation returns the IndustryMutation object of the builder.
func (iu *IndustryUpdate) Mutation() *IndustryMutation {
	return iu.mutation
}

// ClearEmitents clears all "Emitents" edges to the Emitent entity.
func (iu *IndustryUpdate) ClearEmitents() *IndustryUpdate {
	iu.mutation.ClearEmitents()
	return iu
}

// RemoveEmitentIDs removes the "Emitents" edge to Emitent entities by IDs.
func (iu *IndustryUpdate) RemoveEmitentIDs(ids ...xid.ID) *IndustryUpdate {
	iu.mutation.RemoveEmitentIDs(ids...)
	return iu
}

// RemoveEmitents removes "Emitents" edges to Emitent entities.
func (iu *IndustryUpdate) RemoveEmitents(e ...*Emitent) *IndustryUpdate {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.RemoveEmitentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IndustryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndustryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IndustryUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IndustryUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IndustryUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IndustryUpdate) check() error {
	if v, ok := iu.mutation.Descr(); ok {
		if err := industry.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Industry.Descr": %w`, err)}
		}
	}
	return nil
}

func (iu *IndustryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   industry.Table,
			Columns: industry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: industry.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: industry.FieldDescr,
		})
	}
	if iu.mutation.EmitentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedEmitentsIDs(); len(nodes) > 0 && !iu.mutation.EmitentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.EmitentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{industry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IndustryUpdateOne is the builder for updating a single Industry entity.
type IndustryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IndustryMutation
}

// SetDescr sets the "Descr" field.
func (iuo *IndustryUpdateOne) SetDescr(s string) *IndustryUpdateOne {
	iuo.mutation.SetDescr(s)
	return iuo
}

// AddEmitentIDs adds the "Emitents" edge to the Emitent entity by IDs.
func (iuo *IndustryUpdateOne) AddEmitentIDs(ids ...xid.ID) *IndustryUpdateOne {
	iuo.mutation.AddEmitentIDs(ids...)
	return iuo
}

// AddEmitents adds the "Emitents" edges to the Emitent entity.
func (iuo *IndustryUpdateOne) AddEmitents(e ...*Emitent) *IndustryUpdateOne {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.AddEmitentIDs(ids...)
}

// Mutation returns the IndustryMutation object of the builder.
func (iuo *IndustryUpdateOne) Mutation() *IndustryMutation {
	return iuo.mutation
}

// ClearEmitents clears all "Emitents" edges to the Emitent entity.
func (iuo *IndustryUpdateOne) ClearEmitents() *IndustryUpdateOne {
	iuo.mutation.ClearEmitents()
	return iuo
}

// RemoveEmitentIDs removes the "Emitents" edge to Emitent entities by IDs.
func (iuo *IndustryUpdateOne) RemoveEmitentIDs(ids ...xid.ID) *IndustryUpdateOne {
	iuo.mutation.RemoveEmitentIDs(ids...)
	return iuo
}

// RemoveEmitents removes "Emitents" edges to Emitent entities.
func (iuo *IndustryUpdateOne) RemoveEmitents(e ...*Emitent) *IndustryUpdateOne {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.RemoveEmitentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IndustryUpdateOne) Select(field string, fields ...string) *IndustryUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Industry entity.
func (iuo *IndustryUpdateOne) Save(ctx context.Context) (*Industry, error) {
	var (
		err  error
		node *Industry
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndustryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IndustryUpdateOne) SaveX(ctx context.Context) *Industry {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IndustryUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IndustryUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IndustryUpdateOne) check() error {
	if v, ok := iuo.mutation.Descr(); ok {
		if err := industry.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Industry.Descr": %w`, err)}
		}
	}
	return nil
}

func (iuo *IndustryUpdateOne) sqlSave(ctx context.Context) (_node *Industry, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   industry.Table,
			Columns: industry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: industry.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Industry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, industry.FieldID)
		for _, f := range fields {
			if !industry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != industry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: industry.FieldDescr,
		})
	}
	if iuo.mutation.EmitentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedEmitentsIDs(); len(nodes) > 0 && !iuo.mutation.EmitentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.EmitentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   industry.EmitentsTable,
			Columns: []string{industry.EmitentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emitent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Industry{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{industry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

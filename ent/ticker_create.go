// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/divpayout"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/ent/quote"
	"github.com/softilium/mb4/ent/ticker"
)

// TickerCreate is the builder for creating a Ticker entity.
type TickerCreate struct {
	config
	mutation *TickerMutation
	hooks    []Hook
}

// SetDescr sets the "Descr" field.
func (tc *TickerCreate) SetDescr(s string) *TickerCreate {
	tc.mutation.SetDescr(s)
	return tc
}

// SetKind sets the "Kind" field.
func (tc *TickerCreate) SetKind(i int32) *TickerCreate {
	tc.mutation.SetKind(i)
	return tc
}

// SetNillableKind sets the "Kind" field if the given value is not nil.
func (tc *TickerCreate) SetNillableKind(i *int32) *TickerCreate {
	if i != nil {
		tc.SetKind(*i)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TickerCreate) SetID(s string) *TickerCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetEmitentID sets the "Emitent" edge to the Emitent entity by ID.
func (tc *TickerCreate) SetEmitentID(id xid.ID) *TickerCreate {
	tc.mutation.SetEmitentID(id)
	return tc
}

// SetNillableEmitentID sets the "Emitent" edge to the Emitent entity by ID if the given value is not nil.
func (tc *TickerCreate) SetNillableEmitentID(id *xid.ID) *TickerCreate {
	if id != nil {
		tc = tc.SetEmitentID(*id)
	}
	return tc
}

// SetEmitent sets the "Emitent" edge to the Emitent entity.
func (tc *TickerCreate) SetEmitent(e *Emitent) *TickerCreate {
	return tc.SetEmitentID(e.ID)
}

// AddQuoteIDs adds the "Quotes" edge to the Quote entity by IDs.
func (tc *TickerCreate) AddQuoteIDs(ids ...xid.ID) *TickerCreate {
	tc.mutation.AddQuoteIDs(ids...)
	return tc
}

// AddQuotes adds the "Quotes" edges to the Quote entity.
func (tc *TickerCreate) AddQuotes(q ...*Quote) *TickerCreate {
	ids := make([]xid.ID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tc.AddQuoteIDs(ids...)
}

// AddDivPayoutIDs adds the "DivPayouts" edge to the DivPayout entity by IDs.
func (tc *TickerCreate) AddDivPayoutIDs(ids ...int) *TickerCreate {
	tc.mutation.AddDivPayoutIDs(ids...)
	return tc
}

// AddDivPayouts adds the "DivPayouts" edges to the DivPayout entity.
func (tc *TickerCreate) AddDivPayouts(d ...*DivPayout) *TickerCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tc.AddDivPayoutIDs(ids...)
}

// Mutation returns the TickerMutation object of the builder.
func (tc *TickerCreate) Mutation() *TickerMutation {
	return tc.mutation
}

// Save creates the Ticker in the database.
func (tc *TickerCreate) Save(ctx context.Context) (*Ticker, error) {
	var (
		err  error
		node *Ticker
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TickerCreate) SaveX(ctx context.Context) *Ticker {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TickerCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TickerCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TickerCreate) defaults() {
	if _, ok := tc.mutation.Kind(); !ok {
		v := ticker.DefaultKind
		tc.mutation.SetKind(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TickerCreate) check() error {
	if _, ok := tc.mutation.Descr(); !ok {
		return &ValidationError{Name: "Descr", err: errors.New(`ent: missing required field "Ticker.Descr"`)}
	}
	if v, ok := tc.mutation.Descr(); ok {
		if err := ticker.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Ticker.Descr": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Kind(); !ok {
		return &ValidationError{Name: "Kind", err: errors.New(`ent: missing required field "Ticker.Kind"`)}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := ticker.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Ticker.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TickerCreate) sqlSave(ctx context.Context) (*Ticker, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Ticker.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TickerCreate) createSpec() (*Ticker, *sqlgraph.CreateSpec) {
	var (
		_node = &Ticker{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ticker.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: ticker.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Descr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticker.FieldDescr,
		})
		_node.Descr = value
	}
	if value, ok := tc.mutation.Kind(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: ticker.FieldKind,
		})
		_node.Kind = value
	}
	if nodes := tc.mutation.EmitentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticker.EmitentTable,
			Columns: []string{ticker.EmitentColumn},
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
		_node.emitent_tickers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.QuotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.QuotesTable,
			Columns: []string{ticker.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: quote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DivPayoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.DivPayoutsTable,
			Columns: []string{ticker.DivPayoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: divpayout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TickerCreateBulk is the builder for creating many Ticker entities in bulk.
type TickerCreateBulk struct {
	config
	builders []*TickerCreate
}

// Save creates the Ticker entities in the database.
func (tcb *TickerCreateBulk) Save(ctx context.Context) ([]*Ticker, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Ticker, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TickerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TickerCreateBulk) SaveX(ctx context.Context) []*Ticker {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TickerCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TickerCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

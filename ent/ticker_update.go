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
	"github.com/softilium/mb4/ent/divpayout"
	"github.com/softilium/mb4/ent/emission"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/ent/predicate"
	"github.com/softilium/mb4/ent/quote"
	"github.com/softilium/mb4/ent/ticker"
)

// TickerUpdate is the builder for updating Ticker entities.
type TickerUpdate struct {
	config
	hooks    []Hook
	mutation *TickerMutation
}

// Where appends a list predicates to the TickerUpdate builder.
func (tu *TickerUpdate) Where(ps ...predicate.Ticker) *TickerUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDescr sets the "Descr" field.
func (tu *TickerUpdate) SetDescr(s string) *TickerUpdate {
	tu.mutation.SetDescr(s)
	return tu
}

// SetKind sets the "Kind" field.
func (tu *TickerUpdate) SetKind(i int32) *TickerUpdate {
	tu.mutation.ResetKind()
	tu.mutation.SetKind(i)
	return tu
}

// SetNillableKind sets the "Kind" field if the given value is not nil.
func (tu *TickerUpdate) SetNillableKind(i *int32) *TickerUpdate {
	if i != nil {
		tu.SetKind(*i)
	}
	return tu
}

// AddKind adds i to the "Kind" field.
func (tu *TickerUpdate) AddKind(i int32) *TickerUpdate {
	tu.mutation.AddKind(i)
	return tu
}

// SetEmitentID sets the "Emitent" edge to the Emitent entity by ID.
func (tu *TickerUpdate) SetEmitentID(id xid.ID) *TickerUpdate {
	tu.mutation.SetEmitentID(id)
	return tu
}

// SetEmitent sets the "Emitent" edge to the Emitent entity.
func (tu *TickerUpdate) SetEmitent(e *Emitent) *TickerUpdate {
	return tu.SetEmitentID(e.ID)
}

// AddQuoteIDs adds the "Quotes" edge to the Quote entity by IDs.
func (tu *TickerUpdate) AddQuoteIDs(ids ...xid.ID) *TickerUpdate {
	tu.mutation.AddQuoteIDs(ids...)
	return tu
}

// AddQuotes adds the "Quotes" edges to the Quote entity.
func (tu *TickerUpdate) AddQuotes(q ...*Quote) *TickerUpdate {
	ids := make([]xid.ID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tu.AddQuoteIDs(ids...)
}

// AddDivPayoutIDs adds the "DivPayouts" edge to the DivPayout entity by IDs.
func (tu *TickerUpdate) AddDivPayoutIDs(ids ...int) *TickerUpdate {
	tu.mutation.AddDivPayoutIDs(ids...)
	return tu
}

// AddDivPayouts adds the "DivPayouts" edges to the DivPayout entity.
func (tu *TickerUpdate) AddDivPayouts(d ...*DivPayout) *TickerUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tu.AddDivPayoutIDs(ids...)
}

// AddEmissionIDs adds the "Emissions" edge to the Emission entity by IDs.
func (tu *TickerUpdate) AddEmissionIDs(ids ...xid.ID) *TickerUpdate {
	tu.mutation.AddEmissionIDs(ids...)
	return tu
}

// AddEmissions adds the "Emissions" edges to the Emission entity.
func (tu *TickerUpdate) AddEmissions(e ...*Emission) *TickerUpdate {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.AddEmissionIDs(ids...)
}

// Mutation returns the TickerMutation object of the builder.
func (tu *TickerUpdate) Mutation() *TickerMutation {
	return tu.mutation
}

// ClearEmitent clears the "Emitent" edge to the Emitent entity.
func (tu *TickerUpdate) ClearEmitent() *TickerUpdate {
	tu.mutation.ClearEmitent()
	return tu
}

// ClearQuotes clears all "Quotes" edges to the Quote entity.
func (tu *TickerUpdate) ClearQuotes() *TickerUpdate {
	tu.mutation.ClearQuotes()
	return tu
}

// RemoveQuoteIDs removes the "Quotes" edge to Quote entities by IDs.
func (tu *TickerUpdate) RemoveQuoteIDs(ids ...xid.ID) *TickerUpdate {
	tu.mutation.RemoveQuoteIDs(ids...)
	return tu
}

// RemoveQuotes removes "Quotes" edges to Quote entities.
func (tu *TickerUpdate) RemoveQuotes(q ...*Quote) *TickerUpdate {
	ids := make([]xid.ID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tu.RemoveQuoteIDs(ids...)
}

// ClearDivPayouts clears all "DivPayouts" edges to the DivPayout entity.
func (tu *TickerUpdate) ClearDivPayouts() *TickerUpdate {
	tu.mutation.ClearDivPayouts()
	return tu
}

// RemoveDivPayoutIDs removes the "DivPayouts" edge to DivPayout entities by IDs.
func (tu *TickerUpdate) RemoveDivPayoutIDs(ids ...int) *TickerUpdate {
	tu.mutation.RemoveDivPayoutIDs(ids...)
	return tu
}

// RemoveDivPayouts removes "DivPayouts" edges to DivPayout entities.
func (tu *TickerUpdate) RemoveDivPayouts(d ...*DivPayout) *TickerUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tu.RemoveDivPayoutIDs(ids...)
}

// ClearEmissions clears all "Emissions" edges to the Emission entity.
func (tu *TickerUpdate) ClearEmissions() *TickerUpdate {
	tu.mutation.ClearEmissions()
	return tu
}

// RemoveEmissionIDs removes the "Emissions" edge to Emission entities by IDs.
func (tu *TickerUpdate) RemoveEmissionIDs(ids ...xid.ID) *TickerUpdate {
	tu.mutation.RemoveEmissionIDs(ids...)
	return tu
}

// RemoveEmissions removes "Emissions" edges to Emission entities.
func (tu *TickerUpdate) RemoveEmissions(e ...*Emission) *TickerUpdate {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.RemoveEmissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TickerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TickerUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TickerUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TickerUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TickerUpdate) check() error {
	if v, ok := tu.mutation.Descr(); ok {
		if err := ticker.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Ticker.Descr": %w`, err)}
		}
	}
	if _, ok := tu.mutation.EmitentID(); tu.mutation.EmitentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ticker.Emitent"`)
	}
	return nil
}

func (tu *TickerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticker.Table,
			Columns: ticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: ticker.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticker.FieldDescr,
		})
	}
	if value, ok := tu.mutation.Kind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: ticker.FieldKind,
		})
	}
	if value, ok := tu.mutation.AddedKind(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: ticker.FieldKind,
		})
	}
	if tu.mutation.EmitentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EmitentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.QuotesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedQuotesIDs(); len(nodes) > 0 && !tu.mutation.QuotesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.QuotesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.DivPayoutsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedDivPayoutsIDs(); len(nodes) > 0 && !tu.mutation.DivPayoutsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.DivPayoutsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.EmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedEmissionsIDs(); len(nodes) > 0 && !tu.mutation.EmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TickerUpdateOne is the builder for updating a single Ticker entity.
type TickerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TickerMutation
}

// SetDescr sets the "Descr" field.
func (tuo *TickerUpdateOne) SetDescr(s string) *TickerUpdateOne {
	tuo.mutation.SetDescr(s)
	return tuo
}

// SetKind sets the "Kind" field.
func (tuo *TickerUpdateOne) SetKind(i int32) *TickerUpdateOne {
	tuo.mutation.ResetKind()
	tuo.mutation.SetKind(i)
	return tuo
}

// SetNillableKind sets the "Kind" field if the given value is not nil.
func (tuo *TickerUpdateOne) SetNillableKind(i *int32) *TickerUpdateOne {
	if i != nil {
		tuo.SetKind(*i)
	}
	return tuo
}

// AddKind adds i to the "Kind" field.
func (tuo *TickerUpdateOne) AddKind(i int32) *TickerUpdateOne {
	tuo.mutation.AddKind(i)
	return tuo
}

// SetEmitentID sets the "Emitent" edge to the Emitent entity by ID.
func (tuo *TickerUpdateOne) SetEmitentID(id xid.ID) *TickerUpdateOne {
	tuo.mutation.SetEmitentID(id)
	return tuo
}

// SetEmitent sets the "Emitent" edge to the Emitent entity.
func (tuo *TickerUpdateOne) SetEmitent(e *Emitent) *TickerUpdateOne {
	return tuo.SetEmitentID(e.ID)
}

// AddQuoteIDs adds the "Quotes" edge to the Quote entity by IDs.
func (tuo *TickerUpdateOne) AddQuoteIDs(ids ...xid.ID) *TickerUpdateOne {
	tuo.mutation.AddQuoteIDs(ids...)
	return tuo
}

// AddQuotes adds the "Quotes" edges to the Quote entity.
func (tuo *TickerUpdateOne) AddQuotes(q ...*Quote) *TickerUpdateOne {
	ids := make([]xid.ID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tuo.AddQuoteIDs(ids...)
}

// AddDivPayoutIDs adds the "DivPayouts" edge to the DivPayout entity by IDs.
func (tuo *TickerUpdateOne) AddDivPayoutIDs(ids ...int) *TickerUpdateOne {
	tuo.mutation.AddDivPayoutIDs(ids...)
	return tuo
}

// AddDivPayouts adds the "DivPayouts" edges to the DivPayout entity.
func (tuo *TickerUpdateOne) AddDivPayouts(d ...*DivPayout) *TickerUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tuo.AddDivPayoutIDs(ids...)
}

// AddEmissionIDs adds the "Emissions" edge to the Emission entity by IDs.
func (tuo *TickerUpdateOne) AddEmissionIDs(ids ...xid.ID) *TickerUpdateOne {
	tuo.mutation.AddEmissionIDs(ids...)
	return tuo
}

// AddEmissions adds the "Emissions" edges to the Emission entity.
func (tuo *TickerUpdateOne) AddEmissions(e ...*Emission) *TickerUpdateOne {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.AddEmissionIDs(ids...)
}

// Mutation returns the TickerMutation object of the builder.
func (tuo *TickerUpdateOne) Mutation() *TickerMutation {
	return tuo.mutation
}

// ClearEmitent clears the "Emitent" edge to the Emitent entity.
func (tuo *TickerUpdateOne) ClearEmitent() *TickerUpdateOne {
	tuo.mutation.ClearEmitent()
	return tuo
}

// ClearQuotes clears all "Quotes" edges to the Quote entity.
func (tuo *TickerUpdateOne) ClearQuotes() *TickerUpdateOne {
	tuo.mutation.ClearQuotes()
	return tuo
}

// RemoveQuoteIDs removes the "Quotes" edge to Quote entities by IDs.
func (tuo *TickerUpdateOne) RemoveQuoteIDs(ids ...xid.ID) *TickerUpdateOne {
	tuo.mutation.RemoveQuoteIDs(ids...)
	return tuo
}

// RemoveQuotes removes "Quotes" edges to Quote entities.
func (tuo *TickerUpdateOne) RemoveQuotes(q ...*Quote) *TickerUpdateOne {
	ids := make([]xid.ID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tuo.RemoveQuoteIDs(ids...)
}

// ClearDivPayouts clears all "DivPayouts" edges to the DivPayout entity.
func (tuo *TickerUpdateOne) ClearDivPayouts() *TickerUpdateOne {
	tuo.mutation.ClearDivPayouts()
	return tuo
}

// RemoveDivPayoutIDs removes the "DivPayouts" edge to DivPayout entities by IDs.
func (tuo *TickerUpdateOne) RemoveDivPayoutIDs(ids ...int) *TickerUpdateOne {
	tuo.mutation.RemoveDivPayoutIDs(ids...)
	return tuo
}

// RemoveDivPayouts removes "DivPayouts" edges to DivPayout entities.
func (tuo *TickerUpdateOne) RemoveDivPayouts(d ...*DivPayout) *TickerUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tuo.RemoveDivPayoutIDs(ids...)
}

// ClearEmissions clears all "Emissions" edges to the Emission entity.
func (tuo *TickerUpdateOne) ClearEmissions() *TickerUpdateOne {
	tuo.mutation.ClearEmissions()
	return tuo
}

// RemoveEmissionIDs removes the "Emissions" edge to Emission entities by IDs.
func (tuo *TickerUpdateOne) RemoveEmissionIDs(ids ...xid.ID) *TickerUpdateOne {
	tuo.mutation.RemoveEmissionIDs(ids...)
	return tuo
}

// RemoveEmissions removes "Emissions" edges to Emission entities.
func (tuo *TickerUpdateOne) RemoveEmissions(e ...*Emission) *TickerUpdateOne {
	ids := make([]xid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.RemoveEmissionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TickerUpdateOne) Select(field string, fields ...string) *TickerUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Ticker entity.
func (tuo *TickerUpdateOne) Save(ctx context.Context) (*Ticker, error) {
	var (
		err  error
		node *Ticker
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TickerUpdateOne) SaveX(ctx context.Context) *Ticker {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TickerUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TickerUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TickerUpdateOne) check() error {
	if v, ok := tuo.mutation.Descr(); ok {
		if err := ticker.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Ticker.Descr": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.EmitentID(); tuo.mutation.EmitentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ticker.Emitent"`)
	}
	return nil
}

func (tuo *TickerUpdateOne) sqlSave(ctx context.Context) (_node *Ticker, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticker.Table,
			Columns: ticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: ticker.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ticker.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticker.FieldID)
		for _, f := range fields {
			if !ticker.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ticker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticker.FieldDescr,
		})
	}
	if value, ok := tuo.mutation.Kind(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: ticker.FieldKind,
		})
	}
	if value, ok := tuo.mutation.AddedKind(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: ticker.FieldKind,
		})
	}
	if tuo.mutation.EmitentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EmitentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.QuotesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedQuotesIDs(); len(nodes) > 0 && !tuo.mutation.QuotesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.QuotesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.DivPayoutsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedDivPayoutsIDs(); len(nodes) > 0 && !tuo.mutation.DivPayoutsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.DivPayoutsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.EmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedEmissionsIDs(); len(nodes) > 0 && !tuo.mutation.EmissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EmissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticker.EmissionsTable,
			Columns: []string{ticker.EmissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: emission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ticker{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
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
	"github.com/softilium/mb4/ent/report"
	"github.com/softilium/mb4/ent/ticker"
)

// EmitentUpdate is the builder for updating Emitent entities.
type EmitentUpdate struct {
	config
	hooks    []Hook
	mutation *EmitentMutation
}

// Where appends a list predicates to the EmitentUpdate builder.
func (eu *EmitentUpdate) Where(ps ...predicate.Emitent) *EmitentUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetDescr sets the "Descr" field.
func (eu *EmitentUpdate) SetDescr(s string) *EmitentUpdate {
	eu.mutation.SetDescr(s)
	return eu
}

// SetIndustryID sets the "Industry" edge to the Industry entity by ID.
func (eu *EmitentUpdate) SetIndustryID(id string) *EmitentUpdate {
	eu.mutation.SetIndustryID(id)
	return eu
}

// SetIndustry sets the "Industry" edge to the Industry entity.
func (eu *EmitentUpdate) SetIndustry(i *Industry) *EmitentUpdate {
	return eu.SetIndustryID(i.ID)
}

// AddTickerIDs adds the "Tickers" edge to the Ticker entity by IDs.
func (eu *EmitentUpdate) AddTickerIDs(ids ...string) *EmitentUpdate {
	eu.mutation.AddTickerIDs(ids...)
	return eu
}

// AddTickers adds the "Tickers" edges to the Ticker entity.
func (eu *EmitentUpdate) AddTickers(t ...*Ticker) *EmitentUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.AddTickerIDs(ids...)
}

// AddReportIDs adds the "Reports" edge to the Report entity by IDs.
func (eu *EmitentUpdate) AddReportIDs(ids ...xid.ID) *EmitentUpdate {
	eu.mutation.AddReportIDs(ids...)
	return eu
}

// AddReports adds the "Reports" edges to the Report entity.
func (eu *EmitentUpdate) AddReports(r ...*Report) *EmitentUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eu.AddReportIDs(ids...)
}

// Mutation returns the EmitentMutation object of the builder.
func (eu *EmitentUpdate) Mutation() *EmitentMutation {
	return eu.mutation
}

// ClearIndustry clears the "Industry" edge to the Industry entity.
func (eu *EmitentUpdate) ClearIndustry() *EmitentUpdate {
	eu.mutation.ClearIndustry()
	return eu
}

// ClearTickers clears all "Tickers" edges to the Ticker entity.
func (eu *EmitentUpdate) ClearTickers() *EmitentUpdate {
	eu.mutation.ClearTickers()
	return eu
}

// RemoveTickerIDs removes the "Tickers" edge to Ticker entities by IDs.
func (eu *EmitentUpdate) RemoveTickerIDs(ids ...string) *EmitentUpdate {
	eu.mutation.RemoveTickerIDs(ids...)
	return eu
}

// RemoveTickers removes "Tickers" edges to Ticker entities.
func (eu *EmitentUpdate) RemoveTickers(t ...*Ticker) *EmitentUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.RemoveTickerIDs(ids...)
}

// ClearReports clears all "Reports" edges to the Report entity.
func (eu *EmitentUpdate) ClearReports() *EmitentUpdate {
	eu.mutation.ClearReports()
	return eu
}

// RemoveReportIDs removes the "Reports" edge to Report entities by IDs.
func (eu *EmitentUpdate) RemoveReportIDs(ids ...xid.ID) *EmitentUpdate {
	eu.mutation.RemoveReportIDs(ids...)
	return eu
}

// RemoveReports removes "Reports" edges to Report entities.
func (eu *EmitentUpdate) RemoveReports(r ...*Report) *EmitentUpdate {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eu.RemoveReportIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmitentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmitentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmitentUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmitentUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmitentUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmitentUpdate) check() error {
	if v, ok := eu.mutation.Descr(); ok {
		if err := emitent.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Emitent.Descr": %w`, err)}
		}
	}
	if _, ok := eu.mutation.IndustryID(); eu.mutation.IndustryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Emitent.Industry"`)
	}
	return nil
}

func (eu *EmitentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emitent.Table,
			Columns: emitent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: emitent.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emitent.FieldDescr,
		})
	}
	if eu.mutation.IndustryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emitent.IndustryTable,
			Columns: []string{emitent.IndustryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: industry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.IndustryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emitent.IndustryTable,
			Columns: []string{emitent.IndustryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: industry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.TickersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedTickersIDs(); len(nodes) > 0 && !eu.mutation.TickersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.TickersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ReportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedReportsIDs(); len(nodes) > 0 && !eu.mutation.ReportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ReportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emitent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EmitentUpdateOne is the builder for updating a single Emitent entity.
type EmitentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmitentMutation
}

// SetDescr sets the "Descr" field.
func (euo *EmitentUpdateOne) SetDescr(s string) *EmitentUpdateOne {
	euo.mutation.SetDescr(s)
	return euo
}

// SetIndustryID sets the "Industry" edge to the Industry entity by ID.
func (euo *EmitentUpdateOne) SetIndustryID(id string) *EmitentUpdateOne {
	euo.mutation.SetIndustryID(id)
	return euo
}

// SetIndustry sets the "Industry" edge to the Industry entity.
func (euo *EmitentUpdateOne) SetIndustry(i *Industry) *EmitentUpdateOne {
	return euo.SetIndustryID(i.ID)
}

// AddTickerIDs adds the "Tickers" edge to the Ticker entity by IDs.
func (euo *EmitentUpdateOne) AddTickerIDs(ids ...string) *EmitentUpdateOne {
	euo.mutation.AddTickerIDs(ids...)
	return euo
}

// AddTickers adds the "Tickers" edges to the Ticker entity.
func (euo *EmitentUpdateOne) AddTickers(t ...*Ticker) *EmitentUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.AddTickerIDs(ids...)
}

// AddReportIDs adds the "Reports" edge to the Report entity by IDs.
func (euo *EmitentUpdateOne) AddReportIDs(ids ...xid.ID) *EmitentUpdateOne {
	euo.mutation.AddReportIDs(ids...)
	return euo
}

// AddReports adds the "Reports" edges to the Report entity.
func (euo *EmitentUpdateOne) AddReports(r ...*Report) *EmitentUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return euo.AddReportIDs(ids...)
}

// Mutation returns the EmitentMutation object of the builder.
func (euo *EmitentUpdateOne) Mutation() *EmitentMutation {
	return euo.mutation
}

// ClearIndustry clears the "Industry" edge to the Industry entity.
func (euo *EmitentUpdateOne) ClearIndustry() *EmitentUpdateOne {
	euo.mutation.ClearIndustry()
	return euo
}

// ClearTickers clears all "Tickers" edges to the Ticker entity.
func (euo *EmitentUpdateOne) ClearTickers() *EmitentUpdateOne {
	euo.mutation.ClearTickers()
	return euo
}

// RemoveTickerIDs removes the "Tickers" edge to Ticker entities by IDs.
func (euo *EmitentUpdateOne) RemoveTickerIDs(ids ...string) *EmitentUpdateOne {
	euo.mutation.RemoveTickerIDs(ids...)
	return euo
}

// RemoveTickers removes "Tickers" edges to Ticker entities.
func (euo *EmitentUpdateOne) RemoveTickers(t ...*Ticker) *EmitentUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.RemoveTickerIDs(ids...)
}

// ClearReports clears all "Reports" edges to the Report entity.
func (euo *EmitentUpdateOne) ClearReports() *EmitentUpdateOne {
	euo.mutation.ClearReports()
	return euo
}

// RemoveReportIDs removes the "Reports" edge to Report entities by IDs.
func (euo *EmitentUpdateOne) RemoveReportIDs(ids ...xid.ID) *EmitentUpdateOne {
	euo.mutation.RemoveReportIDs(ids...)
	return euo
}

// RemoveReports removes "Reports" edges to Report entities.
func (euo *EmitentUpdateOne) RemoveReports(r ...*Report) *EmitentUpdateOne {
	ids := make([]xid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return euo.RemoveReportIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmitentUpdateOne) Select(field string, fields ...string) *EmitentUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Emitent entity.
func (euo *EmitentUpdateOne) Save(ctx context.Context) (*Emitent, error) {
	var (
		err  error
		node *Emitent
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmitentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmitentUpdateOne) SaveX(ctx context.Context) *Emitent {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmitentUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmitentUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmitentUpdateOne) check() error {
	if v, ok := euo.mutation.Descr(); ok {
		if err := emitent.DescrValidator(v); err != nil {
			return &ValidationError{Name: "Descr", err: fmt.Errorf(`ent: validator failed for field "Emitent.Descr": %w`, err)}
		}
	}
	if _, ok := euo.mutation.IndustryID(); euo.mutation.IndustryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Emitent.Industry"`)
	}
	return nil
}

func (euo *EmitentUpdateOne) sqlSave(ctx context.Context) (_node *Emitent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emitent.Table,
			Columns: emitent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: emitent.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Emitent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emitent.FieldID)
		for _, f := range fields {
			if !emitent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emitent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Descr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emitent.FieldDescr,
		})
	}
	if euo.mutation.IndustryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emitent.IndustryTable,
			Columns: []string{emitent.IndustryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: industry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.IndustryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emitent.IndustryTable,
			Columns: []string{emitent.IndustryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: industry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.TickersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedTickersIDs(); len(nodes) > 0 && !euo.mutation.TickersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.TickersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.TickersTable,
			Columns: []string{emitent.TickersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: ticker.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ReportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedReportsIDs(); len(nodes) > 0 && !euo.mutation.ReportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ReportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emitent.ReportsTable,
			Columns: []string{emitent.ReportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: report.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Emitent{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emitent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

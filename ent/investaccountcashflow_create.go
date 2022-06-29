// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/investaccountcashflow"
)

// InvestAccountCashflowCreate is the builder for creating a InvestAccountCashflow entity.
type InvestAccountCashflowCreate struct {
	config
	mutation *InvestAccountCashflowMutation
	hooks    []Hook
}

// SetRecDate sets the "RecDate" field.
func (iacc *InvestAccountCashflowCreate) SetRecDate(t time.Time) *InvestAccountCashflowCreate {
	iacc.mutation.SetRecDate(t)
	return iacc
}

// SetQty sets the "Qty" field.
func (iacc *InvestAccountCashflowCreate) SetQty(f float64) *InvestAccountCashflowCreate {
	iacc.mutation.SetQty(f)
	return iacc
}

// SetID sets the "id" field.
func (iacc *InvestAccountCashflowCreate) SetID(x xid.ID) *InvestAccountCashflowCreate {
	iacc.mutation.SetID(x)
	return iacc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (iacc *InvestAccountCashflowCreate) SetNillableID(x *xid.ID) *InvestAccountCashflowCreate {
	if x != nil {
		iacc.SetID(*x)
	}
	return iacc
}

// SetOwnerID sets the "Owner" edge to the InvestAccount entity by ID.
func (iacc *InvestAccountCashflowCreate) SetOwnerID(id xid.ID) *InvestAccountCashflowCreate {
	iacc.mutation.SetOwnerID(id)
	return iacc
}

// SetOwner sets the "Owner" edge to the InvestAccount entity.
func (iacc *InvestAccountCashflowCreate) SetOwner(i *InvestAccount) *InvestAccountCashflowCreate {
	return iacc.SetOwnerID(i.ID)
}

// Mutation returns the InvestAccountCashflowMutation object of the builder.
func (iacc *InvestAccountCashflowCreate) Mutation() *InvestAccountCashflowMutation {
	return iacc.mutation
}

// Save creates the InvestAccountCashflow in the database.
func (iacc *InvestAccountCashflowCreate) Save(ctx context.Context) (*InvestAccountCashflow, error) {
	var (
		err  error
		node *InvestAccountCashflow
	)
	iacc.defaults()
	if len(iacc.hooks) == 0 {
		if err = iacc.check(); err != nil {
			return nil, err
		}
		node, err = iacc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InvestAccountCashflowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iacc.check(); err != nil {
				return nil, err
			}
			iacc.mutation = mutation
			if node, err = iacc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(iacc.hooks) - 1; i >= 0; i-- {
			if iacc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iacc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iacc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iacc *InvestAccountCashflowCreate) SaveX(ctx context.Context) *InvestAccountCashflow {
	v, err := iacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iacc *InvestAccountCashflowCreate) Exec(ctx context.Context) error {
	_, err := iacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iacc *InvestAccountCashflowCreate) ExecX(ctx context.Context) {
	if err := iacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iacc *InvestAccountCashflowCreate) defaults() {
	if _, ok := iacc.mutation.ID(); !ok {
		v := investaccountcashflow.DefaultID()
		iacc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iacc *InvestAccountCashflowCreate) check() error {
	if _, ok := iacc.mutation.RecDate(); !ok {
		return &ValidationError{Name: "RecDate", err: errors.New(`ent: missing required field "InvestAccountCashflow.RecDate"`)}
	}
	if _, ok := iacc.mutation.Qty(); !ok {
		return &ValidationError{Name: "Qty", err: errors.New(`ent: missing required field "InvestAccountCashflow.Qty"`)}
	}
	if v, ok := iacc.mutation.ID(); ok {
		if err := investaccountcashflow.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "InvestAccountCashflow.id": %w`, err)}
		}
	}
	if _, ok := iacc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "Owner", err: errors.New(`ent: missing required edge "InvestAccountCashflow.Owner"`)}
	}
	return nil
}

func (iacc *InvestAccountCashflowCreate) sqlSave(ctx context.Context) (*InvestAccountCashflow, error) {
	_node, _spec := iacc.createSpec()
	if err := sqlgraph.CreateNode(ctx, iacc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (iacc *InvestAccountCashflowCreate) createSpec() (*InvestAccountCashflow, *sqlgraph.CreateSpec) {
	var (
		_node = &InvestAccountCashflow{config: iacc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: investaccountcashflow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: investaccountcashflow.FieldID,
			},
		}
	)
	if id, ok := iacc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := iacc.mutation.RecDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: investaccountcashflow.FieldRecDate,
		})
		_node.RecDate = value
	}
	if value, ok := iacc.mutation.Qty(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: investaccountcashflow.FieldQty,
		})
		_node.Qty = value
	}
	if nodes := iacc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   investaccountcashflow.OwnerTable,
			Columns: []string{investaccountcashflow.OwnerColumn},
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
		_node.invest_account_cashflows = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvestAccountCashflowCreateBulk is the builder for creating many InvestAccountCashflow entities in bulk.
type InvestAccountCashflowCreateBulk struct {
	config
	builders []*InvestAccountCashflowCreate
}

// Save creates the InvestAccountCashflow entities in the database.
func (iaccb *InvestAccountCashflowCreateBulk) Save(ctx context.Context) ([]*InvestAccountCashflow, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iaccb.builders))
	nodes := make([]*InvestAccountCashflow, len(iaccb.builders))
	mutators := make([]Mutator, len(iaccb.builders))
	for i := range iaccb.builders {
		func(i int, root context.Context) {
			builder := iaccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvestAccountCashflowMutation)
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
					_, err = mutators[i+1].Mutate(root, iaccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iaccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iaccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iaccb *InvestAccountCashflowCreateBulk) SaveX(ctx context.Context) []*InvestAccountCashflow {
	v, err := iaccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iaccb *InvestAccountCashflowCreateBulk) Exec(ctx context.Context) error {
	_, err := iaccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iaccb *InvestAccountCashflowCreateBulk) ExecX(ctx context.Context) {
	if err := iaccb.Exec(ctx); err != nil {
		panic(err)
	}
}

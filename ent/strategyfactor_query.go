// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/predicate"
	"github.com/softilium/mb4/ent/strategy"
	"github.com/softilium/mb4/ent/strategyfactor"
)

// StrategyFactorQuery is the builder for querying StrategyFactor entities.
type StrategyFactorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StrategyFactor
	// eager-loading edges.
	withStrategy *StrategyQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StrategyFactorQuery builder.
func (sfq *StrategyFactorQuery) Where(ps ...predicate.StrategyFactor) *StrategyFactorQuery {
	sfq.predicates = append(sfq.predicates, ps...)
	return sfq
}

// Limit adds a limit step to the query.
func (sfq *StrategyFactorQuery) Limit(limit int) *StrategyFactorQuery {
	sfq.limit = &limit
	return sfq
}

// Offset adds an offset step to the query.
func (sfq *StrategyFactorQuery) Offset(offset int) *StrategyFactorQuery {
	sfq.offset = &offset
	return sfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sfq *StrategyFactorQuery) Unique(unique bool) *StrategyFactorQuery {
	sfq.unique = &unique
	return sfq
}

// Order adds an order step to the query.
func (sfq *StrategyFactorQuery) Order(o ...OrderFunc) *StrategyFactorQuery {
	sfq.order = append(sfq.order, o...)
	return sfq
}

// QueryStrategy chains the current query on the "Strategy" edge.
func (sfq *StrategyFactorQuery) QueryStrategy() *StrategyQuery {
	query := &StrategyQuery{config: sfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(strategyfactor.Table, strategyfactor.FieldID, selector),
			sqlgraph.To(strategy.Table, strategy.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, strategyfactor.StrategyTable, strategyfactor.StrategyColumn),
		)
		fromU = sqlgraph.SetNeighbors(sfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StrategyFactor entity from the query.
// Returns a *NotFoundError when no StrategyFactor was found.
func (sfq *StrategyFactorQuery) First(ctx context.Context) (*StrategyFactor, error) {
	nodes, err := sfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{strategyfactor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sfq *StrategyFactorQuery) FirstX(ctx context.Context) *StrategyFactor {
	node, err := sfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StrategyFactor ID from the query.
// Returns a *NotFoundError when no StrategyFactor ID was found.
func (sfq *StrategyFactorQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = sfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{strategyfactor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sfq *StrategyFactorQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := sfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StrategyFactor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StrategyFactor entity is found.
// Returns a *NotFoundError when no StrategyFactor entities are found.
func (sfq *StrategyFactorQuery) Only(ctx context.Context) (*StrategyFactor, error) {
	nodes, err := sfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{strategyfactor.Label}
	default:
		return nil, &NotSingularError{strategyfactor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sfq *StrategyFactorQuery) OnlyX(ctx context.Context) *StrategyFactor {
	node, err := sfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StrategyFactor ID in the query.
// Returns a *NotSingularError when more than one StrategyFactor ID is found.
// Returns a *NotFoundError when no entities are found.
func (sfq *StrategyFactorQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = sfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = &NotSingularError{strategyfactor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sfq *StrategyFactorQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := sfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StrategyFactors.
func (sfq *StrategyFactorQuery) All(ctx context.Context) ([]*StrategyFactor, error) {
	if err := sfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sfq *StrategyFactorQuery) AllX(ctx context.Context) []*StrategyFactor {
	nodes, err := sfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StrategyFactor IDs.
func (sfq *StrategyFactorQuery) IDs(ctx context.Context) ([]xid.ID, error) {
	var ids []xid.ID
	if err := sfq.Select(strategyfactor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sfq *StrategyFactorQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := sfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sfq *StrategyFactorQuery) Count(ctx context.Context) (int, error) {
	if err := sfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sfq *StrategyFactorQuery) CountX(ctx context.Context) int {
	count, err := sfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sfq *StrategyFactorQuery) Exist(ctx context.Context) (bool, error) {
	if err := sfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sfq *StrategyFactorQuery) ExistX(ctx context.Context) bool {
	exist, err := sfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StrategyFactorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sfq *StrategyFactorQuery) Clone() *StrategyFactorQuery {
	if sfq == nil {
		return nil
	}
	return &StrategyFactorQuery{
		config:       sfq.config,
		limit:        sfq.limit,
		offset:       sfq.offset,
		order:        append([]OrderFunc{}, sfq.order...),
		predicates:   append([]predicate.StrategyFactor{}, sfq.predicates...),
		withStrategy: sfq.withStrategy.Clone(),
		// clone intermediate query.
		sql:    sfq.sql.Clone(),
		path:   sfq.path,
		unique: sfq.unique,
	}
}

// WithStrategy tells the query-builder to eager-load the nodes that are connected to
// the "Strategy" edge. The optional arguments are used to configure the query builder of the edge.
func (sfq *StrategyFactorQuery) WithStrategy(opts ...func(*StrategyQuery)) *StrategyFactorQuery {
	query := &StrategyQuery{config: sfq.config}
	for _, opt := range opts {
		opt(query)
	}
	sfq.withStrategy = query
	return sfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LineNum int `json:"LineNum,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StrategyFactor.Query().
//		GroupBy(strategyfactor.FieldLineNum).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sfq *StrategyFactorQuery) GroupBy(field string, fields ...string) *StrategyFactorGroupBy {
	group := &StrategyFactorGroupBy{config: sfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LineNum int `json:"LineNum,omitempty"`
//	}
//
//	client.StrategyFactor.Query().
//		Select(strategyfactor.FieldLineNum).
//		Scan(ctx, &v)
//
func (sfq *StrategyFactorQuery) Select(fields ...string) *StrategyFactorSelect {
	sfq.fields = append(sfq.fields, fields...)
	return &StrategyFactorSelect{StrategyFactorQuery: sfq}
}

func (sfq *StrategyFactorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sfq.fields {
		if !strategyfactor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sfq.path != nil {
		prev, err := sfq.path(ctx)
		if err != nil {
			return err
		}
		sfq.sql = prev
	}
	return nil
}

func (sfq *StrategyFactorQuery) sqlAll(ctx context.Context) ([]*StrategyFactor, error) {
	var (
		nodes       = []*StrategyFactor{}
		withFKs     = sfq.withFKs
		_spec       = sfq.querySpec()
		loadedTypes = [1]bool{
			sfq.withStrategy != nil,
		}
	)
	if sfq.withStrategy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, strategyfactor.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StrategyFactor{config: sfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, sfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sfq.withStrategy; query != nil {
		ids := make([]xid.ID, 0, len(nodes))
		nodeids := make(map[xid.ID][]*StrategyFactor)
		for i := range nodes {
			if nodes[i].strategy_factors == nil {
				continue
			}
			fk := *nodes[i].strategy_factors
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(strategy.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "strategy_factors" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Strategy = n
			}
		}
	}

	return nodes, nil
}

func (sfq *StrategyFactorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sfq.querySpec()
	_spec.Node.Columns = sfq.fields
	if len(sfq.fields) > 0 {
		_spec.Unique = sfq.unique != nil && *sfq.unique
	}
	return sqlgraph.CountNodes(ctx, sfq.driver, _spec)
}

func (sfq *StrategyFactorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sfq *StrategyFactorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   strategyfactor.Table,
			Columns: strategyfactor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: strategyfactor.FieldID,
			},
		},
		From:   sfq.sql,
		Unique: true,
	}
	if unique := sfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, strategyfactor.FieldID)
		for i := range fields {
			if fields[i] != strategyfactor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sfq *StrategyFactorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sfq.driver.Dialect())
	t1 := builder.Table(strategyfactor.Table)
	columns := sfq.fields
	if len(columns) == 0 {
		columns = strategyfactor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sfq.sql != nil {
		selector = sfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sfq.unique != nil && *sfq.unique {
		selector.Distinct()
	}
	for _, p := range sfq.predicates {
		p(selector)
	}
	for _, p := range sfq.order {
		p(selector)
	}
	if offset := sfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StrategyFactorGroupBy is the group-by builder for StrategyFactor entities.
type StrategyFactorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sfgb *StrategyFactorGroupBy) Aggregate(fns ...AggregateFunc) *StrategyFactorGroupBy {
	sfgb.fns = append(sfgb.fns, fns...)
	return sfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sfgb *StrategyFactorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sfgb.path(ctx)
	if err != nil {
		return err
	}
	sfgb.sql = query
	return sfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sfgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) StringsX(ctx context.Context) []string {
	v, err := sfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) StringX(ctx context.Context) string {
	v, err := sfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sfgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) IntsX(ctx context.Context) []int {
	v, err := sfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) IntX(ctx context.Context) int {
	v, err := sfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sfgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sfgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sfgb *StrategyFactorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sfgb *StrategyFactorGroupBy) BoolX(ctx context.Context) bool {
	v, err := sfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sfgb *StrategyFactorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sfgb.fields {
		if !strategyfactor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sfgb *StrategyFactorGroupBy) sqlQuery() *sql.Selector {
	selector := sfgb.sql.Select()
	aggregation := make([]string, 0, len(sfgb.fns))
	for _, fn := range sfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sfgb.fields)+len(sfgb.fns))
		for _, f := range sfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sfgb.fields...)...)
}

// StrategyFactorSelect is the builder for selecting fields of StrategyFactor entities.
type StrategyFactorSelect struct {
	*StrategyFactorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sfs *StrategyFactorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sfs.prepareQuery(ctx); err != nil {
		return err
	}
	sfs.sql = sfs.StrategyFactorQuery.sqlQuery(ctx)
	return sfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sfs *StrategyFactorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sfs.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sfs *StrategyFactorSelect) StringsX(ctx context.Context) []string {
	v, err := sfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sfs *StrategyFactorSelect) StringX(ctx context.Context) string {
	v, err := sfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sfs.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sfs *StrategyFactorSelect) IntsX(ctx context.Context) []int {
	v, err := sfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sfs *StrategyFactorSelect) IntX(ctx context.Context) int {
	v, err := sfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sfs.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sfs *StrategyFactorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sfs *StrategyFactorSelect) Float64X(ctx context.Context) float64 {
	v, err := sfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sfs.fields) > 1 {
		return nil, errors.New("ent: StrategyFactorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sfs *StrategyFactorSelect) BoolsX(ctx context.Context) []bool {
	v, err := sfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sfs *StrategyFactorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfactor.Label}
	default:
		err = fmt.Errorf("ent: StrategyFactorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sfs *StrategyFactorSelect) BoolX(ctx context.Context) bool {
	v, err := sfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sfs *StrategyFactorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sfs.sql.Query()
	if err := sfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

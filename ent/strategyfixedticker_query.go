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
	"github.com/softilium/mb4/ent/strategyfixedticker"
)

// StrategyFixedTickerQuery is the builder for querying StrategyFixedTicker entities.
type StrategyFixedTickerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StrategyFixedTicker
	// eager-loading edges.
	withStrategy *StrategyQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StrategyFixedTickerQuery builder.
func (sftq *StrategyFixedTickerQuery) Where(ps ...predicate.StrategyFixedTicker) *StrategyFixedTickerQuery {
	sftq.predicates = append(sftq.predicates, ps...)
	return sftq
}

// Limit adds a limit step to the query.
func (sftq *StrategyFixedTickerQuery) Limit(limit int) *StrategyFixedTickerQuery {
	sftq.limit = &limit
	return sftq
}

// Offset adds an offset step to the query.
func (sftq *StrategyFixedTickerQuery) Offset(offset int) *StrategyFixedTickerQuery {
	sftq.offset = &offset
	return sftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sftq *StrategyFixedTickerQuery) Unique(unique bool) *StrategyFixedTickerQuery {
	sftq.unique = &unique
	return sftq
}

// Order adds an order step to the query.
func (sftq *StrategyFixedTickerQuery) Order(o ...OrderFunc) *StrategyFixedTickerQuery {
	sftq.order = append(sftq.order, o...)
	return sftq
}

// QueryStrategy chains the current query on the "Strategy" edge.
func (sftq *StrategyFixedTickerQuery) QueryStrategy() *StrategyQuery {
	query := &StrategyQuery{config: sftq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sftq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(strategyfixedticker.Table, strategyfixedticker.FieldID, selector),
			sqlgraph.To(strategy.Table, strategy.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, strategyfixedticker.StrategyTable, strategyfixedticker.StrategyColumn),
		)
		fromU = sqlgraph.SetNeighbors(sftq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StrategyFixedTicker entity from the query.
// Returns a *NotFoundError when no StrategyFixedTicker was found.
func (sftq *StrategyFixedTickerQuery) First(ctx context.Context) (*StrategyFixedTicker, error) {
	nodes, err := sftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{strategyfixedticker.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) FirstX(ctx context.Context) *StrategyFixedTicker {
	node, err := sftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StrategyFixedTicker ID from the query.
// Returns a *NotFoundError when no StrategyFixedTicker ID was found.
func (sftq *StrategyFixedTickerQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = sftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{strategyfixedticker.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := sftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StrategyFixedTicker entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StrategyFixedTicker entity is found.
// Returns a *NotFoundError when no StrategyFixedTicker entities are found.
func (sftq *StrategyFixedTickerQuery) Only(ctx context.Context) (*StrategyFixedTicker, error) {
	nodes, err := sftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{strategyfixedticker.Label}
	default:
		return nil, &NotSingularError{strategyfixedticker.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) OnlyX(ctx context.Context) *StrategyFixedTicker {
	node, err := sftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StrategyFixedTicker ID in the query.
// Returns a *NotSingularError when more than one StrategyFixedTicker ID is found.
// Returns a *NotFoundError when no entities are found.
func (sftq *StrategyFixedTickerQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = sftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = &NotSingularError{strategyfixedticker.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := sftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StrategyFixedTickers.
func (sftq *StrategyFixedTickerQuery) All(ctx context.Context) ([]*StrategyFixedTicker, error) {
	if err := sftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sftq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) AllX(ctx context.Context) []*StrategyFixedTicker {
	nodes, err := sftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StrategyFixedTicker IDs.
func (sftq *StrategyFixedTickerQuery) IDs(ctx context.Context) ([]xid.ID, error) {
	var ids []xid.ID
	if err := sftq.Select(strategyfixedticker.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := sftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sftq *StrategyFixedTickerQuery) Count(ctx context.Context) (int, error) {
	if err := sftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sftq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) CountX(ctx context.Context) int {
	count, err := sftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sftq *StrategyFixedTickerQuery) Exist(ctx context.Context) (bool, error) {
	if err := sftq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sftq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sftq *StrategyFixedTickerQuery) ExistX(ctx context.Context) bool {
	exist, err := sftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StrategyFixedTickerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sftq *StrategyFixedTickerQuery) Clone() *StrategyFixedTickerQuery {
	if sftq == nil {
		return nil
	}
	return &StrategyFixedTickerQuery{
		config:       sftq.config,
		limit:        sftq.limit,
		offset:       sftq.offset,
		order:        append([]OrderFunc{}, sftq.order...),
		predicates:   append([]predicate.StrategyFixedTicker{}, sftq.predicates...),
		withStrategy: sftq.withStrategy.Clone(),
		// clone intermediate query.
		sql:    sftq.sql.Clone(),
		path:   sftq.path,
		unique: sftq.unique,
	}
}

// WithStrategy tells the query-builder to eager-load the nodes that are connected to
// the "Strategy" edge. The optional arguments are used to configure the query builder of the edge.
func (sftq *StrategyFixedTickerQuery) WithStrategy(opts ...func(*StrategyQuery)) *StrategyFixedTickerQuery {
	query := &StrategyQuery{config: sftq.config}
	for _, opt := range opts {
		opt(query)
	}
	sftq.withStrategy = query
	return sftq
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
//	client.StrategyFixedTicker.Query().
//		GroupBy(strategyfixedticker.FieldLineNum).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sftq *StrategyFixedTickerQuery) GroupBy(field string, fields ...string) *StrategyFixedTickerGroupBy {
	group := &StrategyFixedTickerGroupBy{config: sftq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sftq.sqlQuery(ctx), nil
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
//	client.StrategyFixedTicker.Query().
//		Select(strategyfixedticker.FieldLineNum).
//		Scan(ctx, &v)
//
func (sftq *StrategyFixedTickerQuery) Select(fields ...string) *StrategyFixedTickerSelect {
	sftq.fields = append(sftq.fields, fields...)
	return &StrategyFixedTickerSelect{StrategyFixedTickerQuery: sftq}
}

func (sftq *StrategyFixedTickerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sftq.fields {
		if !strategyfixedticker.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sftq.path != nil {
		prev, err := sftq.path(ctx)
		if err != nil {
			return err
		}
		sftq.sql = prev
	}
	return nil
}

func (sftq *StrategyFixedTickerQuery) sqlAll(ctx context.Context) ([]*StrategyFixedTicker, error) {
	var (
		nodes       = []*StrategyFixedTicker{}
		withFKs     = sftq.withFKs
		_spec       = sftq.querySpec()
		loadedTypes = [1]bool{
			sftq.withStrategy != nil,
		}
	)
	if sftq.withStrategy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, strategyfixedticker.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StrategyFixedTicker{config: sftq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sftq.withStrategy; query != nil {
		ids := make([]xid.ID, 0, len(nodes))
		nodeids := make(map[xid.ID][]*StrategyFixedTicker)
		for i := range nodes {
			if nodes[i].strategy_fixed_tickers == nil {
				continue
			}
			fk := *nodes[i].strategy_fixed_tickers
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
				return nil, fmt.Errorf(`unexpected foreign-key "strategy_fixed_tickers" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Strategy = n
			}
		}
	}

	return nodes, nil
}

func (sftq *StrategyFixedTickerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sftq.querySpec()
	_spec.Node.Columns = sftq.fields
	if len(sftq.fields) > 0 {
		_spec.Unique = sftq.unique != nil && *sftq.unique
	}
	return sqlgraph.CountNodes(ctx, sftq.driver, _spec)
}

func (sftq *StrategyFixedTickerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sftq *StrategyFixedTickerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   strategyfixedticker.Table,
			Columns: strategyfixedticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: strategyfixedticker.FieldID,
			},
		},
		From:   sftq.sql,
		Unique: true,
	}
	if unique := sftq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sftq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, strategyfixedticker.FieldID)
		for i := range fields {
			if fields[i] != strategyfixedticker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sftq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sftq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sftq *StrategyFixedTickerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sftq.driver.Dialect())
	t1 := builder.Table(strategyfixedticker.Table)
	columns := sftq.fields
	if len(columns) == 0 {
		columns = strategyfixedticker.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sftq.sql != nil {
		selector = sftq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sftq.unique != nil && *sftq.unique {
		selector.Distinct()
	}
	for _, p := range sftq.predicates {
		p(selector)
	}
	for _, p := range sftq.order {
		p(selector)
	}
	if offset := sftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StrategyFixedTickerGroupBy is the group-by builder for StrategyFixedTicker entities.
type StrategyFixedTickerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sftgb *StrategyFixedTickerGroupBy) Aggregate(fns ...AggregateFunc) *StrategyFixedTickerGroupBy {
	sftgb.fns = append(sftgb.fns, fns...)
	return sftgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sftgb *StrategyFixedTickerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sftgb.path(ctx)
	if err != nil {
		return err
	}
	sftgb.sql = query
	return sftgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sftgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) StringsX(ctx context.Context) []string {
	v, err := sftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sftgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) StringX(ctx context.Context) string {
	v, err := sftgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sftgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) IntsX(ctx context.Context) []int {
	v, err := sftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sftgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) IntX(ctx context.Context) int {
	v, err := sftgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sftgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sftgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sftgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sftgb.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sftgb *StrategyFixedTickerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sftgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sftgb *StrategyFixedTickerGroupBy) BoolX(ctx context.Context) bool {
	v, err := sftgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sftgb *StrategyFixedTickerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sftgb.fields {
		if !strategyfixedticker.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sftgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sftgb *StrategyFixedTickerGroupBy) sqlQuery() *sql.Selector {
	selector := sftgb.sql.Select()
	aggregation := make([]string, 0, len(sftgb.fns))
	for _, fn := range sftgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sftgb.fields)+len(sftgb.fns))
		for _, f := range sftgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sftgb.fields...)...)
}

// StrategyFixedTickerSelect is the builder for selecting fields of StrategyFixedTicker entities.
type StrategyFixedTickerSelect struct {
	*StrategyFixedTickerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sfts *StrategyFixedTickerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sfts.prepareQuery(ctx); err != nil {
		return err
	}
	sfts.sql = sfts.StrategyFixedTickerQuery.sqlQuery(ctx)
	return sfts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sfts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sfts.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sfts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) StringsX(ctx context.Context) []string {
	v, err := sfts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sfts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) StringX(ctx context.Context) string {
	v, err := sfts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sfts.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sfts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) IntsX(ctx context.Context) []int {
	v, err := sfts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sfts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) IntX(ctx context.Context) int {
	v, err := sfts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sfts.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sfts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sfts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sfts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) Float64X(ctx context.Context) float64 {
	v, err := sfts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sfts.fields) > 1 {
		return nil, errors.New("ent: StrategyFixedTickerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sfts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) BoolsX(ctx context.Context) []bool {
	v, err := sfts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sfts *StrategyFixedTickerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sfts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{strategyfixedticker.Label}
	default:
		err = fmt.Errorf("ent: StrategyFixedTickerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sfts *StrategyFixedTickerSelect) BoolX(ctx context.Context) bool {
	v, err := sfts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sfts *StrategyFixedTickerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sfts.sql.Query()
	if err := sfts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

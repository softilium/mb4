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
	"github.com/softilium/mb4/ent/divpayout"
	"github.com/softilium/mb4/ent/predicate"
	"github.com/softilium/mb4/ent/ticker"
)

// DivPayoutQuery is the builder for querying DivPayout entities.
type DivPayoutQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DivPayout
	// eager-loading edges.
	withTickers *TickerQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DivPayoutQuery builder.
func (dpq *DivPayoutQuery) Where(ps ...predicate.DivPayout) *DivPayoutQuery {
	dpq.predicates = append(dpq.predicates, ps...)
	return dpq
}

// Limit adds a limit step to the query.
func (dpq *DivPayoutQuery) Limit(limit int) *DivPayoutQuery {
	dpq.limit = &limit
	return dpq
}

// Offset adds an offset step to the query.
func (dpq *DivPayoutQuery) Offset(offset int) *DivPayoutQuery {
	dpq.offset = &offset
	return dpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dpq *DivPayoutQuery) Unique(unique bool) *DivPayoutQuery {
	dpq.unique = &unique
	return dpq
}

// Order adds an order step to the query.
func (dpq *DivPayoutQuery) Order(o ...OrderFunc) *DivPayoutQuery {
	dpq.order = append(dpq.order, o...)
	return dpq
}

// QueryTickers chains the current query on the "Tickers" edge.
func (dpq *DivPayoutQuery) QueryTickers() *TickerQuery {
	query := &TickerQuery{config: dpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(divpayout.Table, divpayout.FieldID, selector),
			sqlgraph.To(ticker.Table, ticker.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, divpayout.TickersTable, divpayout.TickersColumn),
		)
		fromU = sqlgraph.SetNeighbors(dpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DivPayout entity from the query.
// Returns a *NotFoundError when no DivPayout was found.
func (dpq *DivPayoutQuery) First(ctx context.Context) (*DivPayout, error) {
	nodes, err := dpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{divpayout.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dpq *DivPayoutQuery) FirstX(ctx context.Context) *DivPayout {
	node, err := dpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DivPayout ID from the query.
// Returns a *NotFoundError when no DivPayout ID was found.
func (dpq *DivPayoutQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{divpayout.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dpq *DivPayoutQuery) FirstIDX(ctx context.Context) int {
	id, err := dpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DivPayout entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DivPayout entity is found.
// Returns a *NotFoundError when no DivPayout entities are found.
func (dpq *DivPayoutQuery) Only(ctx context.Context) (*DivPayout, error) {
	nodes, err := dpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{divpayout.Label}
	default:
		return nil, &NotSingularError{divpayout.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dpq *DivPayoutQuery) OnlyX(ctx context.Context) *DivPayout {
	node, err := dpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DivPayout ID in the query.
// Returns a *NotSingularError when more than one DivPayout ID is found.
// Returns a *NotFoundError when no entities are found.
func (dpq *DivPayoutQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = &NotSingularError{divpayout.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dpq *DivPayoutQuery) OnlyIDX(ctx context.Context) int {
	id, err := dpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DivPayouts.
func (dpq *DivPayoutQuery) All(ctx context.Context) ([]*DivPayout, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dpq *DivPayoutQuery) AllX(ctx context.Context) []*DivPayout {
	nodes, err := dpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DivPayout IDs.
func (dpq *DivPayoutQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dpq.Select(divpayout.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dpq *DivPayoutQuery) IDsX(ctx context.Context) []int {
	ids, err := dpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dpq *DivPayoutQuery) Count(ctx context.Context) (int, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dpq *DivPayoutQuery) CountX(ctx context.Context) int {
	count, err := dpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dpq *DivPayoutQuery) Exist(ctx context.Context) (bool, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dpq *DivPayoutQuery) ExistX(ctx context.Context) bool {
	exist, err := dpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DivPayoutQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dpq *DivPayoutQuery) Clone() *DivPayoutQuery {
	if dpq == nil {
		return nil
	}
	return &DivPayoutQuery{
		config:      dpq.config,
		limit:       dpq.limit,
		offset:      dpq.offset,
		order:       append([]OrderFunc{}, dpq.order...),
		predicates:  append([]predicate.DivPayout{}, dpq.predicates...),
		withTickers: dpq.withTickers.Clone(),
		// clone intermediate query.
		sql:    dpq.sql.Clone(),
		path:   dpq.path,
		unique: dpq.unique,
	}
}

// WithTickers tells the query-builder to eager-load the nodes that are connected to
// the "Tickers" edge. The optional arguments are used to configure the query builder of the edge.
func (dpq *DivPayoutQuery) WithTickers(opts ...func(*TickerQuery)) *DivPayoutQuery {
	query := &TickerQuery{config: dpq.config}
	for _, opt := range opts {
		opt(query)
	}
	dpq.withTickers = query
	return dpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ForYear int `json:"ForYear,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DivPayout.Query().
//		GroupBy(divpayout.FieldForYear).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dpq *DivPayoutQuery) GroupBy(field string, fields ...string) *DivPayoutGroupBy {
	group := &DivPayoutGroupBy{config: dpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ForYear int `json:"ForYear,omitempty"`
//	}
//
//	client.DivPayout.Query().
//		Select(divpayout.FieldForYear).
//		Scan(ctx, &v)
//
func (dpq *DivPayoutQuery) Select(fields ...string) *DivPayoutSelect {
	dpq.fields = append(dpq.fields, fields...)
	return &DivPayoutSelect{DivPayoutQuery: dpq}
}

func (dpq *DivPayoutQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dpq.fields {
		if !divpayout.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dpq.path != nil {
		prev, err := dpq.path(ctx)
		if err != nil {
			return err
		}
		dpq.sql = prev
	}
	return nil
}

func (dpq *DivPayoutQuery) sqlAll(ctx context.Context) ([]*DivPayout, error) {
	var (
		nodes       = []*DivPayout{}
		withFKs     = dpq.withFKs
		_spec       = dpq.querySpec()
		loadedTypes = [1]bool{
			dpq.withTickers != nil,
		}
	)
	if dpq.withTickers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, divpayout.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DivPayout{config: dpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dpq.withTickers; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*DivPayout)
		for i := range nodes {
			if nodes[i].ticker_div_payouts == nil {
				continue
			}
			fk := *nodes[i].ticker_div_payouts
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(ticker.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ticker_div_payouts" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tickers = n
			}
		}
	}

	return nodes, nil
}

func (dpq *DivPayoutQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dpq.querySpec()
	_spec.Node.Columns = dpq.fields
	if len(dpq.fields) > 0 {
		_spec.Unique = dpq.unique != nil && *dpq.unique
	}
	return sqlgraph.CountNodes(ctx, dpq.driver, _spec)
}

func (dpq *DivPayoutQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dpq *DivPayoutQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   divpayout.Table,
			Columns: divpayout.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: divpayout.FieldID,
			},
		},
		From:   dpq.sql,
		Unique: true,
	}
	if unique := dpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, divpayout.FieldID)
		for i := range fields {
			if fields[i] != divpayout.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dpq *DivPayoutQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dpq.driver.Dialect())
	t1 := builder.Table(divpayout.Table)
	columns := dpq.fields
	if len(columns) == 0 {
		columns = divpayout.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dpq.sql != nil {
		selector = dpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dpq.unique != nil && *dpq.unique {
		selector.Distinct()
	}
	for _, p := range dpq.predicates {
		p(selector)
	}
	for _, p := range dpq.order {
		p(selector)
	}
	if offset := dpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DivPayoutGroupBy is the group-by builder for DivPayout entities.
type DivPayoutGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dpgb *DivPayoutGroupBy) Aggregate(fns ...AggregateFunc) *DivPayoutGroupBy {
	dpgb.fns = append(dpgb.fns, fns...)
	return dpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dpgb *DivPayoutGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dpgb.path(ctx)
	if err != nil {
		return err
	}
	dpgb.sql = query
	return dpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DivPayoutGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) StringsX(ctx context.Context) []string {
	v, err := dpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) StringX(ctx context.Context) string {
	v, err := dpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DivPayoutGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) IntsX(ctx context.Context) []int {
	v, err := dpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) IntX(ctx context.Context) int {
	v, err := dpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DivPayoutGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DivPayoutGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DivPayoutGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dpgb *DivPayoutGroupBy) BoolX(ctx context.Context) bool {
	v, err := dpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dpgb *DivPayoutGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dpgb.fields {
		if !divpayout.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dpgb *DivPayoutGroupBy) sqlQuery() *sql.Selector {
	selector := dpgb.sql.Select()
	aggregation := make([]string, 0, len(dpgb.fns))
	for _, fn := range dpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dpgb.fields)+len(dpgb.fns))
		for _, f := range dpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dpgb.fields...)...)
}

// DivPayoutSelect is the builder for selecting fields of DivPayout entities.
type DivPayoutSelect struct {
	*DivPayoutQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dps *DivPayoutSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dps.prepareQuery(ctx); err != nil {
		return err
	}
	dps.sql = dps.DivPayoutQuery.sqlQuery(ctx)
	return dps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dps *DivPayoutSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DivPayoutSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dps *DivPayoutSelect) StringsX(ctx context.Context) []string {
	v, err := dps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dps *DivPayoutSelect) StringX(ctx context.Context) string {
	v, err := dps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DivPayoutSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dps *DivPayoutSelect) IntsX(ctx context.Context) []int {
	v, err := dps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dps *DivPayoutSelect) IntX(ctx context.Context) int {
	v, err := dps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DivPayoutSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dps *DivPayoutSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dps *DivPayoutSelect) Float64X(ctx context.Context) float64 {
	v, err := dps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DivPayoutSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dps *DivPayoutSelect) BoolsX(ctx context.Context) []bool {
	v, err := dps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dps *DivPayoutSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{divpayout.Label}
	default:
		err = fmt.Errorf("ent: DivPayoutSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dps *DivPayoutSelect) BoolX(ctx context.Context) bool {
	v, err := dps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dps *DivPayoutSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dps.sql.Query()
	if err := dps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

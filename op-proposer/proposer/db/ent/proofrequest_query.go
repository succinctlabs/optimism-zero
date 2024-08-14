// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/moongate-forks/kona-sp1/op-proposer/proposer/db/ent/predicate"
	"github.com/moongate-forks/kona-sp1/op-proposer/proposer/db/ent/proofrequest"
)

// ProofRequestQuery is the builder for querying ProofRequest entities.
type ProofRequestQuery struct {
	config
	ctx        *QueryContext
	order      []proofrequest.OrderOption
	inters     []Interceptor
	predicates []predicate.ProofRequest
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProofRequestQuery builder.
func (prq *ProofRequestQuery) Where(ps ...predicate.ProofRequest) *ProofRequestQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit the number of records to be returned by this query.
func (prq *ProofRequestQuery) Limit(limit int) *ProofRequestQuery {
	prq.ctx.Limit = &limit
	return prq
}

// Offset to start from.
func (prq *ProofRequestQuery) Offset(offset int) *ProofRequestQuery {
	prq.ctx.Offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *ProofRequestQuery) Unique(unique bool) *ProofRequestQuery {
	prq.ctx.Unique = &unique
	return prq
}

// Order specifies how the records should be ordered.
func (prq *ProofRequestQuery) Order(o ...proofrequest.OrderOption) *ProofRequestQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// First returns the first ProofRequest entity from the query.
// Returns a *NotFoundError when no ProofRequest was found.
func (prq *ProofRequestQuery) First(ctx context.Context) (*ProofRequest, error) {
	nodes, err := prq.Limit(1).All(setContextOp(ctx, prq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{proofrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *ProofRequestQuery) FirstX(ctx context.Context) *ProofRequest {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProofRequest ID from the query.
// Returns a *NotFoundError when no ProofRequest ID was found.
func (prq *ProofRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{proofrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *ProofRequestQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProofRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProofRequest entity is found.
// Returns a *NotFoundError when no ProofRequest entities are found.
func (prq *ProofRequestQuery) Only(ctx context.Context) (*ProofRequest, error) {
	nodes, err := prq.Limit(2).All(setContextOp(ctx, prq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{proofrequest.Label}
	default:
		return nil, &NotSingularError{proofrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *ProofRequestQuery) OnlyX(ctx context.Context) *ProofRequest {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProofRequest ID in the query.
// Returns a *NotSingularError when more than one ProofRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *ProofRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{proofrequest.Label}
	default:
		err = &NotSingularError{proofrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *ProofRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProofRequests.
func (prq *ProofRequestQuery) All(ctx context.Context) ([]*ProofRequest, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryAll)
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProofRequest, *ProofRequestQuery]()
	return withInterceptors[[]*ProofRequest](ctx, prq, qr, prq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prq *ProofRequestQuery) AllX(ctx context.Context) []*ProofRequest {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProofRequest IDs.
func (prq *ProofRequestQuery) IDs(ctx context.Context) (ids []int, err error) {
	if prq.ctx.Unique == nil && prq.path != nil {
		prq.Unique(true)
	}
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryIDs)
	if err = prq.Select(proofrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *ProofRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *ProofRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryCount)
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prq, querierCount[*ProofRequestQuery](), prq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prq *ProofRequestQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *ProofRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryExist)
	switch _, err := prq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *ProofRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProofRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *ProofRequestQuery) Clone() *ProofRequestQuery {
	if prq == nil {
		return nil
	}
	return &ProofRequestQuery{
		config:     prq.config,
		ctx:        prq.ctx.Clone(),
		order:      append([]proofrequest.OrderOption{}, prq.order...),
		inters:     append([]Interceptor{}, prq.inters...),
		predicates: append([]predicate.ProofRequest{}, prq.predicates...),
		// clone intermediate query.
		sql:  prq.sql.Clone(),
		path: prq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type proofrequest.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProofRequest.Query().
//		GroupBy(proofrequest.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (prq *ProofRequestQuery) GroupBy(field string, fields ...string) *ProofRequestGroupBy {
	prq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProofRequestGroupBy{build: prq}
	grbuild.flds = &prq.ctx.Fields
	grbuild.label = proofrequest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type proofrequest.Type `json:"type,omitempty"`
//	}
//
//	client.ProofRequest.Query().
//		Select(proofrequest.FieldType).
//		Scan(ctx, &v)
func (prq *ProofRequestQuery) Select(fields ...string) *ProofRequestSelect {
	prq.ctx.Fields = append(prq.ctx.Fields, fields...)
	sbuild := &ProofRequestSelect{ProofRequestQuery: prq}
	sbuild.label = proofrequest.Label
	sbuild.flds, sbuild.scan = &prq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProofRequestSelect configured with the given aggregations.
func (prq *ProofRequestQuery) Aggregate(fns ...AggregateFunc) *ProofRequestSelect {
	return prq.Select().Aggregate(fns...)
}

func (prq *ProofRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prq); err != nil {
				return err
			}
		}
	}
	for _, f := range prq.ctx.Fields {
		if !proofrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *ProofRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProofRequest, error) {
	var (
		nodes = []*ProofRequest{}
		_spec = prq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProofRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProofRequest{config: prq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (prq *ProofRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	_spec.Node.Columns = prq.ctx.Fields
	if len(prq.ctx.Fields) > 0 {
		_spec.Unique = prq.ctx.Unique != nil && *prq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *ProofRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(proofrequest.Table, proofrequest.Columns, sqlgraph.NewFieldSpec(proofrequest.FieldID, field.TypeInt))
	_spec.From = prq.sql
	if unique := prq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prq.path != nil {
		_spec.Unique = true
	}
	if fields := prq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, proofrequest.FieldID)
		for i := range fields {
			if fields[i] != proofrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *ProofRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(proofrequest.Table)
	columns := prq.ctx.Fields
	if len(columns) == 0 {
		columns = proofrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.ctx.Unique != nil && *prq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProofRequestGroupBy is the group-by builder for ProofRequest entities.
type ProofRequestGroupBy struct {
	selector
	build *ProofRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *ProofRequestGroupBy) Aggregate(fns ...AggregateFunc) *ProofRequestGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the selector query and scans the result into the given value.
func (prgb *ProofRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prgb.build.ctx, ent.OpQueryGroupBy)
	if err := prgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProofRequestQuery, *ProofRequestGroupBy](ctx, prgb.build, prgb, prgb.build.inters, v)
}

func (prgb *ProofRequestGroupBy) sqlScan(ctx context.Context, root *ProofRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prgb.flds)+len(prgb.fns))
		for _, f := range *prgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProofRequestSelect is the builder for selecting fields of ProofRequest entities.
type ProofRequestSelect struct {
	*ProofRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prs *ProofRequestSelect) Aggregate(fns ...AggregateFunc) *ProofRequestSelect {
	prs.fns = append(prs.fns, fns...)
	return prs
}

// Scan applies the selector query and scans the result into the given value.
func (prs *ProofRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prs.ctx, ent.OpQuerySelect)
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProofRequestQuery, *ProofRequestSelect](ctx, prs.ProofRequestQuery, prs, prs.inters, v)
}

func (prs *ProofRequestSelect) sqlScan(ctx context.Context, root *ProofRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prs.fns))
	for _, fn := range prs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

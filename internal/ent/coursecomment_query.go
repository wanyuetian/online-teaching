// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"online-teaching/internal/ent/coursecomment"
	"online-teaching/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseCommentQuery is the builder for querying CourseComment entities.
type CourseCommentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CourseComment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CourseCommentQuery builder.
func (ccq *CourseCommentQuery) Where(ps ...predicate.CourseComment) *CourseCommentQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit adds a limit step to the query.
func (ccq *CourseCommentQuery) Limit(limit int) *CourseCommentQuery {
	ccq.limit = &limit
	return ccq
}

// Offset adds an offset step to the query.
func (ccq *CourseCommentQuery) Offset(offset int) *CourseCommentQuery {
	ccq.offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *CourseCommentQuery) Unique(unique bool) *CourseCommentQuery {
	ccq.unique = &unique
	return ccq
}

// Order adds an order step to the query.
func (ccq *CourseCommentQuery) Order(o ...OrderFunc) *CourseCommentQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// First returns the first CourseComment entity from the query.
// Returns a *NotFoundError when no CourseComment was found.
func (ccq *CourseCommentQuery) First(ctx context.Context) (*CourseComment, error) {
	nodes, err := ccq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coursecomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *CourseCommentQuery) FirstX(ctx context.Context) *CourseComment {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CourseComment ID from the query.
// Returns a *NotFoundError when no CourseComment ID was found.
func (ccq *CourseCommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coursecomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *CourseCommentQuery) FirstIDX(ctx context.Context) int {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CourseComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CourseComment entity is found.
// Returns a *NotFoundError when no CourseComment entities are found.
func (ccq *CourseCommentQuery) Only(ctx context.Context) (*CourseComment, error) {
	nodes, err := ccq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coursecomment.Label}
	default:
		return nil, &NotSingularError{coursecomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *CourseCommentQuery) OnlyX(ctx context.Context) *CourseComment {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CourseComment ID in the query.
// Returns a *NotSingularError when more than one CourseComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccq *CourseCommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ccq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coursecomment.Label}
	default:
		err = &NotSingularError{coursecomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *CourseCommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CourseComments.
func (ccq *CourseCommentQuery) All(ctx context.Context) ([]*CourseComment, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ccq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ccq *CourseCommentQuery) AllX(ctx context.Context) []*CourseComment {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CourseComment IDs.
func (ccq *CourseCommentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ccq.Select(coursecomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *CourseCommentQuery) IDsX(ctx context.Context) []int {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *CourseCommentQuery) Count(ctx context.Context) (int, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ccq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *CourseCommentQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *CourseCommentQuery) Exist(ctx context.Context) (bool, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ccq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *CourseCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CourseCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *CourseCommentQuery) Clone() *CourseCommentQuery {
	if ccq == nil {
		return nil
	}
	return &CourseCommentQuery{
		config:     ccq.config,
		limit:      ccq.limit,
		offset:     ccq.offset,
		order:      append([]OrderFunc{}, ccq.order...),
		predicates: append([]predicate.CourseComment{}, ccq.predicates...),
		// clone intermediate query.
		sql:    ccq.sql.Clone(),
		path:   ccq.path,
		unique: ccq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CourseComment.Query().
//		GroupBy(coursecomment.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ccq *CourseCommentQuery) GroupBy(field string, fields ...string) *CourseCommentGroupBy {
	grbuild := &CourseCommentGroupBy{config: ccq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ccq.sqlQuery(ctx), nil
	}
	grbuild.label = coursecomment.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.CourseComment.Query().
//		Select(coursecomment.FieldUsername).
//		Scan(ctx, &v)
//
func (ccq *CourseCommentQuery) Select(fields ...string) *CourseCommentSelect {
	ccq.fields = append(ccq.fields, fields...)
	selbuild := &CourseCommentSelect{CourseCommentQuery: ccq}
	selbuild.label = coursecomment.Label
	selbuild.flds, selbuild.scan = &ccq.fields, selbuild.Scan
	return selbuild
}

func (ccq *CourseCommentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ccq.fields {
		if !coursecomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *CourseCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CourseComment, error) {
	var (
		nodes = []*CourseComment{}
		_spec = ccq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CourseComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CourseComment{config: ccq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccq *CourseCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	_spec.Node.Columns = ccq.fields
	if len(ccq.fields) > 0 {
		_spec.Unique = ccq.unique != nil && *ccq.unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *CourseCommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ccq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ccq *CourseCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coursecomment.Table,
			Columns: coursecomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coursecomment.FieldID,
			},
		},
		From:   ccq.sql,
		Unique: true,
	}
	if unique := ccq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ccq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coursecomment.FieldID)
		for i := range fields {
			if fields[i] != coursecomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *CourseCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(coursecomment.Table)
	columns := ccq.fields
	if len(columns) == 0 {
		columns = coursecomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.unique != nil && *ccq.unique {
		selector.Distinct()
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CourseCommentGroupBy is the group-by builder for CourseComment entities.
type CourseCommentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *CourseCommentGroupBy) Aggregate(fns ...AggregateFunc) *CourseCommentGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ccgb *CourseCommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ccgb.path(ctx)
	if err != nil {
		return err
	}
	ccgb.sql = query
	return ccgb.sqlScan(ctx, v)
}

func (ccgb *CourseCommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ccgb.fields {
		if !coursecomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ccgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ccgb *CourseCommentGroupBy) sqlQuery() *sql.Selector {
	selector := ccgb.sql.Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ccgb.fields)+len(ccgb.fns))
		for _, f := range ccgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ccgb.fields...)...)
}

// CourseCommentSelect is the builder for selecting fields of CourseComment entities.
type CourseCommentSelect struct {
	*CourseCommentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *CourseCommentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	ccs.sql = ccs.CourseCommentQuery.sqlQuery(ctx)
	return ccs.sqlScan(ctx, v)
}

func (ccs *CourseCommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ccs.sql.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
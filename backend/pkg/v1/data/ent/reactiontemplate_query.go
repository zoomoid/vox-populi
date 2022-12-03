// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reactiontemplate"
)

// ReactionTemplateQuery is the builder for querying ReactionTemplate entities.
type ReactionTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ReactionTemplate
	withPoll   *PollQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReactionTemplateQuery builder.
func (rtq *ReactionTemplateQuery) Where(ps ...predicate.ReactionTemplate) *ReactionTemplateQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit adds a limit step to the query.
func (rtq *ReactionTemplateQuery) Limit(limit int) *ReactionTemplateQuery {
	rtq.limit = &limit
	return rtq
}

// Offset adds an offset step to the query.
func (rtq *ReactionTemplateQuery) Offset(offset int) *ReactionTemplateQuery {
	rtq.offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *ReactionTemplateQuery) Unique(unique bool) *ReactionTemplateQuery {
	rtq.unique = &unique
	return rtq
}

// Order adds an order step to the query.
func (rtq *ReactionTemplateQuery) Order(o ...OrderFunc) *ReactionTemplateQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryPoll chains the current query on the "poll" edge.
func (rtq *ReactionTemplateQuery) QueryPoll() *PollQuery {
	query := &PollQuery{config: rtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reactiontemplate.Table, reactiontemplate.FieldID, selector),
			sqlgraph.To(poll.Table, poll.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reactiontemplate.PollTable, reactiontemplate.PollColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReactionTemplate entity from the query.
// Returns a *NotFoundError when no ReactionTemplate was found.
func (rtq *ReactionTemplateQuery) First(ctx context.Context) (*ReactionTemplate, error) {
	nodes, err := rtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reactiontemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) FirstX(ctx context.Context) *ReactionTemplate {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReactionTemplate ID from the query.
// Returns a *NotFoundError when no ReactionTemplate ID was found.
func (rtq *ReactionTemplateQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reactiontemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) FirstIDX(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReactionTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReactionTemplate entity is found.
// Returns a *NotFoundError when no ReactionTemplate entities are found.
func (rtq *ReactionTemplateQuery) Only(ctx context.Context) (*ReactionTemplate, error) {
	nodes, err := rtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reactiontemplate.Label}
	default:
		return nil, &NotSingularError{reactiontemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) OnlyX(ctx context.Context) *ReactionTemplate {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReactionTemplate ID in the query.
// Returns a *NotSingularError when more than one ReactionTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *ReactionTemplateQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reactiontemplate.Label}
	default:
		err = &NotSingularError{reactiontemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReactionTemplates.
func (rtq *ReactionTemplateQuery) All(ctx context.Context) ([]*ReactionTemplate, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) AllX(ctx context.Context) []*ReactionTemplate {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReactionTemplate IDs.
func (rtq *ReactionTemplateQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rtq.Select(reactiontemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *ReactionTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *ReactionTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *ReactionTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReactionTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *ReactionTemplateQuery) Clone() *ReactionTemplateQuery {
	if rtq == nil {
		return nil
	}
	return &ReactionTemplateQuery{
		config:     rtq.config,
		limit:      rtq.limit,
		offset:     rtq.offset,
		order:      append([]OrderFunc{}, rtq.order...),
		predicates: append([]predicate.ReactionTemplate{}, rtq.predicates...),
		withPoll:   rtq.withPoll.Clone(),
		// clone intermediate query.
		sql:    rtq.sql.Clone(),
		path:   rtq.path,
		unique: rtq.unique,
	}
}

// WithPoll tells the query-builder to eager-load the nodes that are connected to
// the "poll" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *ReactionTemplateQuery) WithPoll(opts ...func(*PollQuery)) *ReactionTemplateQuery {
	query := &PollQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	rtq.withPoll = query
	return rtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReactionTemplate.Query().
//		GroupBy(reactiontemplate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rtq *ReactionTemplateQuery) GroupBy(field string, fields ...string) *ReactionTemplateGroupBy {
	grbuild := &ReactionTemplateGroupBy{config: rtq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rtq.sqlQuery(ctx), nil
	}
	grbuild.label = reactiontemplate.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ReactionTemplate.Query().
//		Select(reactiontemplate.FieldCreatedAt).
//		Scan(ctx, &v)
func (rtq *ReactionTemplateQuery) Select(fields ...string) *ReactionTemplateSelect {
	rtq.fields = append(rtq.fields, fields...)
	selbuild := &ReactionTemplateSelect{ReactionTemplateQuery: rtq}
	selbuild.label = reactiontemplate.Label
	selbuild.flds, selbuild.scan = &rtq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ReactionTemplateSelect configured with the given aggregations.
func (rtq *ReactionTemplateQuery) Aggregate(fns ...AggregateFunc) *ReactionTemplateSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *ReactionTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rtq.fields {
		if !reactiontemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *ReactionTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReactionTemplate, error) {
	var (
		nodes       = []*ReactionTemplate{}
		withFKs     = rtq.withFKs
		_spec       = rtq.querySpec()
		loadedTypes = [1]bool{
			rtq.withPoll != nil,
		}
	)
	if rtq.withPoll != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, reactiontemplate.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReactionTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReactionTemplate{config: rtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rtq.withPoll; query != nil {
		if err := rtq.loadPoll(ctx, query, nodes, nil,
			func(n *ReactionTemplate, e *Poll) { n.Edges.Poll = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rtq *ReactionTemplateQuery) loadPoll(ctx context.Context, query *PollQuery, nodes []*ReactionTemplate, init func(*ReactionTemplate), assign func(*ReactionTemplate, *Poll)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ReactionTemplate)
	for i := range nodes {
		if nodes[i].poll_reaction_templates == nil {
			continue
		}
		fk := *nodes[i].poll_reaction_templates
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(poll.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_reaction_templates" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rtq *ReactionTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	_spec.Node.Columns = rtq.fields
	if len(rtq.fields) > 0 {
		_spec.Unique = rtq.unique != nil && *rtq.unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *ReactionTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rtq *ReactionTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reactiontemplate.Table,
			Columns: reactiontemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reactiontemplate.FieldID,
			},
		},
		From:   rtq.sql,
		Unique: true,
	}
	if unique := rtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reactiontemplate.FieldID)
		for i := range fields {
			if fields[i] != reactiontemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *ReactionTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(reactiontemplate.Table)
	columns := rtq.fields
	if len(columns) == 0 {
		columns = reactiontemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.unique != nil && *rtq.unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReactionTemplateGroupBy is the group-by builder for ReactionTemplate entities.
type ReactionTemplateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *ReactionTemplateGroupBy) Aggregate(fns ...AggregateFunc) *ReactionTemplateGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rtgb *ReactionTemplateGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rtgb.path(ctx)
	if err != nil {
		return err
	}
	rtgb.sql = query
	return rtgb.sqlScan(ctx, v)
}

func (rtgb *ReactionTemplateGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rtgb.fields {
		if !reactiontemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rtgb *ReactionTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := rtgb.sql.Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rtgb.fields)+len(rtgb.fns))
		for _, f := range rtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rtgb.fields...)...)
}

// ReactionTemplateSelect is the builder for selecting fields of ReactionTemplate entities.
type ReactionTemplateSelect struct {
	*ReactionTemplateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *ReactionTemplateSelect) Aggregate(fns ...AggregateFunc) *ReactionTemplateSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *ReactionTemplateSelect) Scan(ctx context.Context, v any) error {
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	rts.sql = rts.ReactionTemplateQuery.sqlQuery(ctx)
	return rts.sqlScan(ctx, v)
}

func (rts *ReactionTemplateSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(rts.sql))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rts.sql.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
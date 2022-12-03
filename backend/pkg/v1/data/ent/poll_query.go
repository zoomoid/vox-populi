// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reaction"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reactiontemplate"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/vote"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/votetemplate"
)

// PollQuery is the builder for querying Poll entities.
type PollQuery struct {
	config
	limit                 *int
	offset                *int
	unique                *bool
	order                 []OrderFunc
	fields                []string
	predicates            []predicate.Poll
	withVoteTemplates     *VoteTemplateQuery
	withVotes             *VoteQuery
	withReactionTemplates *ReactionTemplateQuery
	withReactions         *ReactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PollQuery builder.
func (pq *PollQuery) Where(ps ...predicate.Poll) *PollQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PollQuery) Limit(limit int) *PollQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PollQuery) Offset(offset int) *PollQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PollQuery) Unique(unique bool) *PollQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PollQuery) Order(o ...OrderFunc) *PollQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryVoteTemplates chains the current query on the "vote_templates" edge.
func (pq *PollQuery) QueryVoteTemplates() *VoteTemplateQuery {
	query := &VoteTemplateQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, selector),
			sqlgraph.To(votetemplate.Table, votetemplate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, poll.VoteTemplatesTable, poll.VoteTemplatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVotes chains the current query on the "votes" edge.
func (pq *PollQuery) QueryVotes() *VoteQuery {
	query := &VoteQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, selector),
			sqlgraph.To(vote.Table, vote.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, poll.VotesTable, poll.VotesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReactionTemplates chains the current query on the "reaction_templates" edge.
func (pq *PollQuery) QueryReactionTemplates() *ReactionTemplateQuery {
	query := &ReactionTemplateQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, selector),
			sqlgraph.To(reactiontemplate.Table, reactiontemplate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, poll.ReactionTemplatesTable, poll.ReactionTemplatesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReactions chains the current query on the "reactions" edge.
func (pq *PollQuery) QueryReactions() *ReactionQuery {
	query := &ReactionQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, selector),
			sqlgraph.To(reaction.Table, reaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, poll.ReactionsTable, poll.ReactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Poll entity from the query.
// Returns a *NotFoundError when no Poll was found.
func (pq *PollQuery) First(ctx context.Context) (*Poll, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{poll.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PollQuery) FirstX(ctx context.Context) *Poll {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Poll ID from the query.
// Returns a *NotFoundError when no Poll ID was found.
func (pq *PollQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{poll.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PollQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Poll entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Poll entity is found.
// Returns a *NotFoundError when no Poll entities are found.
func (pq *PollQuery) Only(ctx context.Context) (*Poll, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{poll.Label}
	default:
		return nil, &NotSingularError{poll.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PollQuery) OnlyX(ctx context.Context) *Poll {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Poll ID in the query.
// Returns a *NotSingularError when more than one Poll ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PollQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{poll.Label}
	default:
		err = &NotSingularError{poll.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PollQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Polls.
func (pq *PollQuery) All(ctx context.Context) ([]*Poll, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PollQuery) AllX(ctx context.Context) []*Poll {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Poll IDs.
func (pq *PollQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(poll.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PollQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PollQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PollQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PollQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PollQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PollQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PollQuery) Clone() *PollQuery {
	if pq == nil {
		return nil
	}
	return &PollQuery{
		config:                pq.config,
		limit:                 pq.limit,
		offset:                pq.offset,
		order:                 append([]OrderFunc{}, pq.order...),
		predicates:            append([]predicate.Poll{}, pq.predicates...),
		withVoteTemplates:     pq.withVoteTemplates.Clone(),
		withVotes:             pq.withVotes.Clone(),
		withReactionTemplates: pq.withReactionTemplates.Clone(),
		withReactions:         pq.withReactions.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithVoteTemplates tells the query-builder to eager-load the nodes that are connected to
// the "vote_templates" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PollQuery) WithVoteTemplates(opts ...func(*VoteTemplateQuery)) *PollQuery {
	query := &VoteTemplateQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withVoteTemplates = query
	return pq
}

// WithVotes tells the query-builder to eager-load the nodes that are connected to
// the "votes" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PollQuery) WithVotes(opts ...func(*VoteQuery)) *PollQuery {
	query := &VoteQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withVotes = query
	return pq
}

// WithReactionTemplates tells the query-builder to eager-load the nodes that are connected to
// the "reaction_templates" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PollQuery) WithReactionTemplates(opts ...func(*ReactionTemplateQuery)) *PollQuery {
	query := &ReactionTemplateQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withReactionTemplates = query
	return pq
}

// WithReactions tells the query-builder to eager-load the nodes that are connected to
// the "reactions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PollQuery) WithReactions(opts ...func(*ReactionQuery)) *PollQuery {
	query := &ReactionQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withReactions = query
	return pq
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
//	client.Poll.Query().
//		GroupBy(poll.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PollQuery) GroupBy(field string, fields ...string) *PollGroupBy {
	grbuild := &PollGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = poll.Label
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
//	client.Poll.Query().
//		Select(poll.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PollQuery) Select(fields ...string) *PollSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PollSelect{PollQuery: pq}
	selbuild.label = poll.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a PollSelect configured with the given aggregations.
func (pq *PollQuery) Aggregate(fns ...AggregateFunc) *PollSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PollQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !poll.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PollQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Poll, error) {
	var (
		nodes       = []*Poll{}
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withVoteTemplates != nil,
			pq.withVotes != nil,
			pq.withReactionTemplates != nil,
			pq.withReactions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Poll).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Poll{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withVoteTemplates; query != nil {
		if err := pq.loadVoteTemplates(ctx, query, nodes,
			func(n *Poll) { n.Edges.VoteTemplates = []*VoteTemplate{} },
			func(n *Poll, e *VoteTemplate) { n.Edges.VoteTemplates = append(n.Edges.VoteTemplates, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withVotes; query != nil {
		if err := pq.loadVotes(ctx, query, nodes,
			func(n *Poll) { n.Edges.Votes = []*Vote{} },
			func(n *Poll, e *Vote) { n.Edges.Votes = append(n.Edges.Votes, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withReactionTemplates; query != nil {
		if err := pq.loadReactionTemplates(ctx, query, nodes,
			func(n *Poll) { n.Edges.ReactionTemplates = []*ReactionTemplate{} },
			func(n *Poll, e *ReactionTemplate) { n.Edges.ReactionTemplates = append(n.Edges.ReactionTemplates, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withReactions; query != nil {
		if err := pq.loadReactions(ctx, query, nodes,
			func(n *Poll) { n.Edges.Reactions = []*Reaction{} },
			func(n *Poll, e *Reaction) { n.Edges.Reactions = append(n.Edges.Reactions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PollQuery) loadVoteTemplates(ctx context.Context, query *VoteTemplateQuery, nodes []*Poll, init func(*Poll), assign func(*Poll, *VoteTemplate)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Poll)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.VoteTemplate(func(s *sql.Selector) {
		s.Where(sql.InValues(poll.VoteTemplatesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.poll_vote_templates
		if fk == nil {
			return fmt.Errorf(`foreign-key "poll_vote_templates" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_vote_templates" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PollQuery) loadVotes(ctx context.Context, query *VoteQuery, nodes []*Poll, init func(*Poll), assign func(*Poll, *Vote)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Poll)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Vote(func(s *sql.Selector) {
		s.Where(sql.InValues(poll.VotesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.poll_votes
		if fk == nil {
			return fmt.Errorf(`foreign-key "poll_votes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_votes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PollQuery) loadReactionTemplates(ctx context.Context, query *ReactionTemplateQuery, nodes []*Poll, init func(*Poll), assign func(*Poll, *ReactionTemplate)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Poll)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ReactionTemplate(func(s *sql.Selector) {
		s.Where(sql.InValues(poll.ReactionTemplatesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.poll_reaction_templates
		if fk == nil {
			return fmt.Errorf(`foreign-key "poll_reaction_templates" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_reaction_templates" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PollQuery) loadReactions(ctx context.Context, query *ReactionQuery, nodes []*Poll, init func(*Poll), assign func(*Poll, *Reaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Poll)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.InValues(poll.ReactionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.poll_reactions
		if fk == nil {
			return fmt.Errorf(`foreign-key "poll_reactions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "poll_reactions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PollQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PollQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (pq *PollQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   poll.Table,
			Columns: poll.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: poll.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poll.FieldID)
		for i := range fields {
			if fields[i] != poll.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PollQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(poll.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = poll.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PollGroupBy is the group-by builder for Poll entities.
type PollGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PollGroupBy) Aggregate(fns ...AggregateFunc) *PollGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PollGroupBy) Scan(ctx context.Context, v any) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PollGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range pgb.fields {
		if !poll.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PollGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PollSelect is the builder for selecting fields of Poll entities.
type PollSelect struct {
	*PollQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PollSelect) Aggregate(fns ...AggregateFunc) *PollSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PollSelect) Scan(ctx context.Context, v any) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PollQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PollSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(ps.sql))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

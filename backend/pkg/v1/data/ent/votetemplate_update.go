// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/votetemplate"
)

// VoteTemplateUpdate is the builder for updating VoteTemplate entities.
type VoteTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *VoteTemplateMutation
}

// Where appends a list predicates to the VoteTemplateUpdate builder.
func (vtu *VoteTemplateUpdate) Where(ps ...predicate.VoteTemplate) *VoteTemplateUpdate {
	vtu.mutation.Where(ps...)
	return vtu
}

// SetUpdatedAt sets the "updated_at" field.
func (vtu *VoteTemplateUpdate) SetUpdatedAt(t time.Time) *VoteTemplateUpdate {
	vtu.mutation.SetUpdatedAt(t)
	return vtu
}

// SetAnswer sets the "answer" field.
func (vtu *VoteTemplateUpdate) SetAnswer(s string) *VoteTemplateUpdate {
	vtu.mutation.SetAnswer(s)
	return vtu
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (vtu *VoteTemplateUpdate) SetPollID(id int) *VoteTemplateUpdate {
	vtu.mutation.SetPollID(id)
	return vtu
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (vtu *VoteTemplateUpdate) SetNillablePollID(id *int) *VoteTemplateUpdate {
	if id != nil {
		vtu = vtu.SetPollID(*id)
	}
	return vtu
}

// SetPoll sets the "poll" edge to the Poll entity.
func (vtu *VoteTemplateUpdate) SetPoll(p *Poll) *VoteTemplateUpdate {
	return vtu.SetPollID(p.ID)
}

// Mutation returns the VoteTemplateMutation object of the builder.
func (vtu *VoteTemplateUpdate) Mutation() *VoteTemplateMutation {
	return vtu.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (vtu *VoteTemplateUpdate) ClearPoll() *VoteTemplateUpdate {
	vtu.mutation.ClearPoll()
	return vtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vtu *VoteTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	vtu.defaults()
	if len(vtu.hooks) == 0 {
		if err = vtu.check(); err != nil {
			return 0, err
		}
		affected, err = vtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VoteTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vtu.check(); err != nil {
				return 0, err
			}
			vtu.mutation = mutation
			affected, err = vtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vtu.hooks) - 1; i >= 0; i-- {
			if vtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vtu *VoteTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := vtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vtu *VoteTemplateUpdate) Exec(ctx context.Context) error {
	_, err := vtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtu *VoteTemplateUpdate) ExecX(ctx context.Context) {
	if err := vtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtu *VoteTemplateUpdate) defaults() {
	if _, ok := vtu.mutation.UpdatedAt(); !ok {
		v := votetemplate.UpdateDefaultUpdatedAt()
		vtu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vtu *VoteTemplateUpdate) check() error {
	if v, ok := vtu.mutation.Answer(); ok {
		if err := votetemplate.AnswerValidator(v); err != nil {
			return &ValidationError{Name: "answer", err: fmt.Errorf(`ent: validator failed for field "VoteTemplate.answer": %w`, err)}
		}
	}
	return nil
}

func (vtu *VoteTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   votetemplate.Table,
			Columns: votetemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: votetemplate.FieldID,
			},
		},
	}
	if ps := vtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtu.mutation.UpdatedAt(); ok {
		_spec.SetField(votetemplate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vtu.mutation.Answer(); ok {
		_spec.SetField(votetemplate.FieldAnswer, field.TypeString, value)
	}
	if vtu.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   votetemplate.PollTable,
			Columns: []string{votetemplate.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: poll.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vtu.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   votetemplate.PollTable,
			Columns: []string{votetemplate.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: poll.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{votetemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// VoteTemplateUpdateOne is the builder for updating a single VoteTemplate entity.
type VoteTemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VoteTemplateMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (vtuo *VoteTemplateUpdateOne) SetUpdatedAt(t time.Time) *VoteTemplateUpdateOne {
	vtuo.mutation.SetUpdatedAt(t)
	return vtuo
}

// SetAnswer sets the "answer" field.
func (vtuo *VoteTemplateUpdateOne) SetAnswer(s string) *VoteTemplateUpdateOne {
	vtuo.mutation.SetAnswer(s)
	return vtuo
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (vtuo *VoteTemplateUpdateOne) SetPollID(id int) *VoteTemplateUpdateOne {
	vtuo.mutation.SetPollID(id)
	return vtuo
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (vtuo *VoteTemplateUpdateOne) SetNillablePollID(id *int) *VoteTemplateUpdateOne {
	if id != nil {
		vtuo = vtuo.SetPollID(*id)
	}
	return vtuo
}

// SetPoll sets the "poll" edge to the Poll entity.
func (vtuo *VoteTemplateUpdateOne) SetPoll(p *Poll) *VoteTemplateUpdateOne {
	return vtuo.SetPollID(p.ID)
}

// Mutation returns the VoteTemplateMutation object of the builder.
func (vtuo *VoteTemplateUpdateOne) Mutation() *VoteTemplateMutation {
	return vtuo.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (vtuo *VoteTemplateUpdateOne) ClearPoll() *VoteTemplateUpdateOne {
	vtuo.mutation.ClearPoll()
	return vtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vtuo *VoteTemplateUpdateOne) Select(field string, fields ...string) *VoteTemplateUpdateOne {
	vtuo.fields = append([]string{field}, fields...)
	return vtuo
}

// Save executes the query and returns the updated VoteTemplate entity.
func (vtuo *VoteTemplateUpdateOne) Save(ctx context.Context) (*VoteTemplate, error) {
	var (
		err  error
		node *VoteTemplate
	)
	vtuo.defaults()
	if len(vtuo.hooks) == 0 {
		if err = vtuo.check(); err != nil {
			return nil, err
		}
		node, err = vtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VoteTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vtuo.check(); err != nil {
				return nil, err
			}
			vtuo.mutation = mutation
			node, err = vtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vtuo.hooks) - 1; i >= 0; i-- {
			if vtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vtuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vtuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*VoteTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from VoteTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vtuo *VoteTemplateUpdateOne) SaveX(ctx context.Context) *VoteTemplate {
	node, err := vtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vtuo *VoteTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := vtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtuo *VoteTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := vtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtuo *VoteTemplateUpdateOne) defaults() {
	if _, ok := vtuo.mutation.UpdatedAt(); !ok {
		v := votetemplate.UpdateDefaultUpdatedAt()
		vtuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vtuo *VoteTemplateUpdateOne) check() error {
	if v, ok := vtuo.mutation.Answer(); ok {
		if err := votetemplate.AnswerValidator(v); err != nil {
			return &ValidationError{Name: "answer", err: fmt.Errorf(`ent: validator failed for field "VoteTemplate.answer": %w`, err)}
		}
	}
	return nil
}

func (vtuo *VoteTemplateUpdateOne) sqlSave(ctx context.Context) (_node *VoteTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   votetemplate.Table,
			Columns: votetemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: votetemplate.FieldID,
			},
		},
	}
	id, ok := vtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VoteTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, votetemplate.FieldID)
		for _, f := range fields {
			if !votetemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != votetemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(votetemplate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vtuo.mutation.Answer(); ok {
		_spec.SetField(votetemplate.FieldAnswer, field.TypeString, value)
	}
	if vtuo.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   votetemplate.PollTable,
			Columns: []string{votetemplate.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: poll.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vtuo.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   votetemplate.PollTable,
			Columns: []string{votetemplate.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: poll.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VoteTemplate{config: vtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{votetemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

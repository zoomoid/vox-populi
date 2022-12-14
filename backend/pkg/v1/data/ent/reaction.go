// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reaction"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reactiontemplate"
)

// Reaction is the model entity for the Reaction schema.
type Reaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReactionQuery when eager-loading is set.
	Edges             ReactionEdges `json:"edges"`
	poll_reactions    *int
	reaction_template *int
}

// ReactionEdges holds the relations/edges for other nodes in the graph.
type ReactionEdges struct {
	// Template holds the value of the template edge.
	Template *ReactionTemplate `json:"template,omitempty"`
	// Poll holds the value of the poll edge.
	Poll *Poll `json:"poll,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReactionEdges) TemplateOrErr() (*ReactionTemplate, error) {
	if e.loadedTypes[0] {
		if e.Template == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: reactiontemplate.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// PollOrErr returns the Poll value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReactionEdges) PollOrErr() (*Poll, error) {
	if e.loadedTypes[1] {
		if e.Poll == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: poll.Label}
		}
		return e.Poll, nil
	}
	return nil, &NotLoadedError{edge: "poll"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reaction.FieldID:
			values[i] = new(sql.NullInt64)
		case reaction.FieldCreatedAt, reaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case reaction.ForeignKeys[0]: // poll_reactions
			values[i] = new(sql.NullInt64)
		case reaction.ForeignKeys[1]: // reaction_template
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Reaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reaction fields.
func (r *Reaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case reaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field poll_reactions", value)
			} else if value.Valid {
				r.poll_reactions = new(int)
				*r.poll_reactions = int(value.Int64)
			}
		case reaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field reaction_template", value)
			} else if value.Valid {
				r.reaction_template = new(int)
				*r.reaction_template = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTemplate queries the "template" edge of the Reaction entity.
func (r *Reaction) QueryTemplate() *ReactionTemplateQuery {
	return (&ReactionClient{config: r.config}).QueryTemplate(r)
}

// QueryPoll queries the "poll" edge of the Reaction entity.
func (r *Reaction) QueryPoll() *PollQuery {
	return (&ReactionClient{config: r.config}).QueryPoll(r)
}

// Update returns a builder for updating this Reaction.
// Note that you need to call Reaction.Unwrap() before calling this method if this Reaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reaction) Update() *ReactionUpdateOne {
	return (&ReactionClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Reaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reaction) Unwrap() *Reaction {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reaction is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reaction) String() string {
	var builder strings.Builder
	builder.WriteString("Reaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reactions is a parsable slice of Reaction.
type Reactions []*Reaction

func (r Reactions) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

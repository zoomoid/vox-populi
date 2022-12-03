// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/vote"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/votetemplate"
)

// Vote is the model entity for the Vote schema.
type Vote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VoteQuery when eager-loading is set.
	Edges         VoteEdges `json:"edges"`
	poll_votes    *int
	vote_template *int
}

// VoteEdges holds the relations/edges for other nodes in the graph.
type VoteEdges struct {
	// Template holds the value of the template edge.
	Template *VoteTemplate `json:"template,omitempty"`
	// Poll holds the value of the poll edge.
	Poll *Poll `json:"poll,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VoteEdges) TemplateOrErr() (*VoteTemplate, error) {
	if e.loadedTypes[0] {
		if e.Template == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: votetemplate.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// PollOrErr returns the Poll value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VoteEdges) PollOrErr() (*Poll, error) {
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
func (*Vote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vote.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case vote.FieldID:
			values[i] = new(sql.NullInt64)
		case vote.FieldNonce:
			values[i] = new(sql.NullString)
		case vote.FieldCreatedAt, vote.FieldUpdatedAt, vote.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case vote.ForeignKeys[0]: // poll_votes
			values[i] = new(sql.NullInt64)
		case vote.ForeignKeys[1]: // vote_template
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vote fields.
func (v *Vote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vote.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case vote.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case vote.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				v.IsDeleted = value.Bool
			}
		case vote.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				v.DeletedAt = new(time.Time)
				*v.DeletedAt = value.Time
			}
		case vote.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				v.Nonce = value.String
			}
		case vote.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field poll_votes", value)
			} else if value.Valid {
				v.poll_votes = new(int)
				*v.poll_votes = int(value.Int64)
			}
		case vote.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field vote_template", value)
			} else if value.Valid {
				v.vote_template = new(int)
				*v.vote_template = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTemplate queries the "template" edge of the Vote entity.
func (v *Vote) QueryTemplate() *VoteTemplateQuery {
	return (&VoteClient{config: v.config}).QueryTemplate(v)
}

// QueryPoll queries the "poll" edge of the Vote entity.
func (v *Vote) QueryPoll() *PollQuery {
	return (&VoteClient{config: v.config}).QueryPoll(v)
}

// Update returns a builder for updating this Vote.
// Note that you need to call Vote.Unwrap() before calling this method if this Vote
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vote) Update() *VoteUpdateOne {
	return (&VoteClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Vote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vote) Unwrap() *Vote {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vote is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vote) String() string {
	var builder strings.Builder
	builder.WriteString("Vote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", v.IsDeleted))
	builder.WriteString(", ")
	if v := v.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(v.Nonce)
	builder.WriteByte(')')
	return builder.String()
}

// Votes is a parsable slice of Vote.
type Votes []*Vote

func (v Votes) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}

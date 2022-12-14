// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
)

// Poll is the model entity for the Poll schema.
type Poll struct {
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
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsLive holds the value of the "is_live" field.
	IsLive bool `json:"is_live,omitempty"`
	// CanVote holds the value of the "can_vote" field.
	CanVote bool `json:"can_vote,omitempty"`
	// CanReact holds the value of the "can_react" field.
	CanReact bool `json:"can_react,omitempty"`
	// CanSeeResults holds the value of the "can_see_results" field.
	CanSeeResults bool `json:"can_see_results,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// UnpublishedAt holds the value of the "unpublished_at" field.
	UnpublishedAt *time.Time `json:"unpublished_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PollQuery when eager-loading is set.
	Edges PollEdges `json:"edges"`
}

// PollEdges holds the relations/edges for other nodes in the graph.
type PollEdges struct {
	// VoteTemplates holds the value of the vote_templates edge.
	VoteTemplates []*VoteTemplate `json:"vote_templates,omitempty"`
	// Votes holds the value of the votes edge.
	Votes []*Vote `json:"votes,omitempty"`
	// ReactionTemplates holds the value of the reaction_templates edge.
	ReactionTemplates []*ReactionTemplate `json:"reaction_templates,omitempty"`
	// Reactions holds the value of the reactions edge.
	Reactions []*Reaction `json:"reactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// VoteTemplatesOrErr returns the VoteTemplates value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) VoteTemplatesOrErr() ([]*VoteTemplate, error) {
	if e.loadedTypes[0] {
		return e.VoteTemplates, nil
	}
	return nil, &NotLoadedError{edge: "vote_templates"}
}

// VotesOrErr returns the Votes value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) VotesOrErr() ([]*Vote, error) {
	if e.loadedTypes[1] {
		return e.Votes, nil
	}
	return nil, &NotLoadedError{edge: "votes"}
}

// ReactionTemplatesOrErr returns the ReactionTemplates value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) ReactionTemplatesOrErr() ([]*ReactionTemplate, error) {
	if e.loadedTypes[2] {
		return e.ReactionTemplates, nil
	}
	return nil, &NotLoadedError{edge: "reaction_templates"}
}

// ReactionsOrErr returns the Reactions value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) ReactionsOrErr() ([]*Reaction, error) {
	if e.loadedTypes[3] {
		return e.Reactions, nil
	}
	return nil, &NotLoadedError{edge: "reactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Poll) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case poll.FieldIsDeleted, poll.FieldIsLive, poll.FieldCanVote, poll.FieldCanReact, poll.FieldCanSeeResults:
			values[i] = new(sql.NullBool)
		case poll.FieldID:
			values[i] = new(sql.NullInt64)
		case poll.FieldTitle, poll.FieldDescription:
			values[i] = new(sql.NullString)
		case poll.FieldCreatedAt, poll.FieldUpdatedAt, poll.FieldDeletedAt, poll.FieldPublishedAt, poll.FieldUnpublishedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Poll", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Poll fields.
func (po *Poll) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case poll.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case poll.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case poll.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case poll.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				po.IsDeleted = value.Bool
			}
		case poll.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				po.DeletedAt = new(time.Time)
				*po.DeletedAt = value.Time
			}
		case poll.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case poll.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				po.Description = value.String
			}
		case poll.FieldIsLive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_live", values[i])
			} else if value.Valid {
				po.IsLive = value.Bool
			}
		case poll.FieldCanVote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_vote", values[i])
			} else if value.Valid {
				po.CanVote = value.Bool
			}
		case poll.FieldCanReact:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_react", values[i])
			} else if value.Valid {
				po.CanReact = value.Bool
			}
		case poll.FieldCanSeeResults:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_see_results", values[i])
			} else if value.Valid {
				po.CanSeeResults = value.Bool
			}
		case poll.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				po.PublishedAt = new(time.Time)
				*po.PublishedAt = value.Time
			}
		case poll.FieldUnpublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field unpublished_at", values[i])
			} else if value.Valid {
				po.UnpublishedAt = new(time.Time)
				*po.UnpublishedAt = value.Time
			}
		}
	}
	return nil
}

// QueryVoteTemplates queries the "vote_templates" edge of the Poll entity.
func (po *Poll) QueryVoteTemplates() *VoteTemplateQuery {
	return (&PollClient{config: po.config}).QueryVoteTemplates(po)
}

// QueryVotes queries the "votes" edge of the Poll entity.
func (po *Poll) QueryVotes() *VoteQuery {
	return (&PollClient{config: po.config}).QueryVotes(po)
}

// QueryReactionTemplates queries the "reaction_templates" edge of the Poll entity.
func (po *Poll) QueryReactionTemplates() *ReactionTemplateQuery {
	return (&PollClient{config: po.config}).QueryReactionTemplates(po)
}

// QueryReactions queries the "reactions" edge of the Poll entity.
func (po *Poll) QueryReactions() *ReactionQuery {
	return (&PollClient{config: po.config}).QueryReactions(po)
}

// Update returns a builder for updating this Poll.
// Note that you need to call Poll.Unwrap() before calling this method if this Poll
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Poll) Update() *PollUpdateOne {
	return (&PollClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Poll entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Poll) Unwrap() *Poll {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Poll is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Poll) String() string {
	var builder strings.Builder
	builder.WriteString("Poll(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", po.IsDeleted))
	builder.WriteString(", ")
	if v := po.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(po.Description)
	builder.WriteString(", ")
	builder.WriteString("is_live=")
	builder.WriteString(fmt.Sprintf("%v", po.IsLive))
	builder.WriteString(", ")
	builder.WriteString("can_vote=")
	builder.WriteString(fmt.Sprintf("%v", po.CanVote))
	builder.WriteString(", ")
	builder.WriteString("can_react=")
	builder.WriteString(fmt.Sprintf("%v", po.CanReact))
	builder.WriteString(", ")
	builder.WriteString("can_see_results=")
	builder.WriteString(fmt.Sprintf("%v", po.CanSeeResults))
	builder.WriteString(", ")
	if v := po.PublishedAt; v != nil {
		builder.WriteString("published_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := po.UnpublishedAt; v != nil {
		builder.WriteString("unpublished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Polls is a parsable slice of Poll.
type Polls []*Poll

func (po Polls) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}

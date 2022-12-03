// Code generated by ent, DO NOT EDIT.

package vote

import (
	"time"
)

const (
	// Label holds the string label denoting the vote type in the database.
	Label = "vote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// EdgePoll holds the string denoting the poll edge name in mutations.
	EdgePoll = "poll"
	// Table holds the table name of the vote in the database.
	Table = "votes"
	// TemplateTable is the table that holds the template relation/edge.
	TemplateTable = "votes"
	// TemplateInverseTable is the table name for the VoteTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "votetemplate" package.
	TemplateInverseTable = "vote_templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "vote_template"
	// PollTable is the table that holds the poll relation/edge.
	PollTable = "votes"
	// PollInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollInverseTable = "polls"
	// PollColumn is the table column denoting the poll relation/edge.
	PollColumn = "poll_votes"
)

// Columns holds all SQL columns for vote fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsDeleted,
	FieldDeletedAt,
	FieldNonce,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "votes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"poll_votes",
	"vote_template",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)
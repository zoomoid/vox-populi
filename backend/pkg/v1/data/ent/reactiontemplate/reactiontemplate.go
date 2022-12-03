// Code generated by ent, DO NOT EDIT.

package reactiontemplate

import (
	"time"
)

const (
	// Label holds the string label denoting the reactiontemplate type in the database.
	Label = "reaction_template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldReaction holds the string denoting the reaction field in the database.
	FieldReaction = "reaction"
	// EdgePoll holds the string denoting the poll edge name in mutations.
	EdgePoll = "poll"
	// Table holds the table name of the reactiontemplate in the database.
	Table = "reaction_templates"
	// PollTable is the table that holds the poll relation/edge.
	PollTable = "reaction_templates"
	// PollInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollInverseTable = "polls"
	// PollColumn is the table column denoting the poll relation/edge.
	PollColumn = "poll_reaction_templates"
)

// Columns holds all SQL columns for reactiontemplate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldReaction,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reaction_templates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"poll_reaction_templates",
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
	// ReactionValidator is a validator for the "reaction" field. It is called by the builders before save.
	ReactionValidator func(string) error
)

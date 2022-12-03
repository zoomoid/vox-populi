// Code generated by ent, DO NOT EDIT.

package poll

import (
	"time"
)

const (
	// Label holds the string label denoting the poll type in the database.
	Label = "poll"
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
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsLive holds the string denoting the is_live field in the database.
	FieldIsLive = "is_live"
	// FieldCanVote holds the string denoting the can_vote field in the database.
	FieldCanVote = "can_vote"
	// FieldCanReact holds the string denoting the can_react field in the database.
	FieldCanReact = "can_react"
	// FieldCanSeeResults holds the string denoting the can_see_results field in the database.
	FieldCanSeeResults = "can_see_results"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldUnpublishedAt holds the string denoting the unpublished_at field in the database.
	FieldUnpublishedAt = "unpublished_at"
	// EdgeVoteTemplates holds the string denoting the vote_templates edge name in mutations.
	EdgeVoteTemplates = "vote_templates"
	// EdgeVotes holds the string denoting the votes edge name in mutations.
	EdgeVotes = "votes"
	// EdgeReactionTemplates holds the string denoting the reaction_templates edge name in mutations.
	EdgeReactionTemplates = "reaction_templates"
	// EdgeReactions holds the string denoting the reactions edge name in mutations.
	EdgeReactions = "reactions"
	// Table holds the table name of the poll in the database.
	Table = "polls"
	// VoteTemplatesTable is the table that holds the vote_templates relation/edge.
	VoteTemplatesTable = "vote_templates"
	// VoteTemplatesInverseTable is the table name for the VoteTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "votetemplate" package.
	VoteTemplatesInverseTable = "vote_templates"
	// VoteTemplatesColumn is the table column denoting the vote_templates relation/edge.
	VoteTemplatesColumn = "poll_vote_templates"
	// VotesTable is the table that holds the votes relation/edge.
	VotesTable = "votes"
	// VotesInverseTable is the table name for the Vote entity.
	// It exists in this package in order to avoid circular dependency with the "vote" package.
	VotesInverseTable = "votes"
	// VotesColumn is the table column denoting the votes relation/edge.
	VotesColumn = "poll_votes"
	// ReactionTemplatesTable is the table that holds the reaction_templates relation/edge.
	ReactionTemplatesTable = "reaction_templates"
	// ReactionTemplatesInverseTable is the table name for the ReactionTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "reactiontemplate" package.
	ReactionTemplatesInverseTable = "reaction_templates"
	// ReactionTemplatesColumn is the table column denoting the reaction_templates relation/edge.
	ReactionTemplatesColumn = "poll_reaction_templates"
	// ReactionsTable is the table that holds the reactions relation/edge.
	ReactionsTable = "reactions"
	// ReactionsInverseTable is the table name for the Reaction entity.
	// It exists in this package in order to avoid circular dependency with the "reaction" package.
	ReactionsInverseTable = "reactions"
	// ReactionsColumn is the table column denoting the reactions relation/edge.
	ReactionsColumn = "poll_reactions"
)

// Columns holds all SQL columns for poll fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsDeleted,
	FieldDeletedAt,
	FieldTitle,
	FieldDescription,
	FieldIsLive,
	FieldCanVote,
	FieldCanReact,
	FieldCanSeeResults,
	FieldPublishedAt,
	FieldUnpublishedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultIsLive holds the default value on creation for the "is_live" field.
	DefaultIsLive bool
	// DefaultCanVote holds the default value on creation for the "can_vote" field.
	DefaultCanVote bool
	// DefaultCanReact holds the default value on creation for the "can_react" field.
	DefaultCanReact bool
	// DefaultCanSeeResults holds the default value on creation for the "can_see_results" field.
	DefaultCanSeeResults bool
)

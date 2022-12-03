package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description"),
		field.Bool("is_live").Default(false),
		field.Bool("can_vote").Default(true),
		field.Bool("can_react").Default(true),
		field.Bool("can_see_results").Default(true),
		field.Time("published_at").Nillable().Default(nil),
		field.Time("unpublished_at").Nillable().Default(nil),
	}
}

func (Poll) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
		SoftDeleteMixin{},
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vote_templates", VoteTemplate.Type),
		edge.To("votes", Vote.Type),
		edge.To("reaction_templates", ReactionTemplate.Type),
		edge.To("reactions", Reaction.Type),
	}
}

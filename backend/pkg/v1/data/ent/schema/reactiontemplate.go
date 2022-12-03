package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReactionTemplate holds the schema definition for the ReactionTemplate entity.
type ReactionTemplate struct {
	ent.Schema
}

// Fields of the ReactionTemplate.
func (ReactionTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("reaction").NotEmpty(),
	}
}

func (ReactionTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
	}
}

// Edges of the ReactionTemplate.
func (ReactionTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).Ref("reaction_templates").Unique(),
	}
}

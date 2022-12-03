package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VoteTemplate holds the schema definition for the VoteTemplate entity.
type VoteTemplate struct {
	ent.Schema
}

// Fields of the VoteTemplate.
func (VoteTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("answer").NotEmpty(),
	}
}

func (VoteTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
	}
}

// Edges of the VoteTemplate.
func (VoteTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).Ref("vote_templates").Unique(),
	}
}

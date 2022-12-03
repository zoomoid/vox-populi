package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Reaction holds the schema definition for the Reaction entity.
type Reaction struct {
	ent.Schema
}

// Fields of the Reaction.
func (Reaction) Fields() []ent.Field {
	return []ent.Field{}
}

func (Reaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
	}
}

// Edges of the Reaction.
func (Reaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("template", ReactionTemplate.Type).Unique().Required(),
		edge.From("poll", Poll.Type).Ref("reactions").Unique(),
	}
}

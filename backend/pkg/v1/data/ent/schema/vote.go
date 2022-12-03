package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.String("nonce"),
	}
}

func (Vote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MetadataMixin{},
		SoftDeleteMixin{},
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("template", VoteTemplate.Type).Unique().Required(),
		edge.From("poll", Poll.Type).Ref("votes").Unique(),
	}
}

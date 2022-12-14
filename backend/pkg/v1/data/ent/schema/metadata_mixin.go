package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type MetadataMixin struct {
	mixin.Schema
}

func (MetadataMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		// field.Int("revision").
		// 	Positive().
		// 	Default(0).
		// 	UpdateDefault(),
	}
}

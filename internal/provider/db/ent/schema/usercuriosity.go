package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCuriosity holds the schema definition for the UserCuriosity entity.
type UserCuriosity struct {
	ent.Schema
}

// Fields of the UserCuriosity.
func (UserCuriosity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Bool("viewed").Default(false),
		field.Bool("liked").Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the UserCuriosity.
func (UserCuriosity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			StorageKey(edge.Column("user_id")),
		edge.To("curiosity", Curiosity.Type).
			Unique().
			StorageKey(edge.Column("curiosity_id")),
	}
}

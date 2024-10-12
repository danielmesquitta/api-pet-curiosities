package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// View holds the schema definition for the View entity.
type View struct {
	ent.Schema
}

// Fields of the View.
func (View) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Time("createdAt").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the View.
func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			StorageKey(edge.Column("user_id")),
		edge.To("curiosity", Curiosity.Type).
			Unique().
			StorageKey(edge.Column("curiosity_id")),
	}
}

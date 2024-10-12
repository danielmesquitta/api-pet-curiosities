package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Curiosity holds the schema definition for the Curiosity entity.
type Curiosity struct {
	ent.Schema
}

// Fields of the Curiosity.
func (Curiosity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
		field.Time("createdAt").
			Default(time.Now).
			Immutable(),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Curiosity.
func (Curiosity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pet", Pet.Type).Unique().StorageKey(edge.Column("pet_id")),
		edge.From("likes", Like.Type).Ref("curiosity"),
		edge.From("views", View.Type).Ref("curiosity"),
	}
}

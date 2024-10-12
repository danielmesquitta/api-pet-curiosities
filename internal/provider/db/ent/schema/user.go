package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.Enum("tier").
			Values("PRO", "FREE").
			Default("FREE"),
		field.Time("subscriptionExpiresAt").
			Nillable().
			Optional(),
		field.Time("createdAt").
			Default(time.Now).
			Immutable(),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
		edge.From("likes", Like.Type).Ref("user"),
		edge.From("views", View.Type).Ref("user"),
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
		Immutable().
		Unique(),
		field.String("tweet"),
		field.Time("created_at").
		Default(time.Now),
		field.String("user_id").
		Optional(),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author",User.Type).
		Ref("tweets").
		Unique().
		Field("user_id"),
    }
}

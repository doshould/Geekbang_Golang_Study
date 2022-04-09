package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.String("geners"),
	}
}

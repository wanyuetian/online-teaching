package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseSection holds the schema definition for the CourseSection entity.
type CourseSection struct {
	ent.Schema
}

// Fields of the CourseSection.
func (CourseSection) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("type").Default("video"),
		field.String("video"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at"),
	}
}

// Edges of the CourseSection.
func (CourseSection) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseTeacher holds the schema definition for the CourseTeacher entity.
type CourseTeacher struct {
	ent.Schema
}

// Fields of the CourseTeacher.
func (CourseTeacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("desc"),
		field.String("avatar"),
		field.String("title"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at"),
	}
}

// Edges of the CourseTeacher.
func (CourseTeacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("course", Course.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseChapter holds the schema definition for the CourseChapter entity.
type CourseChapter struct {
	ent.Schema
}

// Fields of the CourseChapter.
func (CourseChapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Comment("章标题"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CourseChapter.
func (CourseChapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("course_section", CourseSection.Type),
	}
}

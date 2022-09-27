package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseInfo holds the schema definition for the CourseInfo entity.
type CourseInfo struct {
	ent.Schema
}

// Fields of the CourseInfo.
func (CourseInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("click_count"),
		field.Int("learn_count"),
		field.Int("total_duration"),
		field.Int("section_count"),
		field.Float("price"),
		field.String("detail").Default(""),
		field.String("state").Default("on"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at"),
	}
}

// Edges of the CourseInfo.
func (CourseInfo) Edges() []ent.Edge {
	return nil
}

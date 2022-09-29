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
		field.String("title").NotEmpty().Comment("小节标题"),
		field.String("type").Default("video").Comment("小节类型"),
		field.String("video").Comment("小节视频地址"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CourseSection.
func (CourseSection) Edges() []ent.Edge {
	return nil
}

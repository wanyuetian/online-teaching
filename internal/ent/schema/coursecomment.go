package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseComment holds the schema definition for the CourseComment entity.
type CourseComment struct {
	ent.Schema
}

// Fields of the CourseComment.
func (CourseComment) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("用户名"),
		field.String("comment").Comment("评论内容"),
		field.Int("order").Default(0).Comment("序号"),
		field.Bool("is_show").Default(false).Comment("是否展示"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CourseComment.
func (CourseComment) Edges() []ent.Edge {
	return nil
}

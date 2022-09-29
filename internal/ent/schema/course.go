package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Comment("课程描述"),
		field.String("desc").Comment("课程描述"),
		field.String("image").Comment("课程图片"),
		field.String("tags").Comment("课程标签"),
		field.String("classification").Comment("课程分类"),
		field.Bool("is_deleted").Default(false).Comment("是否已删除"),
		field.Time("created_at").Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teachers", Teacher.Type).Ref("courses"),
		edge.To("infos", CourseInfo.Type),
		edge.To("chapters", CourseChapter.Type),
		edge.To("sections", CourseSection.Type),
		edge.To("swipers", CourseSwiper.Type),
		edge.From("users", User.Type).Ref("courses"),
	}
}

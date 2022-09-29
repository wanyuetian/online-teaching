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
		field.Int("click_count").Default(0).Comment("课程点击量"),
		field.Int("learn_count").Default(0).Comment("课程学习人数"),
		field.Int("total_duration").Default(0).Comment("课程总时长"),
		field.Int("section_count").Default(0).Comment("课程小节数"),
		field.Float("price").Comment("课程价格"),
		field.String("detail").Default("").MaxLen(8192).Comment("课程详细描述"),
		field.Int("state").Default(0).Comment("课程状态"), // 默认不展示
		field.Int("order").Default(0).Comment("序号"),
		field.Bool("is_quality").Default(false).Comment("是否精品"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CourseInfo.
func (CourseInfo) Edges() []ent.Edge {
	return nil
}

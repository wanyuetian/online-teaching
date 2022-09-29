package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CourseSwiper holds the schema definition for the CourseSwiper entity.
type CourseSwiper struct {
	ent.Schema
}

// Fields of the CourseSwiper.
func (CourseSwiper) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("轮播图名称"),
		field.Int("order").Unique().Comment("顺序"),
		field.String("image").Comment("轮播图地址"),
		field.Bool("is_show").Comment("是否展示"),
		field.Bool("is_deleted").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CourseSwiper.
func (CourseSwiper) Edges() []ent.Edge {
	return nil
}

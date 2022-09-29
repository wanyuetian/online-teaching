// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"online-teaching/internal/ent/courseswiper"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CourseSwiper is the model entity for the CourseSwiper schema.
type CourseSwiper struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 轮播图名称
	Name string `json:"name,omitempty"`
	// 顺序
	Order int `json:"order,omitempty"`
	// 轮播图地址
	Image string `json:"image,omitempty"`
	// 是否展示
	IsShow bool `json:"is_show,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	course_swipers *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CourseSwiper) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case courseswiper.FieldIsShow, courseswiper.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case courseswiper.FieldID, courseswiper.FieldOrder:
			values[i] = new(sql.NullInt64)
		case courseswiper.FieldName, courseswiper.FieldImage:
			values[i] = new(sql.NullString)
		case courseswiper.FieldCreatedAt, courseswiper.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case courseswiper.ForeignKeys[0]: // course_swipers
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CourseSwiper", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CourseSwiper fields.
func (cs *CourseSwiper) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case courseswiper.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cs.ID = int(value.Int64)
		case courseswiper.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cs.Name = value.String
			}
		case courseswiper.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				cs.Order = int(value.Int64)
			}
		case courseswiper.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				cs.Image = value.String
			}
		case courseswiper.FieldIsShow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				cs.IsShow = value.Bool
			}
		case courseswiper.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				cs.IsDeleted = value.Bool
			}
		case courseswiper.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cs.CreatedAt = value.Time
			}
		case courseswiper.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cs.UpdatedAt = value.Time
			}
		case courseswiper.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field course_swipers", value)
			} else if value.Valid {
				cs.course_swipers = new(int)
				*cs.course_swipers = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CourseSwiper.
// Note that you need to call CourseSwiper.Unwrap() before calling this method if this CourseSwiper
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CourseSwiper) Update() *CourseSwiperUpdateOne {
	return (&CourseSwiperClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the CourseSwiper entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *CourseSwiper) Unwrap() *CourseSwiper {
	_tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CourseSwiper is not a transactional entity")
	}
	cs.config.driver = _tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CourseSwiper) String() string {
	var builder strings.Builder
	builder.WriteString("CourseSwiper(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cs.ID))
	builder.WriteString("name=")
	builder.WriteString(cs.Name)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", cs.Order))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(cs.Image)
	builder.WriteString(", ")
	builder.WriteString("is_show=")
	builder.WriteString(fmt.Sprintf("%v", cs.IsShow))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", cs.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(cs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cs.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CourseSwipers is a parsable slice of CourseSwiper.
type CourseSwipers []*CourseSwiper

func (cs CourseSwipers) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
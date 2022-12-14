// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"online-teaching/internal/ent/course"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Course is the model entity for the Course schema.
type Course struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 课程描述
	Name string `json:"name,omitempty"`
	// 课程描述
	Desc string `json:"desc,omitempty"`
	// 课程图片
	Image string `json:"image,omitempty"`
	// 课程标签
	Tags string `json:"tags,omitempty"`
	// 课程分类
	Classification string `json:"classification,omitempty"`
	// 是否已删除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseQuery when eager-loading is set.
	Edges CourseEdges `json:"edges"`
}

// CourseEdges holds the relations/edges for other nodes in the graph.
type CourseEdges struct {
	// Teachers holds the value of the teachers edge.
	Teachers []*Teacher `json:"teachers,omitempty"`
	// Infos holds the value of the infos edge.
	Infos []*CourseInfo `json:"infos,omitempty"`
	// Chapters holds the value of the chapters edge.
	Chapters []*CourseChapter `json:"chapters,omitempty"`
	// Sections holds the value of the sections edge.
	Sections []*CourseSection `json:"sections,omitempty"`
	// Swipers holds the value of the swipers edge.
	Swipers []*CourseSwiper `json:"swipers,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) TeachersOrErr() ([]*Teacher, error) {
	if e.loadedTypes[0] {
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// InfosOrErr returns the Infos value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) InfosOrErr() ([]*CourseInfo, error) {
	if e.loadedTypes[1] {
		return e.Infos, nil
	}
	return nil, &NotLoadedError{edge: "infos"}
}

// ChaptersOrErr returns the Chapters value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) ChaptersOrErr() ([]*CourseChapter, error) {
	if e.loadedTypes[2] {
		return e.Chapters, nil
	}
	return nil, &NotLoadedError{edge: "chapters"}
}

// SectionsOrErr returns the Sections value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) SectionsOrErr() ([]*CourseSection, error) {
	if e.loadedTypes[3] {
		return e.Sections, nil
	}
	return nil, &NotLoadedError{edge: "sections"}
}

// SwipersOrErr returns the Swipers value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) SwipersOrErr() ([]*CourseSwiper, error) {
	if e.loadedTypes[4] {
		return e.Swipers, nil
	}
	return nil, &NotLoadedError{edge: "swipers"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Course) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case course.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case course.FieldID:
			values[i] = new(sql.NullInt64)
		case course.FieldName, course.FieldDesc, course.FieldImage, course.FieldTags, course.FieldClassification:
			values[i] = new(sql.NullString)
		case course.FieldCreatedAt, course.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Course", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Course fields.
func (c *Course) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case course.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case course.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case course.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				c.Desc = value.String
			}
		case course.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				c.Image = value.String
			}
		case course.FieldTags:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value.Valid {
				c.Tags = value.String
			}
		case course.FieldClassification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field classification", values[i])
			} else if value.Valid {
				c.Classification = value.String
			}
		case course.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				c.IsDeleted = value.Bool
			}
		case course.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case course.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTeachers queries the "teachers" edge of the Course entity.
func (c *Course) QueryTeachers() *TeacherQuery {
	return (&CourseClient{config: c.config}).QueryTeachers(c)
}

// QueryInfos queries the "infos" edge of the Course entity.
func (c *Course) QueryInfos() *CourseInfoQuery {
	return (&CourseClient{config: c.config}).QueryInfos(c)
}

// QueryChapters queries the "chapters" edge of the Course entity.
func (c *Course) QueryChapters() *CourseChapterQuery {
	return (&CourseClient{config: c.config}).QueryChapters(c)
}

// QuerySections queries the "sections" edge of the Course entity.
func (c *Course) QuerySections() *CourseSectionQuery {
	return (&CourseClient{config: c.config}).QuerySections(c)
}

// QuerySwipers queries the "swipers" edge of the Course entity.
func (c *Course) QuerySwipers() *CourseSwiperQuery {
	return (&CourseClient{config: c.config}).QuerySwipers(c)
}

// QueryUsers queries the "users" edge of the Course entity.
func (c *Course) QueryUsers() *UserQuery {
	return (&CourseClient{config: c.config}).QueryUsers(c)
}

// Update returns a builder for updating this Course.
// Note that you need to call Course.Unwrap() before calling this method if this Course
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Course) Update() *CourseUpdateOne {
	return (&CourseClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Course entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Course) Unwrap() *Course {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Course is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Course) String() string {
	var builder strings.Builder
	builder.WriteString("Course(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(c.Desc)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(c.Image)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(c.Tags)
	builder.WriteString(", ")
	builder.WriteString("classification=")
	builder.WriteString(c.Classification)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Courses is a parsable slice of Course.
type Courses []*Course

func (c Courses) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

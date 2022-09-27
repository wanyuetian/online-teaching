// Code generated by ent, DO NOT EDIT.

package course

import (
	"time"
)

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldBackgroundImage holds the string denoting the background_image field in the database.
	FieldBackgroundImage = "background_image"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeCourseTeacher holds the string denoting the course_teacher edge name in mutations.
	EdgeCourseTeacher = "course_teacher"
	// EdgeCourseInfo holds the string denoting the course_info edge name in mutations.
	EdgeCourseInfo = "course_info"
	// EdgeCourseChapter holds the string denoting the course_chapter edge name in mutations.
	EdgeCourseChapter = "course_chapter"
	// EdgeCourseSection holds the string denoting the course_section edge name in mutations.
	EdgeCourseSection = "course_section"
	// Table holds the table name of the course in the database.
	Table = "courses"
	// CourseTeacherTable is the table that holds the course_teacher relation/edge.
	CourseTeacherTable = "course_teachers"
	// CourseTeacherInverseTable is the table name for the CourseTeacher entity.
	// It exists in this package in order to avoid circular dependency with the "courseteacher" package.
	CourseTeacherInverseTable = "course_teachers"
	// CourseTeacherColumn is the table column denoting the course_teacher relation/edge.
	CourseTeacherColumn = "course_course_teacher"
	// CourseInfoTable is the table that holds the course_info relation/edge.
	CourseInfoTable = "course_infos"
	// CourseInfoInverseTable is the table name for the CourseInfo entity.
	// It exists in this package in order to avoid circular dependency with the "courseinfo" package.
	CourseInfoInverseTable = "course_infos"
	// CourseInfoColumn is the table column denoting the course_info relation/edge.
	CourseInfoColumn = "course_course_info"
	// CourseChapterTable is the table that holds the course_chapter relation/edge.
	CourseChapterTable = "course_chapters"
	// CourseChapterInverseTable is the table name for the CourseChapter entity.
	// It exists in this package in order to avoid circular dependency with the "coursechapter" package.
	CourseChapterInverseTable = "course_chapters"
	// CourseChapterColumn is the table column denoting the course_chapter relation/edge.
	CourseChapterColumn = "course_course_chapter"
	// CourseSectionTable is the table that holds the course_section relation/edge.
	CourseSectionTable = "course_sections"
	// CourseSectionInverseTable is the table name for the CourseSection entity.
	// It exists in this package in order to avoid circular dependency with the "coursesection" package.
	CourseSectionInverseTable = "course_sections"
	// CourseSectionColumn is the table column denoting the course_section relation/edge.
	CourseSectionColumn = "course_course_section"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDesc,
	FieldBackgroundImage,
	FieldIsDeleted,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "courses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"course_teacher_course",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
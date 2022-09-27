// Code generated by ent, DO NOT EDIT.

package courseteacher

import (
	"time"
)

const (
	// Label holds the string label denoting the courseteacher type in the database.
	Label = "course_teacher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// Table holds the table name of the courseteacher in the database.
	Table = "course_teachers"
	// CourseTable is the table that holds the course relation/edge.
	CourseTable = "courses"
	// CourseInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CourseInverseTable = "courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_teacher_course"
)

// Columns holds all SQL columns for courseteacher fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDesc,
	FieldAvatar,
	FieldTitle,
	FieldIsDeleted,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "course_teachers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"course_course_teacher",
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
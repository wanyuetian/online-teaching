// Code generated by ent, DO NOT EDIT.

package courseinfo

import (
	"time"
)

const (
	// Label holds the string label denoting the courseinfo type in the database.
	Label = "course_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldClickCount holds the string denoting the click_count field in the database.
	FieldClickCount = "click_count"
	// FieldLearnCount holds the string denoting the learn_count field in the database.
	FieldLearnCount = "learn_count"
	// FieldTotalDuration holds the string denoting the total_duration field in the database.
	FieldTotalDuration = "total_duration"
	// FieldSectionCount holds the string denoting the section_count field in the database.
	FieldSectionCount = "section_count"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the courseinfo in the database.
	Table = "course_infos"
)

// Columns holds all SQL columns for courseinfo fields.
var Columns = []string{
	FieldID,
	FieldClickCount,
	FieldLearnCount,
	FieldTotalDuration,
	FieldSectionCount,
	FieldPrice,
	FieldDetail,
	FieldState,
	FieldIsDeleted,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "course_infos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"course_course_info",
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
	// DefaultDetail holds the default value on creation for the "detail" field.
	DefaultDetail string
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

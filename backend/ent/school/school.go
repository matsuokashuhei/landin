// Code generated by entc, DO NOT EDIT.

package school

import (
	"time"
)

const (
	// Label holds the string label denoting the school type in the database.
	Label = "school"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeStudios holds the string denoting the studios edge name in mutations.
	EdgeStudios = "studios"
	// Table holds the table name of the school in the database.
	Table = "schools"
	// StudiosTable is the table that holds the studios relation/edge.
	StudiosTable = "studios"
	// StudiosInverseTable is the table name for the Studio entity.
	// It exists in this package in order to avoid circular dependency with the "studio" package.
	StudiosInverseTable = "studios"
	// StudiosColumn is the table column denoting the studios relation/edge.
	StudiosColumn = "school_studios"
)

// Columns holds all SQL columns for school fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
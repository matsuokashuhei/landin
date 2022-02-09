// Code generated by entc, DO NOT EDIT.

package schedule

import (
	"time"
)

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDayOfWeek holds the string denoting the day_of_week field in the database.
	FieldDayOfWeek = "day_of_week"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// Table holds the table name of the schedule in the database.
	Table = "schedules"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "schedules"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_schedules"
	// ClassesTable is the table that holds the classes relation/edge.
	ClassesTable = "classes"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "classes"
	// ClassesColumn is the table column denoting the classes relation/edge.
	ClassesColumn = "schedule_classes"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDayOfWeek,
	FieldStartTime,
	FieldEndTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "schedules"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"room_schedules",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
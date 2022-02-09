// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "instructor_classes", Type: field.TypeInt, Nullable: true},
		{Name: "schedule_classes", Type: field.TypeInt, Nullable: true},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_instructors_classes",
				Columns:    []*schema.Column{ClassesColumns[5]},
				RefColumns: []*schema.Column{InstructorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_schedules_classes",
				Columns:    []*schema.Column{ClassesColumns[6]},
				RefColumns: []*schema.Column{SchedulesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InstructorsColumns holds the columns for the "instructors" table.
	InstructorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "syllabic_characters", Type: field.TypeString},
		{Name: "biography", Type: field.TypeString, Nullable: true},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
	}
	// InstructorsTable holds the schema information for the "instructors" table.
	InstructorsTable = &schema.Table{
		Name:       "instructors",
		Columns:    InstructorsColumns,
		PrimaryKey: []*schema.Column{InstructorsColumns[0]},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeInt},
		{Name: "studio_rooms", Type: field.TypeInt, Nullable: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rooms_studios_rooms",
				Columns:    []*schema.Column{RoomsColumns[5]},
				RefColumns: []*schema.Column{StudiosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SchedulesColumns holds the columns for the "schedules" table.
	SchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "day_of_week", Type: field.TypeInt},
		{Name: "start_time", Type: field.TypeString},
		{Name: "end_time", Type: field.TypeString},
		{Name: "room_schedules", Type: field.TypeInt, Nullable: true},
	}
	// SchedulesTable holds the schema information for the "schedules" table.
	SchedulesTable = &schema.Table{
		Name:       "schedules",
		Columns:    SchedulesColumns,
		PrimaryKey: []*schema.Column{SchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "schedules_rooms_schedules",
				Columns:    []*schema.Column{SchedulesColumns[6]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SchoolsColumns holds the columns for the "schools" table.
	SchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// SchoolsTable holds the schema information for the "schools" table.
	SchoolsTable = &schema.Table{
		Name:       "schools",
		Columns:    SchoolsColumns,
		PrimaryKey: []*schema.Column{SchoolsColumns[0]},
	}
	// StudiosColumns holds the columns for the "studios" table.
	StudiosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "location", Type: field.TypeString},
		{Name: "school_studios", Type: field.TypeInt, Nullable: true},
	}
	// StudiosTable holds the schema information for the "studios" table.
	StudiosTable = &schema.Table{
		Name:       "studios",
		Columns:    StudiosColumns,
		PrimaryKey: []*schema.Column{StudiosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "studios_schools_studios",
				Columns:    []*schema.Column{StudiosColumns[5]},
				RefColumns: []*schema.Column{SchoolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "firebase_uid", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassesTable,
		InstructorsTable,
		RoomsTable,
		SchedulesTable,
		SchoolsTable,
		StudiosTable,
		UsersTable,
	}
)

func init() {
	ClassesTable.ForeignKeys[0].RefTable = InstructorsTable
	ClassesTable.ForeignKeys[1].RefTable = SchedulesTable
	RoomsTable.ForeignKeys[0].RefTable = StudiosTable
	SchedulesTable.ForeignKeys[0].RefTable = RoomsTable
	StudiosTable.ForeignKeys[0].RefTable = SchoolsTable
}
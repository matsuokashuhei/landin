// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/matsuokashuhei/landin/ent/class"
	"github.com/matsuokashuhei/landin/ent/instructor"
	"github.com/matsuokashuhei/landin/ent/schedule"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Level holds the value of the "level" field.
	Level string `json:"level,omitempty"`
	// Tuition holds the value of the "tuition" field.
	Tuition int `json:"tuition,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges              ClassEdges `json:"edges"`
	instructor_classes *int
	schedule_classes   *int
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule `json:"schedule,omitempty"`
	// Instructor holds the value of the instructor edge.
	Instructor *Instructor `json:"instructor,omitempty"`
	// MembersClasses holds the value of the members_classes edge.
	MembersClasses []*MembersClass `json:"members_classes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[0] {
		if e.Schedule == nil {
			// The edge schedule was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// InstructorOrErr returns the Instructor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) InstructorOrErr() (*Instructor, error) {
	if e.loadedTypes[1] {
		if e.Instructor == nil {
			// The edge instructor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instructor.Label}
		}
		return e.Instructor, nil
	}
	return nil, &NotLoadedError{edge: "instructor"}
}

// MembersClassesOrErr returns the MembersClasses value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) MembersClassesOrErr() ([]*MembersClass, error) {
	if e.loadedTypes[2] {
		return e.MembersClasses, nil
	}
	return nil, &NotLoadedError{edge: "members_classes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldID, class.FieldTuition:
			values[i] = new(sql.NullInt64)
		case class.FieldName, class.FieldLevel:
			values[i] = new(sql.NullString)
		case class.FieldCreateTime, class.FieldUpdateTime, class.FieldStartDate, class.FieldEndDate:
			values[i] = new(sql.NullTime)
		case class.ForeignKeys[0]: // instructor_classes
			values[i] = new(sql.NullInt64)
		case class.ForeignKeys[1]: // schedule_classes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Class", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case class.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case class.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case class.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case class.FieldLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				c.Level = value.String
			}
		case class.FieldTuition:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tuition", values[i])
			} else if value.Valid {
				c.Tuition = int(value.Int64)
			}
		case class.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				c.StartDate = value.Time
			}
		case class.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				c.EndDate = value.Time
			}
		case class.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field instructor_classes", value)
			} else if value.Valid {
				c.instructor_classes = new(int)
				*c.instructor_classes = int(value.Int64)
			}
		case class.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field schedule_classes", value)
			} else if value.Valid {
				c.schedule_classes = new(int)
				*c.schedule_classes = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySchedule queries the "schedule" edge of the Class entity.
func (c *Class) QuerySchedule() *ScheduleQuery {
	return (&ClassClient{config: c.config}).QuerySchedule(c)
}

// QueryInstructor queries the "instructor" edge of the Class entity.
func (c *Class) QueryInstructor() *InstructorQuery {
	return (&ClassClient{config: c.config}).QueryInstructor(c)
}

// QueryMembersClasses queries the "members_classes" edge of the Class entity.
func (c *Class) QueryMembersClasses() *MembersClassQuery {
	return (&ClassClient{config: c.config}).QueryMembersClasses(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return (&ClassClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", level=")
	builder.WriteString(c.Level)
	builder.WriteString(", tuition=")
	builder.WriteString(fmt.Sprintf("%v", c.Tuition))
	builder.WriteString(", start_date=")
	builder.WriteString(c.StartDate.Format(time.ANSIC))
	builder.WriteString(", end_date=")
	builder.WriteString(c.EndDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class

func (c Classes) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

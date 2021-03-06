// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/matsuokashuhei/landin/ent/school"
	"github.com/matsuokashuhei/landin/ent/studio"
)

// Studio is the model entity for the Studio schema.
type Studio struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudioQuery when eager-loading is set.
	Edges          StudioEdges `json:"edges"`
	school_studios *int
}

// StudioEdges holds the relations/edges for other nodes in the graph.
type StudioEdges struct {
	// School holds the value of the school edge.
	School *School `json:"school,omitempty"`
	// Rooms holds the value of the rooms edge.
	Rooms []*Room `json:"rooms,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SchoolOrErr returns the School value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudioEdges) SchoolOrErr() (*School, error) {
	if e.loadedTypes[0] {
		if e.School == nil {
			// The edge school was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: school.Label}
		}
		return e.School, nil
	}
	return nil, &NotLoadedError{edge: "school"}
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading.
func (e StudioEdges) RoomsOrErr() ([]*Room, error) {
	if e.loadedTypes[1] {
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Studio) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studio.FieldID:
			values[i] = new(sql.NullInt64)
		case studio.FieldName, studio.FieldLocation:
			values[i] = new(sql.NullString)
		case studio.FieldCreateTime, studio.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case studio.ForeignKeys[0]: // school_studios
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Studio", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Studio fields.
func (s *Studio) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studio.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case studio.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case studio.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case studio.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case studio.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				s.Location = value.String
			}
		case studio.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field school_studios", value)
			} else if value.Valid {
				s.school_studios = new(int)
				*s.school_studios = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySchool queries the "school" edge of the Studio entity.
func (s *Studio) QuerySchool() *SchoolQuery {
	return (&StudioClient{config: s.config}).QuerySchool(s)
}

// QueryRooms queries the "rooms" edge of the Studio entity.
func (s *Studio) QueryRooms() *RoomQuery {
	return (&StudioClient{config: s.config}).QueryRooms(s)
}

// Update returns a builder for updating this Studio.
// Note that you need to call Studio.Unwrap() before calling this method if this Studio
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Studio) Update() *StudioUpdateOne {
	return (&StudioClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Studio entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Studio) Unwrap() *Studio {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Studio is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Studio) String() string {
	var builder strings.Builder
	builder.WriteString("Studio(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", location=")
	builder.WriteString(s.Location)
	builder.WriteByte(')')
	return builder.String()
}

// Studios is a parsable slice of Studio.
type Studios []*Studio

func (s Studios) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}

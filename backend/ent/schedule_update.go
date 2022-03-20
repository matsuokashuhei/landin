// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/matsuokashuhei/landin/ent/class"
	"github.com/matsuokashuhei/landin/ent/predicate"
	"github.com/matsuokashuhei/landin/ent/room"
	"github.com/matsuokashuhei/landin/ent/schedule"
)

// ScheduleUpdate is the builder for updating Schedule entities.
type ScheduleUpdate struct {
	config
	hooks    []Hook
	mutation *ScheduleMutation
}

// Where appends a list predicates to the ScheduleUpdate builder.
func (su *ScheduleUpdate) Where(ps ...predicate.Schedule) *ScheduleUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *ScheduleUpdate) SetUpdateTime(t time.Time) *ScheduleUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetDayOfWeek sets the "day_of_week" field.
func (su *ScheduleUpdate) SetDayOfWeek(i int) *ScheduleUpdate {
	su.mutation.ResetDayOfWeek()
	su.mutation.SetDayOfWeek(i)
	return su
}

// AddDayOfWeek adds i to the "day_of_week" field.
func (su *ScheduleUpdate) AddDayOfWeek(i int) *ScheduleUpdate {
	su.mutation.AddDayOfWeek(i)
	return su
}

// SetStartTime sets the "start_time" field.
func (su *ScheduleUpdate) SetStartTime(s string) *ScheduleUpdate {
	su.mutation.SetStartTime(s)
	return su
}

// SetEndTime sets the "end_time" field.
func (su *ScheduleUpdate) SetEndTime(s string) *ScheduleUpdate {
	su.mutation.SetEndTime(s)
	return su
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (su *ScheduleUpdate) SetRoomID(id int) *ScheduleUpdate {
	su.mutation.SetRoomID(id)
	return su
}

// SetRoom sets the "room" edge to the Room entity.
func (su *ScheduleUpdate) SetRoom(r *Room) *ScheduleUpdate {
	return su.SetRoomID(r.ID)
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (su *ScheduleUpdate) AddClassIDs(ids ...int) *ScheduleUpdate {
	su.mutation.AddClassIDs(ids...)
	return su
}

// AddClasses adds the "classes" edges to the Class entity.
func (su *ScheduleUpdate) AddClasses(c ...*Class) *ScheduleUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddClassIDs(ids...)
}

// Mutation returns the ScheduleMutation object of the builder.
func (su *ScheduleUpdate) Mutation() *ScheduleMutation {
	return su.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (su *ScheduleUpdate) ClearRoom() *ScheduleUpdate {
	su.mutation.ClearRoom()
	return su
}

// ClearClasses clears all "classes" edges to the Class entity.
func (su *ScheduleUpdate) ClearClasses() *ScheduleUpdate {
	su.mutation.ClearClasses()
	return su
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (su *ScheduleUpdate) RemoveClassIDs(ids ...int) *ScheduleUpdate {
	su.mutation.RemoveClassIDs(ids...)
	return su
}

// RemoveClasses removes "classes" edges to Class entities.
func (su *ScheduleUpdate) RemoveClasses(c ...*Class) *ScheduleUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveClassIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScheduleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScheduleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScheduleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScheduleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ScheduleUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := schedule.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScheduleUpdate) check() error {
	if _, ok := su.mutation.RoomID(); su.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.room"`)
	}
	return nil
}

func (su *ScheduleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.DayOfWeek(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDayOfWeek,
		})
	}
	if value, ok := su.mutation.AddedDayOfWeek(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDayOfWeek,
		})
	}
	if value, ok := su.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldStartTime,
		})
	}
	if value, ok := su.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldEndTime,
		})
	}
	if su.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.RoomTable,
			Columns: []string{schedule.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.RoomTable,
			Columns: []string{schedule.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedClassesIDs(); len(nodes) > 0 && !su.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ScheduleUpdateOne is the builder for updating a single Schedule entity.
type ScheduleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScheduleMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *ScheduleUpdateOne) SetUpdateTime(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetDayOfWeek sets the "day_of_week" field.
func (suo *ScheduleUpdateOne) SetDayOfWeek(i int) *ScheduleUpdateOne {
	suo.mutation.ResetDayOfWeek()
	suo.mutation.SetDayOfWeek(i)
	return suo
}

// AddDayOfWeek adds i to the "day_of_week" field.
func (suo *ScheduleUpdateOne) AddDayOfWeek(i int) *ScheduleUpdateOne {
	suo.mutation.AddDayOfWeek(i)
	return suo
}

// SetStartTime sets the "start_time" field.
func (suo *ScheduleUpdateOne) SetStartTime(s string) *ScheduleUpdateOne {
	suo.mutation.SetStartTime(s)
	return suo
}

// SetEndTime sets the "end_time" field.
func (suo *ScheduleUpdateOne) SetEndTime(s string) *ScheduleUpdateOne {
	suo.mutation.SetEndTime(s)
	return suo
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (suo *ScheduleUpdateOne) SetRoomID(id int) *ScheduleUpdateOne {
	suo.mutation.SetRoomID(id)
	return suo
}

// SetRoom sets the "room" edge to the Room entity.
func (suo *ScheduleUpdateOne) SetRoom(r *Room) *ScheduleUpdateOne {
	return suo.SetRoomID(r.ID)
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (suo *ScheduleUpdateOne) AddClassIDs(ids ...int) *ScheduleUpdateOne {
	suo.mutation.AddClassIDs(ids...)
	return suo
}

// AddClasses adds the "classes" edges to the Class entity.
func (suo *ScheduleUpdateOne) AddClasses(c ...*Class) *ScheduleUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddClassIDs(ids...)
}

// Mutation returns the ScheduleMutation object of the builder.
func (suo *ScheduleUpdateOne) Mutation() *ScheduleMutation {
	return suo.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (suo *ScheduleUpdateOne) ClearRoom() *ScheduleUpdateOne {
	suo.mutation.ClearRoom()
	return suo
}

// ClearClasses clears all "classes" edges to the Class entity.
func (suo *ScheduleUpdateOne) ClearClasses() *ScheduleUpdateOne {
	suo.mutation.ClearClasses()
	return suo
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (suo *ScheduleUpdateOne) RemoveClassIDs(ids ...int) *ScheduleUpdateOne {
	suo.mutation.RemoveClassIDs(ids...)
	return suo
}

// RemoveClasses removes "classes" edges to Class entities.
func (suo *ScheduleUpdateOne) RemoveClasses(c ...*Class) *ScheduleUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveClassIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScheduleUpdateOne) Select(field string, fields ...string) *ScheduleUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Schedule entity.
func (suo *ScheduleUpdateOne) Save(ctx context.Context) (*Schedule, error) {
	var (
		err  error
		node *Schedule
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScheduleUpdateOne) SaveX(ctx context.Context) *Schedule {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScheduleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScheduleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ScheduleUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := schedule.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScheduleUpdateOne) check() error {
	if _, ok := suo.mutation.RoomID(); suo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Schedule.room"`)
	}
	return nil
}

func (suo *ScheduleUpdateOne) sqlSave(ctx context.Context) (_node *Schedule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Schedule.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedule.FieldID)
		for _, f := range fields {
			if !schedule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != schedule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.DayOfWeek(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDayOfWeek,
		})
	}
	if value, ok := suo.mutation.AddedDayOfWeek(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDayOfWeek,
		})
	}
	if value, ok := suo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldStartTime,
		})
	}
	if value, ok := suo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldEndTime,
		})
	}
	if suo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.RoomTable,
			Columns: []string{schedule.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.RoomTable,
			Columns: []string{schedule.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedClassesIDs(); len(nodes) > 0 && !suo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ClassesTable,
			Columns: []string{schedule.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Schedule{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

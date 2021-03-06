package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

func (Schedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("day_of_week").
			Annotations(entgql.OrderField("DAY_OF_WEEK")),
		field.String("start_time").
			Annotations(entgql.OrderField("START_TIME")),
		field.String("end_time"),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("schedules").Unique().Required(),
		edge.To("classes", Class.Type),
	}
}

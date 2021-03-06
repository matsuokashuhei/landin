package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/room"
	"github.com/matsuokashuhei/landin/ent/schedule"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.CreateRoomInput) (*ent.Room, error) {
	// TODO: Move to repository
	room, err := r.client.Room.Create().
		SetName(input.Name).
		SetCapacity(input.Capacity).
		SetStudioID(input.StudioID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, input model.UpdateRoomInput) (*ent.Room, error) {
	// TODO: Move to repository
	room, err := r.client.Room.Query().Where(room.ID(input.ID)).First(ctx)
	if err != nil {
		return nil, err
	}
	update := room.Update()
	if input.Name != nil {
		update = update.SetName(*input.Name)
	}
	if input.Capacity != nil {
		update = update.SetCapacity(*input.Capacity)
	}
	room, err = update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *mutationResolver) DeleteRoom(ctx context.Context, id int) (*ent.Room, error) {
	// TODO: Move to repository
	room, err := r.client.Room.Query().Where(room.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.client.Room.DeleteOne(room).Exec(ctx); err != nil {
		return nil, err
	}
	return room, nil
}

func (r *queryResolver) Room(ctx context.Context, id int) (*ent.Room, error) {
	// TODO: Move to repository
	room, err := r.client.Room.Query().Where(room.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*ent.Room, error) {
	// TODO: Move to repository
	rooms, err := r.client.Room.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *roomResolver) Schedules(ctx context.Context, obj *ent.Room) ([]*ent.Schedule, error) {
	return obj.
		QuerySchedules().
		Order(ent.Asc(schedule.FieldDayOfWeek)).
		Order(ent.Asc(schedule.FieldStartTime)).
		All(ctx)
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }

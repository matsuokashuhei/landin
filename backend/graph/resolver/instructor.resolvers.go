package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*models.Instructor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Instructor(ctx context.Context, id uint) (*models.Instructor, error) {
	repository := repositories.NewInstructorRepository(r.DB)
	instructor, err := repository.Find(id)
	if err != nil {
		return nil, err
	}
	return instructor, nil
}

func (r *queryResolver) Instructors(ctx context.Context, offset *uint, limit *uint) ([]*models.Instructor, error) {
	if offset == nil {
		offset = new(uint)
		*offset = 0
	}
	if limit == nil {
		limit = new(uint)
		*limit = 10
	}
	repository := repositories.NewInstructorRepository(r.DB)
	instructors, err := repository.FindAll(int(*offset), int(*limit))
	if err != nil {
		return nil, err
	}
	return instructors, nil
}

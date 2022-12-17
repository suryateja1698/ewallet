package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/suryateja1698/ewallet/graph/generated"
	"github.com/suryateja1698/ewallet/graph/model"
	"github.com/suryateja1698/ewallet/pkg/db"
	"github.com/suryateja1698/ewallet/pkg/utils"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userID := utils.UserIDGenerator(input.Email, input.UserName)

	user := model.User{
		ID:       userID,
		Name:     input.Name,
		UserName: input.UserName,
		Email:    input.Email,
	}
	err := db.AddUser(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, input string) (*model.User, error) {
	user, err := db.GetUser(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := db.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

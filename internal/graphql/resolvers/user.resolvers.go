package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/thisjustinf/togo-list/internal/graphql/models"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, registerDto *models.RegisterDto) (uuid.UUID, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, loginDto *models.LoginDto) (*models.User, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, updateUserDto *models.UpdateUserDto) (*models.ResponseStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, userID *uuid.UUID) (*models.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

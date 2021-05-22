package service_interfaces

import (
	"context"
	"pristine/models"
)

type UserService interface {
	SignInUser(ctx context.Context, user models.UserCreateOrUpdate) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	SaveUser(ctx context.Context, newOrg models.UserCreateOrUpdate) (*models.User, error)
}

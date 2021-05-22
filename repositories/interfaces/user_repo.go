package repo_interfaces

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pristine/models"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetUserByUserId(ctx context.Context, orgId string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user models.UserCreateOrUpdate) (*models.User, error)
	UpdateUser(ctx context.Context, user models.UserCreateOrUpdate) (*models.User, error)
}

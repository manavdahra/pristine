package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"pristine/models"
	repo_interfaces "pristine/repositories/interfaces"
)

type UserService struct {
	userRepo repo_interfaces.UserRepo
	userMap  map[string]models.User
}

func NewUserService(userRepo repo_interfaces.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
		userMap:  make(map[string]models.User, 0),
	}
}

func (service *UserService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if user, ok := service.userMap[email]; ok {
		return &user, nil
	}
	user, err := service.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	service.userMap[email] = *user
	return user, nil
}

func (service *UserService) SignInUser(ctx context.Context, user models.UserCreateOrUpdate) (*models.User, error) {
	existingUser, err := service.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return service.SaveUser(ctx, user)
	}
	return existingUser, nil
}

func (service *UserService) SignOutUser(ctx context.Context, email string) error {
	_, err := service.FindUserByEmail(ctx, email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}
	delete(service.userMap, email)
	return nil
}

func (service *UserService) SaveUser(ctx context.Context, newUser models.UserCreateOrUpdate) (*models.User, error) {
	if newUser.UserId != "" {
		return service.userRepo.UpdateUser(ctx, newUser)
	}
	return service.userRepo.CreateUser(ctx, newUser)
}

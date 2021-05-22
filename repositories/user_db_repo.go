package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"pristine/models"
	"pristine/providers"
)

type UserDbRepo struct {
	collectionName string
	logger         *zap.SugaredLogger
	mongo          *providers.MongoDAL
	IdGenerator    *providers.IdGenerator
}

func NewUserDbRepo(idGen *providers.IdGenerator, mongo *providers.MongoDAL, logger *zap.SugaredLogger) *OrgDbRepo {
	return &OrgDbRepo{
		collectionName: "user",
		logger:         logger,
		mongo:          mongo,
		IdGenerator:    idGen,
	}
}

func (repo *OrgDbRepo) GetUserById(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user, err = &models.User{}, error(nil)
	res := repo.mongo.Db.Collection(repo.collectionName).FindOne(ctx, bson.D{{"_id", id}})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	if err = res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *OrgDbRepo) GetUserByUserId(ctx context.Context, userId string) (*models.User, error) {
	var user, err = &models.User{}, error(nil)
	res := repo.mongo.Db.Collection(repo.collectionName).FindOne(ctx, bson.D{{"userId", userId}})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	if err = res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *OrgDbRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user, err = &models.User{}, error(nil)
	res := repo.mongo.Db.Collection(repo.collectionName).FindOne(ctx, bson.D{{"email", email}})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	if err = res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *OrgDbRepo) CreateUser(ctx context.Context, newUser models.UserCreateOrUpdate) (*models.User, error) {
	res, err := repo.mongo.Db.Collection(repo.collectionName).InsertOne(ctx, &models.User{
		UserId: repo.IdGenerator.GenerateNewId(),
		Name:   newUser.Name,
		Email:  newUser.Email,
	})
	if err != nil {
		return nil, err
	}
	return repo.GetUserById(ctx, res.InsertedID.(primitive.ObjectID))
}

func (repo *OrgDbRepo) UpdateUser(ctx context.Context, user models.UserCreateOrUpdate) (*models.User, error) {
	if user.UserId != "" {
		return nil, errors.New("user id cannot be empty for updating user")
	}
	res := repo.mongo.Db.Collection(repo.collectionName).FindOneAndUpdate(ctx, bson.D{{"userId", user.UserId}},
		bson.M{
			"name":  user.Name,
			"email": user.Email,
		})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var updatedUser *models.User
	if err := res.Decode(updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

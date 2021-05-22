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

type OrgDbRepo struct {
	collectionName string
	logger         *zap.SugaredLogger
	mongo          *providers.MongoDAL
	IdGenerator    *providers.IdGenerator
}

func NewOrgDbRepo(idGen *providers.IdGenerator, mongo *providers.MongoDAL, logger *zap.SugaredLogger) *OrgDbRepo {
	return &OrgDbRepo{
		collectionName: "organisation",
		logger:         logger,
		mongo:          mongo,
		IdGenerator:    idGen,
	}
}

func (repo *OrgDbRepo) GetOrgById(ctx context.Context, id primitive.ObjectID) (*models.Organisation, error) {
	var org, err = &models.Organisation{}, error(nil)
	res := repo.mongo.Db.Collection(repo.collectionName).FindOne(ctx, bson.D{{"_id", id}})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	if err = res.Decode(org); err != nil {
		return nil, err
	}
	return org, nil
}

func (repo *OrgDbRepo) GetOrgByOrgId(ctx context.Context, orgId string) (*models.Organisation, error) {
	var org, err = &models.Organisation{}, error(nil)
	res := repo.mongo.Db.Collection(repo.collectionName).FindOne(ctx, bson.D{{"orgId", orgId}})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	if err = res.Decode(org); err != nil {
		return nil, err
	}
	return org, nil
}

func (repo *OrgDbRepo) CreateOrg(ctx context.Context, newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	res, err := repo.mongo.Db.Collection(repo.collectionName).InsertOne(ctx, &models.Organisation{
		OrgId: repo.IdGenerator.GenerateNewId(),
		Name:  newOrg.Name,
	})
	if err != nil {
		return nil, err
	}
	return repo.GetOrgById(ctx, res.InsertedID.(primitive.ObjectID))
}

func (repo *OrgDbRepo) UpdateOrg(ctx context.Context, organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	if organisation.OrgId != "" {
		return nil, errors.New("org id cannot be empty for updating organisation")
	}
	res := repo.mongo.Db.Collection(repo.collectionName).FindOneAndUpdate(ctx, bson.D{{"orgId", organisation.OrgId}},
		bson.M{
			"name":  organisation.Name,
		})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var updatedOrg *models.Organisation
	if err := res.Decode(updatedOrg); err != nil {
		return nil, err
	}
	return updatedOrg, nil
}

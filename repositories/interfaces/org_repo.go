package repo_interfaces

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pristine/models"
)

type OrganisationRepo interface {
	GetOrgById(ctx context.Context, id primitive.ObjectID) (*models.Organisation, error)
	GetOrgByOrgId(ctx context.Context, orgId string) (*models.Organisation, error)
	CreateOrg(ctx context.Context, organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error)
	UpdateOrg(ctx context.Context, organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error)
}

package service_interfaces

import (
	"context"
	"pristine/models"
)

type OrganisationService interface {
	FindOrganisationByOrgId(ctx context.Context, orgId string) (*models.Organisation, error)
	SaveOrganisation(ctx context.Context, newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error)
}

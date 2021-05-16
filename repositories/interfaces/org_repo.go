package interfaces

import "pristine/models"

type OrganisationRepo interface {
	GetOrgById(id string) (*models.Organisation, error)
	GetOrgByEmail(email string) (*models.Organisation, error)
	CreateOrg(organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error)
	UpdateOrg(organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error)
}

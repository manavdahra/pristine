package interfaces

import "pristine/models"

type OrganisationService interface {
	FindOrganisation(orgId string) (*models.Organisation, error)
	SaveOrganisation(newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error)
}

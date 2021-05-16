package services

import (
	"pristine/models"
	"pristine/repositories/interfaces"
)

type OrgService struct {
	OrgRepo interfaces.OrganisationRepo
}

func (service *OrgService) FindOrganisation(orgId string) (*models.Organisation, error) {
	return service.OrgRepo.GetOrgById(orgId)
}

func (service *OrgService) SaveOrganisation(newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	if newOrg.Id != "" {
		return service.OrgRepo.UpdateOrg(newOrg)
	}
	return service.OrgRepo.CreateOrg(newOrg)
}
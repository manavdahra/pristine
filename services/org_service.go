package services

import (
	"context"
	"pristine/models"
	"pristine/repositories/interfaces"
)

type OrgService struct {
	OrgRepo repo_interfaces.OrganisationRepo
}

func NewOrgService(repo repo_interfaces.OrganisationRepo) *OrgService {
	return &OrgService{OrgRepo: repo}
}

func (service *OrgService) FindOrganisationByOrgId(ctx context.Context, orgId string) (*models.Organisation, error) {
	return service.OrgRepo.GetOrgByOrgId(ctx, orgId)
}

func (service *OrgService) SaveOrganisation(ctx context.Context, newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	if newOrg.OrgId != "" {
		return service.OrgRepo.UpdateOrg(ctx, newOrg)
	}
	return service.OrgRepo.CreateOrg(ctx, newOrg)
}

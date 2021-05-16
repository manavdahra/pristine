package repositories

import (
	"errors"
	"pristine/models"
	"pristine/providers"
)

type OrgDbRepo struct {
	OrganisationsMap map[string]models.Organisation
	IdGenerator      *providers.IdGenerator
}

func NewOrgDbRepo(idGen *providers.IdGenerator) *OrgDbRepo {
	return &OrgDbRepo{
		OrganisationsMap: make(map[string]models.Organisation, 0),
		IdGenerator: idGen,
	}
}

func (repo *OrgDbRepo) GetOrgById(id string) (*models.Organisation, error) {
	if org, ok := repo.OrganisationsMap[id]; ok {
		return &org, nil
	} else {
		return nil, errors.New("organisation not found")
	}
}

func (repo *OrgDbRepo) GetOrgByEmail(email string) (*models.Organisation, error) {
	for _, org := range repo.OrganisationsMap {
		if org.Email == email {
			return &org, nil
		}
	}
	return nil, errors.New("organisation not found")
}

func (repo *OrgDbRepo) CreateOrg(newOrg models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	orgId := repo.IdGenerator.GenerateNewId()
	model := models.Organisation{
		Id:    orgId,
		Name:  newOrg.Name,
		Email: newOrg.Email,
	}
	repo.OrganisationsMap[orgId] = model
	return &model, nil
}

func (repo *OrgDbRepo) UpdateOrg(organisation models.OrganisationCreateOrUpdate) (*models.Organisation, error) {
	orgId := repo.IdGenerator.GenerateNewId()
	if model, ok := repo.OrganisationsMap[orgId]; ok {
		model.Name = organisation.Name
		model.Email = organisation.Email
		repo.OrganisationsMap[orgId] = model
		return &model, nil
	} else {
		return repo.CreateOrg(models.OrganisationCreateOrUpdate{
			Name:  organisation.Name,
			Email: organisation.Email,
		})
	}
}

package models

type Organisation struct {
	OrgId string
	Name  string
}

type OrganisationCreateOrUpdate struct {
	OrgId string
	Name  string
}

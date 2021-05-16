package models

type Organisation struct {
	Id    string
	Name  string
	Email string
}

type OrganisationCreateOrUpdate struct {
	Id    string
	Name  string
	Email string
}

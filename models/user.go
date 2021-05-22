package models

type User struct {
	UserId string
	Name   string
	Email  string
}

type UserCreateOrUpdate struct {
	UserId string
	Name   string
	Email  string
}

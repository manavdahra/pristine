package response

import (
	"google.golang.org/api/idtoken"
)

type UserDetail struct {
	UserId  string             `json:"userId"`
	Name    string             `json:"name"`
	Email   string             `json:"email"`
	Payload idtoken.Payload    `json:"payload"`
}

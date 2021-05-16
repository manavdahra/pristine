package response

import (
	"google.golang.org/api/idtoken"
)

type UserDetail struct {
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Payload idtoken.Payload `json:"payload"`
}

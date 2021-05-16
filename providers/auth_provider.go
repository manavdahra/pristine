package providers

import (
	"context"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
)

type AuthProvider struct {
	AuthService *idtoken.Validator
}

func NewAuthProvider() *AuthProvider {
	authService, err := idtoken.NewValidator(context.Background(), idtoken.WithHTTPClient(http.DefaultClient))
	if err != nil {
		log.Fatal(err)
	}
	return &AuthProvider{
		AuthService: authService,
	}
}

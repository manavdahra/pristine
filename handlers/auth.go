package handlers

import (
	"encoding/json"
	"net/http"
	"pristine/dto/request"
	"pristine/dto/response"
	"pristine/providers"
	"pristine/repositories/interfaces"
)

type AuthHandler struct {
	OrgRepo      interfaces.OrganisationRepo
	AuthProvider *providers.AuthProvider
}

func NewAuthHandler(orgRepo interfaces.OrganisationRepo, authProvider *providers.AuthProvider) *AuthHandler {
	return &AuthHandler{
		OrgRepo:      orgRepo,
		AuthProvider: authProvider,
	}
}

func (handler *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//var csrfCookie *http.Cookie
	var err error
	//if csrfCookie, err = r.Cookie("g_csrf_token"); err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//fmt.Printf("g_csrf_token: %s\n", csrfCookie.Value)

	var loginReq request.LoginReq
	if err = json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot read login request body"))
		return
	}

	resp, err := handler.AuthProvider.AuthService.Validate(r.Context(), loginReq.IdToken, "214628435307-btgenci9bn7cc4qv7bt14gl5ul56te5d.apps.googleusercontent.com")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot validate user account"))
		return
	}

	userDetail := response.UserDetail{
		Payload: *resp,
		Name:    resp.Claims["name"].(string),
		Email:   resp.Claims["email"].(string),
	}
	respBytes, err := json.Marshal(userDetail)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot validate user account"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

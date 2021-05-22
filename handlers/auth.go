package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/api/idtoken"
	"net/http"
	"os"
	"pristine/dto/request"
	"pristine/dto/response"
	handler_interfaces "pristine/handlers/interfaces"
	"pristine/models"
	service_interfaces "pristine/services/interfaces"
)

type AuthHandler struct {
	logger      *zap.SugaredLogger
	UserService service_interfaces.UserService
	Validator   *idtoken.Validator
	AuthConfig  handler_interfaces.AuthConfig
}

func NewAuthHandler(userService service_interfaces.UserService, authConfig handler_interfaces.AuthConfig, logger *zap.SugaredLogger) *AuthHandler {
	validator, err := idtoken.NewValidator(context.Background(), idtoken.WithHTTPClient(http.DefaultClient))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	return &AuthHandler{
		logger:      logger,
		UserService: userService,
		Validator:   validator,
		AuthConfig:  authConfig,
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
		_, _ = w.Write([]byte("cannot read login request body"))
		return
	}

	handler.logger.Infof("tokenId: %s", loginReq.IdToken)
	ctx := r.Context()
	resp, err := handler.Validator.Validate(ctx, loginReq.IdToken, handler.AuthConfig.GetClientId())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	name := resp.Claims["name"].(string)
	email := resp.Claims["email"].(string)
	user, err := handler.UserService.SignInUser(ctx, models.UserCreateOrUpdate{
		Name:  name,
		Email: email,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	userDetail := response.UserDetail{
		UserId:  user.UserId,
		Payload: *resp,
		Name:    user.Name,
		Email:   user.Email,
	}
	respBytes, err := json.Marshal(userDetail)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("cannot validate user account"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBytes)
}

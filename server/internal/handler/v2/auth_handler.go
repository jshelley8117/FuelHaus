package v2

import (
	"log"
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/lib"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

type AuthHandler struct {
	AuthService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	var auth model.AuthRequest

	if err := lib.DecodeAndValidateRequest(r, &auth); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: err.Error()})
		return
	}

	jwt, err := ah.AuthService.AuthenticateExistingUser(ctx, r.RemoteAddr, r.UserAgent(), r.Method, auth.Email, auth.Password)
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Login Successful", Data: &jwt})
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	var user model.User

	if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: err.Error()})
		return
	}

	jwt, err := ah.AuthService.AuthenticateNewUser(ctx, r.RemoteAddr, r.UserAgent(), r.Method, user)
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{
		Message: "Registration Successful",
		Data: map[string]any{
			"token": &jwt,
		}})
}

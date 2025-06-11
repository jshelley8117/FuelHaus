package v1

import (
	"log"
	"net/http"
	"strings"

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

// Responsible for user login and registration
func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	if !strings.HasPrefix(r.URL.Path, "/api/v1/auth/") {
		lib.WriteJSONResponse(w, http.StatusNotFound, lib.HandlerResponse{Message: "Endpoint not found"})
		return
	}

	action := strings.TrimPrefix(r.URL.Path, "/api/v1/auth/")

	switch action {
	case "login":
		var auth model.AuthRequest
		if err := lib.DecodeAndValidateRequest(r, &auth); err != nil {
			lib.WriteJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		iPAddr := r.RemoteAddr
		userAgent := r.Header.Get("User-Agent")
		method := r.Method
		email := auth.Email
		pw := auth.Password
		jwt, err := ah.AuthService.AuthenticateExistingUser(ctx, iPAddr, userAgent, method, email, pw)
		if err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error(), Token: nil, Data: nil})
			return
		}
		lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Login Successful", Token: &jwt, Data: nil})
		return
	case "register":
		var user model.User
		if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
			lib.WriteJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		iPAddr := r.RemoteAddr
		userAgent := r.Header.Get("User-Agent")
		method := r.Method
		jwt, err := ah.AuthService.AuthenticateNewUser(ctx, iPAddr, userAgent, method, user)
		if err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error(), Token: nil, Data: nil})
			return
		}

		lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Registration Successful", Token: &jwt, Data: nil})
		return
	}
}

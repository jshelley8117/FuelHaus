package handler

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

// Responsible for user login and registration
func (ah AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	switch r.URL.Path {
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
		if err := ah.AuthService.Login(ctx, iPAddr, userAgent, method, email); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
			return
		}
		lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Login Successful"})
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
		if err := ah.AuthService.Register(ctx, iPAddr, userAgent, method, user); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
			return
		}
		lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Registration Successful"})
		return
	}
}

package handler

import (
	"log"
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/lib"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

type UserHandler struct {
	UserService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// ServeHTTP is a UserHandler implementation of the net/http package's "ServeHTTP"
// function, which used to route requests to /users endpoint based on the HTTP Method
func (uh UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("email") {
			userEmail := r.URL.Query().Get("email")
			user, err := uh.UserService.GetUserByEmail(ctx, userEmail)
			if err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
				return
			}
			lib.WriteJSONResponse(w, http.StatusOK, user)
			return
		} else {
			users, err := uh.UserService.GetAllUsers(ctx)
			if err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
				return
			}
			lib.WriteJSONResponse(w, http.StatusOK, users)
			return
		}
	case http.MethodPost:
		var user model.User

		if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
			lib.WriteJSONResponse(w, http.StatusBadRequest, err)
			return
		}
		lib.SanitizeInput(&user)
		if err := uh.UserService.CreateUser(ctx, user); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
			return
		}
		lib.WriteJSONResponse(w, http.StatusOK, nil)
		return
	case http.MethodDelete:
		if r.URL.Query().Has("uid") {
			uid := r.URL.Query().Get("uid")
			if err := uh.UserService.DeleteUser(ctx, uid); err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
				return
			}
			lib.WriteJSONResponse(w, http.StatusOK, nil)
			return
		} else {
			lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: "No matching User found with User ID"})
			return
		}
	case http.MethodPut:
		if r.URL.Query().Has("uid") {
			var user model.User
			user.UserId = r.URL.Query().Get("uid")
			if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
				lib.WriteJSONResponse(w, http.StatusBadRequest, err)
				return
			}
			lib.SanitizeInput(&user)
			if err := uh.UserService.UpdateUser(ctx, user); err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
				return
			}
			lib.WriteJSONResponse(w, http.StatusOK, nil)
			return
		} else {
			lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: "No matching User found with User ID"})
			return
		}
	default:
		lib.WriteJSONResponse(w, http.StatusTeapot, lib.HandlerResponse{Message: "TEAPOT"})
		return
	}
}

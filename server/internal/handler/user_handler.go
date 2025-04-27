package handler

import (
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
	var user model.User

	if err := lib.DecodeAndValidateRequest(r, user); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, err)
	}

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("email") {
			userEmail := r.URL.Query().Get("email")
			user, err := uh.UserService.GetUserByEmail(userEmail)
			if err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
			}
			lib.WriteJSONResponse(w, http.StatusOK, user)
		} else {
			users, err := uh.UserService.GetAllUsers()
			if err != nil {
				lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
			}
			lib.WriteJSONResponse(w, http.StatusOK, users)
		}
	case http.MethodPost:
		if err := uh.UserService.CreateUser(user); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		}
		lib.WriteJSONResponse(w, http.StatusOK, nil)
	case http.MethodDelete:
		if err := uh.UserService.DeleteUser(user.UserId); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		}
		lib.WriteJSONResponse(w, http.StatusOK, nil)
	case http.MethodPut:
		if err := uh.UserService.UpdateUser(user); err != nil {
			lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		}
		lib.WriteJSONResponse(w, http.StatusOK, nil)
	}
}

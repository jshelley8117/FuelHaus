package v2

import (
	"fmt"
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

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	users, err := uh.UserService.GetAllUsers(ctx)
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{
			Message: err.Error(),
		})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, users)
}

func (uh *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	user, err := uh.UserService.GetUserByEmail(ctx, r.PathValue("email"))
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{
			Message: err.Error(),
		})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, user)
}

// func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
// 	ctx := r.Context()
// 	var user model.User

// 	if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
// 		lib.WriteJSONResponse(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	lib.SanitizeInput(&user)

// 	if err := uh.UserService.CreateUser(ctx, user, user.UserId); err != nil {
// 		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
// 		return
// 	}

// 	lib.WriteJSONResponse(w, http.StatusOK, nil)
// }

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()

	if err := uh.UserService.DeleteUser(ctx, r.PathValue("uid")); err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{
			Message: err.Error(),
		})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{
		Message: fmt.Sprintf("User %s deleted successfully", r.PathValue("uid")),
	})
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	var user model.User
	// TODO: Need to analyze whether or not this is needed, or if the model stored in DB needs an ID field since one is supplied by Firestore itself
	user.UserId = r.PathValue("uid")

	if err := lib.DecodeAndValidateRequest(r, &user); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{
			Message: err.Error(),
		})
		return
	}

	lib.SanitizeInput(&user)
	if err := uh.UserService.UpdateUser(ctx, user, user.UserId); err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{
			Message: err.Error(),
		})
		return
	}

	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{
		Message: fmt.Sprintf("User %s updated successfully", user.UserId),
	})
}

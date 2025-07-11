package service

import (
	"context"
	"log"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
	"golang.org/x/crypto/bcrypt"

	"cloud.google.com/go/firestore"

	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]model.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (model.UserResponse, error)
	CreateUser(ctx context.Context, reqUser model.User, uid string) error
	DeleteUser(ctx context.Context, userId string) error
	UpdateUser(ctx context.Context, reqUser model.User, uid string) error
}

type UserService struct {
	UserClient      *client.UserClient
	FirebaseService resource.FirebaseServices
}

func NewUserService(userClient *client.UserClient, firebaseService resource.FirebaseServices) *UserService {
	return &UserService{
		UserClient:      userClient,
		FirebaseService: firebaseService,
	}
}

// Returns a slice of Users
func (us *UserService) GetAllUsers(ctx context.Context) ([]model.UserResponse, error) {
	fsUsers, err := us.UserClient.FetchAllUsers(ctx, us.FirebaseService)
	if err != nil {
		return nil, err
	}

	return fsUsers, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (model.UserResponse, error) {
	fsUser, err := us.UserClient.FetchUserByEmail(ctx, us.FirebaseService, email)
	if err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		UserId:       fsUser.UserId,
		FirstName:    fsUser.FirstName,
		LastName:     fsUser.LastName,
		Email:        fsUser.Email,
		IsUserActive: fsUser.IsUserActive,
		CreatedAt:    fsUser.CreatedAt,
		UpdatedAt:    fsUser.UpdatedAt,
	}, nil
}

func (us *UserService) CreateUser(ctx context.Context, reqUser model.User, uid string) error {
	log.Println("Entered Service: CreateUser")

	reqUser.CreatedAt = time.Now()
	reqUser.UpdatedAt = time.Now()
	reqUser.IsUserActive = true
	if err := us.UserClient.CreateUser(ctx, us.FirebaseService, uid, model.User{
		CreatedAt:    reqUser.CreatedAt,
		Email:        strings.ToLower(reqUser.Email),
		FirstName:    reqUser.FirstName,
		LastName:     reqUser.LastName,
		IsUserActive: reqUser.IsUserActive,
		UpdatedAt:    reqUser.UpdatedAt,
	}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, userId string) error {
	log.Println("Entered Service: DeleteUser")

	if err := us.UserClient.DeleteUser(ctx, us.FirebaseService, userId); err != nil {
		return err
	}

	if err := us.FirebaseService.Auth.DeleteUser(ctx, userId); err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, u model.User, uid string) error {

	var updates []firestore.Update
	if strings.TrimSpace(u.FirstName) != "" {
		updates = append(updates, firestore.Update{Path: "first_name", Value: u.FirstName})
	}
	if strings.TrimSpace(u.LastName) != "" {
		updates = append(updates, firestore.Update{Path: "last_name", Value: u.LastName})
	}
	if strings.TrimSpace(u.Email) != "" {
		updates = append(updates, firestore.Update{Path: "email", Value: u.Email})
	}
	if strings.TrimSpace(u.Password) != "" {
		updates = append(updates, firestore.Update{Path: "password", Value: u.Password})
	}
	updates = append(updates, firestore.Update{Path: "updated_at", Value: time.Now()})

	// Update Firestore Document
	if err := us.UserClient.UpdateUser(ctx, us.FirebaseService, u.UserId, updates); err != nil {
		return err
	}

	// Update Firebase Auth User
	userToUpdate := (&auth.UserToUpdate{}).
		Email(u.Email).
		DisplayName(u.FirstName + " " + u.LastName)

	if strings.TrimSpace(u.Password) != "" {
		userToUpdate = userToUpdate.Password(u.Password)
	}

	if _, err := us.FirebaseService.Auth.UpdateUser(ctx, uid, userToUpdate); err != nil {
		return err
	}

	return nil
}

// PRIVATE FUNCTIONS BELOW

// Maps Database Response for GetAllUsers to UserList model
func mapUserModelToUserList(fsResp []model.User) []model.UserResponse {
	var uList []model.UserResponse
	for _, v := range fsResp {
		uList = append(uList, model.UserResponse{
			FirstName:    v.FirstName,
			LastName:     v.LastName,
			Email:        v.Email,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			IsUserActive: v.IsUserActive,
		})
	}
	return uList
}

func hashPassword(pw string) (string, error) {
	hpw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return pw, err
	}
	return string(hpw), nil
}

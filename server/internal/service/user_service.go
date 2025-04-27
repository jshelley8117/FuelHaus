package service

import (
	"context"
	"time"

	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IUserService interface {
	GetAllUsers(ctx context.Context) ([]UserList, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, reqUser model.User) error
	DeleteUser(ctx context.Context, userId string) error
	UpdateUser(ctx context.Context, reqUser model.User) error
}

type UserService struct {
	UserClient      client.UserClient
	FirebaseService resource.FirebaseServices
}

type UserList struct {
	UserId       string
	FirstName    string
	LastName     string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsUserActive bool
}

func NewUserService(userClient client.UserClient, firebaseService resource.FirebaseServices) UserService {
	return UserService{
		UserClient:      userClient,
		FirebaseService: firebaseService,
	}
}

// Returns a slice of Users - TODO: Need to map to some user response so that we aren't returning unnecessary data.
func (us *UserService) GetAllUsers(ctx context.Context) ([]UserList, error) {

	fsUsers, err := us.UserClient.FetchAllUsers(ctx, us.FirebaseService)
	if err != nil {
		return nil, err
	}
	users := mapUserModelToUserList(fsUsers)

	return users, nil
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	return user, nil
}

func (us *UserService) CreateUser(ctx context.Context, reqUser model.User) error {

	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, userId string) error {

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, reqUser model.User) error {

	return nil
}

// PRIVATE FUNCTIONS BELOW

// Maps Database Response for GetAllUsers to UserList model
func mapUserModelToUserList(fsResp []model.User) []UserList {
	var uList []UserList
	for _, v := range fsResp {
		uList = append(uList, UserList{
			UserId:       v.UserId,
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

package service

import (
	"time"

	"github.com/jshelley8117/FuelHaus/internal/model"
)

type IUserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
	CreateUser(reqUser model.User) error
	DeleteUser(userId uint16) error
	UpdateUser(reqUser model.User) error
}

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

type UserList struct {
	UserId       uint16
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsUserActive bool
}

// Returns a slice of Users - TODO: Need to map to some user response so that we aren't returning unnecessary data.
func GetAllUsers() ([]model.User, error) {

	// after successfully receiving users, map response to UserList
	return nil, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User

	return user, nil
}

func CreateUser(reqUser model.User) error {

	return nil
}

func DeleteUser(userId uint16) error {

	return nil
}

func UpdateUser(reqUser model.User) error {

	return nil
}

// PRIVATE FUNCTIONS BELOW

// Maps Database Response for GetAllUsers to UserResponseList model
func mapUserModelToUserList([]model.User) []UserList {
	var uList []UserList

	return uList
}

package service

import (
	"context"
	"time"

	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/lib"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IAuthService interface {
	Login(ctx context.Context, ipAddr string, userAgent string, method string, email string) error
	Register(ctx context.Context, ipAddr string, userAgent string, method string, u model.User) error
}

type AuthService struct {
	UserService     IUserService
	FirebaseService resource.FirebaseServices
	AuthClient      client.AuthClient
}

type AuthRequest struct {
}

func NewAuthService(userService IUserService, firebaseService resource.FirebaseServices, authClient client.AuthClient) AuthService {
	return AuthService{
		UserService:     userService,
		FirebaseService: firebaseService,
		AuthClient:      authClient,
	}
}

// Service Layer Implementation to perform Login Authentication
func (as *AuthService) Login(ctx context.Context, ipAddr string, userAgent string, method string, email string) error {
	// needs to perform a GET on the users table to ensure u exists in the users table
	response, err := as.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	// after a successful GET on users table, perform a POST to the authentication table to log a successful authentication
	authAttempt := model.AuthFirestoreRequest{
		UserId:    response.UserId,
		Email:     response.Email,
		CreatedAt: time.Now(),
		IPAddress: ipAddr,
		UserAgent: userAgent,
		Status:    lib.SUCCESSS,
		Method:    method,
	}
	as.AuthClient.CreateLoginRequest(ctx, as.FirebaseService, authAttempt)
	return nil
}

func (as *AuthService) Register(ctx context.Context, ipAddr string, userAgent string, method string, u model.User) error {
	// needs to perform a POST to the users table

	// after a successful POST to the users table, perform a POST to the authentication table to log a successful authentication

	return nil
}

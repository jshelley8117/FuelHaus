package service

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"

	"github.com/jshelley8117/FuelHaus/internal/client"
	// "github.com/jshelley8117/FuelHaus/internal/lib"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IAuthService interface {
	AuthenticateExistingUser(ctx context.Context, ipAddr, userAgent, method, email, pw string) (string, error)
	AuthenticateNewUser(ctx context.Context, ipAddr string, userAgent string, method string, u model.User) (string, error)
}

type AuthService struct {
	UserService     IUserService
	FirebaseService resource.FirebaseServices
	AuthClient      *client.AuthClient
}

type AuthRequest struct {
}

func NewAuthService(userService IUserService, firebaseService resource.FirebaseServices, authClient *client.AuthClient) *AuthService {
	return &AuthService{
		UserService:     userService,
		FirebaseService: firebaseService,
		AuthClient:      authClient,
	}
}

// Service Layer Implementation to perform Login Authentication
func (as *AuthService) AuthenticateExistingUser(ctx context.Context, ipAddr, userAgent, method, email, pw string) (string, error) {
	log.Println("Entered AuthenticateExistingUser")
	fbAuthClient := as.FirebaseService.Auth
	// var status string

	// verify user in firebase auth -> checks to see if email exists in firestore
	userRecord, err := fbAuthClient.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	jwt, err := generateJWT(ctx, fbAuthClient, userRecord.UID)
	if err != nil {
		return "", nil
	}
	return jwt, nil
}

func (as *AuthService) AuthenticateNewUser(ctx context.Context, ipAddr string, userAgent string, method string, u model.User) (string, error) {
	log.Println("Entered AuthenticateNewUser")
	fbAuthClient := as.FirebaseService.Auth

	// create params for firebase auth create user request
	params := (&auth.UserToCreate{}).
		Email(u.Email).
		Password(u.Password).
		EmailVerified(false).
		Disabled(u.IsUserActive)

	userRecord, err := fbAuthClient.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}
	u.UserId = userRecord.UID

	// store newly created user in internal Firestore
	if err := as.UserService.CreateUser(ctx, u, u.UserId); err != nil {
		return "", err
	}

	jwt, err := generateJWT(ctx, fbAuthClient, userRecord.UID)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

// Orchestrates the process for creating and sending an Authentication Request
// func (as *AuthService) logAuthRequest(ctx context.Context, email, ipAddr, userAgent, method, authType, status string) error {
// 	authAttempt := generateAuthRequestPayload(email, ipAddr, userAgent, method, authType)
// 	authAttempt.Status = status
// 	if err := as.AuthClient.CreateAuthenticationRequest(ctx, as.FirebaseService, authAttempt); err != nil {
// 		return err
// 	}
// 	return nil
// }

// Creates an Authentication History request payload
// func generateAuthRequestPayload(email, ipAddr, userAgent, method, authType string) model.AuthFirestoreRequest {
// 	return model.AuthFirestoreRequest{
// 		Email:     email,
// 		CreatedAt: time.Now(),
// 		IPAddress: ipAddr,
// 		UserAgent: userAgent,
// 		Method:    method,
// 		AuthType:  authType,
// 	}
// }

// Generates a Custom Client token to be used on subsequent requests to the server
func generateJWT(ctx context.Context, fbAuthClient *auth.Client, uid string) (string, error) {
	jwtToken, err := fbAuthClient.CustomToken(ctx, uid)
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}

// Returns an error string pertaining to what the specific Firebase Auth Error is
// func handleFirebaseAuthError(err error) string {
// 	if status.Code(err) == codes.NotFound {
// 		return "Email does not exist"
// 	}
// 	return "Firebase Auth Error: " + err.Error()
// }

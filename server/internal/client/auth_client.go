package client

import (
	"context"

	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type AuthClient struct{}

func NewAuthClient() AuthClient {
	return AuthClient{}
}

func (ac *AuthClient) CreateLoginRequest(ctx context.Context, firebaseService resource.FirebaseServices, a model.AuthFirestoreRequest) {
	firestoreClient := firebaseService.Firestore
	firestoreClient.Doc("auth").Create(ctx, a)
}

func (ac *AuthClient) CreateRegisterRequest() {

}

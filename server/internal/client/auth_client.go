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

func (ac *AuthClient) CreateAuthenticationRequest(ctx context.Context, firebaseService resource.FirebaseServices, a model.AuthFirestoreRequest) error {
	firestoreClient := firebaseService.Firestore
	_, _, err := firestoreClient.Collection("auth_history").Add(ctx, a)
	if err != nil {
		return err
	}
	return nil
}

package client

import (
	"context"
	"log"

	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"google.golang.org/api/iterator"
)

type UserClient struct{}

func NewUserClient() UserClient {
	return UserClient{}
}

func (uc *UserClient) FetchAllUsers(ctx context.Context, firebaseServices resource.FirebaseServices) ([]model.User, error) {
	firestoreClient := firebaseServices.Firestore

	docIter := firestoreClient.Collection("users").Documents(ctx)
	defer docIter.Stop()

	var users []model.User
	for {
		doc, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user model.User
		err = doc.DataTo(&user)
		if err != nil {
			log.Printf("Failed to map Firestore document to User struct: %v", err)
		}
		user.UserId = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

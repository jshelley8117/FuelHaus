package client

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"google.golang.org/api/iterator"
)

type UserClient struct{}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (uc *UserClient) FetchAllUsers(ctx context.Context, firebaseServices resource.FirebaseServices) ([]model.UserResponse, error) {
	firestoreClient := firebaseServices.Firestore
	docIter := firestoreClient.Collection("users").Documents(ctx)
	defer docIter.Stop()
	var users []model.UserResponse
	for {
		doc, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user model.UserResponse
		if err := doc.DataTo(&user); err != nil {
			log.Printf("Client Error: Failed to map Firestore document to UserResponse struct [%v]", err)
			return nil, err
		}
		user.UserId = doc.Ref.ID
		users = append(users, user)
	}
	return users, nil
}

func (uc *UserClient) FetchUserByEmail(ctx context.Context, firebaseServices resource.FirebaseServices, email string) (model.UserResponse, error) {
	firestoreClient := firebaseServices.Firestore
	query := firestoreClient.Collection("users").Where("email", "==", email).Limit(1)
	docIter := query.Documents(ctx)
	defer docIter.Stop()
	doc, err := docIter.Next()
	if err == iterator.Done {
		return model.UserResponse{}, fmt.Errorf("the following email could not be found [%s]", email)
	}
	if err != nil {
		return model.UserResponse{}, err
	}
	var user model.UserResponse
	if err := doc.DataTo(&user); err != nil {
		log.Printf("Client Error: Failed to fetch user by email in firestore [%v]", err)
		return model.UserResponse{}, err
	}
	user.UserId = doc.Ref.ID
	return user, nil
}

func (uc *UserClient) CreateUser(ctx context.Context, firebaseServices resource.FirebaseServices, uid string, u model.User) error {
	log.Println("Entered Client: CreateUser")
	firestoreClient := firebaseServices.Firestore
	_, err := firestoreClient.Collection("users").Doc(uid).Create(ctx, u)
	if err != nil {
		log.Printf("Client Error: Failed to create user in firestore [%v]", err)
		return err
	}
	return nil
}

func (uc *UserClient) DeleteUser(ctx context.Context, firebaseServices resource.FirebaseServices, userId string) error {
	log.Println("Entered Client: DeleteUser")
	firestoreClient := firebaseServices.Firestore
	_, err := firestoreClient.Collection("users").Doc(userId).Delete(ctx)
	if err != nil {
		log.Printf("Client Error: Failed to delete user in firestore [%v]", err)
		return err
	}
	return nil
}

func (uc *UserClient) UpdateUser(ctx context.Context, firebaseServices resource.FirebaseServices, uid string, updates []firestore.Update) error {
	log.Println("Entered Client: UpdateUser")
	firestoreClient := firebaseServices.Firestore
	_, err := firestoreClient.Collection("users").Doc(uid).Update(ctx, updates)
	if err != nil {
		log.Printf("Client Error: Failed to update user in firestore [%v]", err)
		return err
	}
	return nil
}

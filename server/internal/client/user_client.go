package client

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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

func (uc *UserClient) FetchUserByEmail(ctx context.Context, firebaseServices resource.FirebaseServices, email string) (model.User, error) {
	firestoreClient := firebaseServices.Firestore

	query := firestoreClient.Collection("users").Where("email", "==", email).Limit(1)
	docIter := query.Documents(ctx)
	defer docIter.Stop()

	doc, err := docIter.Next()
	if err == iterator.Done {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	if err := doc.DataTo(&user); err != nil {
		log.Printf("Failed to fetch user by email in firestore: %v", err)
		return model.User{}, err
	}
	user.UserId = doc.Ref.ID
	return user, nil
}

func (uc *UserClient) CreateUser(ctx context.Context, firebaseServices resource.FirebaseServices, u model.User) error {
	log.Println("Entered Client: CreateUser")
	firestoreClient := firebaseServices.Firestore
	_, _, err := firestoreClient.Collection("users").Add(ctx, u)
	if err != nil {
		log.Printf("Failed to create user in firestore: %v", err)
		return err
	}
	return nil
}

func (uc *UserClient) DeleteUser(ctx context.Context, firebaseServices resource.FirebaseServices, userId string) error {
	log.Println("Entered Client: DeleteUser")
	// TODO: Need to add functionality that also deletes the user from firebase auth service
	firestoreClient := firebaseServices.Firestore
	_, err := firestoreClient.Collection("users").Doc(userId).Delete(ctx)
	if err != nil {
		log.Printf("Failed to delete user in firestore: %v", err)
		return err
	}
	return nil
}

func (uc *UserClient) UpdateUser(ctx context.Context, firebaseServices resource.FirebaseServices, uid string, updates []firestore.Update) error {
	log.Println("Entered Client: UpdateUser")
	firestoreClient := firebaseServices.Firestore
	_, err := firestoreClient.Collection("users").Doc(uid).Update(ctx, updates)
	if err != nil {
		log.Printf("Failed to update user in firestore: %v", err)
		return err
	}
	return nil
}

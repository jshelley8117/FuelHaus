package model

import "time"

type User struct {
	UserId       string    `firestore:"userId,omitempty" json:"userId"`
	FirstName    string    `firestore:"first_name" json:"first_name"`
	LastName     string    `firestore:"last_name" json:"last_name"`
	Email        string    `firestore:"email" json:"email"`
	Password     string    `firestore:"password" json:"password"`
	CreatedAt    time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at" json:"updated_at"`
	IsUserActive bool      `firestore:"is_active" json:"is_active"`
}
type UserResponse struct {
	UserId       string
	FirstName    string
	LastName     string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsUserActive bool
}

package model

import "time"

type User struct {
	UserId       string    `firestore:"userId" json:"userId"`
	FirstName    string    `firestore:"first_name" json:"first_name" validate:"required"`
	LastName     string    `firestore:"last_name" json:"last_name" validate:"required"`
	Email        string    `firestore:"email" json:"email" validate:"required,email"`
	Password     string    `firestore:"password" json:"password" validate:"required"`
	CreatedAt    time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at" json:"updated_at"`
	IsUserActive bool      `firestore:"is_active" json:"is_active"`
}

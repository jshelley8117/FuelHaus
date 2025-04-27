package model

import "time"

type User struct {
	UserId    string    `firestore:"userId"`
	FirstName string    `firestore:"first_name" validate:"required"`
	LastName  string    `firestore:"last_name" validate:"required"`
	Email     string    `firestore:"email" validate:"required,email"`
	Password  string    `firestore:"password" validate:"required"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
	Is_Active bool      `firestore:"is_active"`
}

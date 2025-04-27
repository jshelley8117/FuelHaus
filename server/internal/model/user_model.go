package model

import "time"

type User struct {
	UserId       uint16    `json:"userId,omitempty"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	IsUserActive bool      `json:"isUserActive,omitempty"`
}

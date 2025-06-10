package model

import (
	"time"
)

// TODO: need to define struct tags
type AuthFirestoreRequest struct {
	Email     string
	CreatedAt time.Time
	IPAddress string
	UserAgent string
	Status    string
	Location  string // can be for future implementation
	Method    string
	AuthType  string
}

type AuthRequest struct {
	Email    string
	Password string
}

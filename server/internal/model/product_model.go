package model

import "time"

// Product struct that will be stored in DB
type Product struct {
	ProductId       string    `firestore:"-" json:"productId"`
	Name            string    `firestore:"name" json:"name"`
	Description     string    `firestore:"description,omitempty" json:"description"`
	Category        string    `firestore:"category" json:"category"`
	Price           int       `firestore:"priceInCents" json:"price"`
	CreatedAt       time.Time `firestore:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time `firestore:"updatedAt" json:"updatedAt"`
	IsProductActive bool      `firestore:"isProductActive" json:"isProductActive"`
}

// Product struct that comes in as a Request from Client
type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price" validate:"required,min=0"` //  Accept as as float from the client
}

type ProductResponse struct {
	ProductId       string    `json:"productId"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	Price           float64   `json:"price"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	IsProductActive bool      `json:"isProductActive"`
}

// Convert cents to dollar amount
func CentsToDollarAmount(cents int) float64 {
	return float64(cents) / 100.0
}

// Convert dollar amount to cents
func DollarAmountToCents(dollars float64) int {
	return int(dollars*100 + 0.5) // + 0.5 for proper rounding
}

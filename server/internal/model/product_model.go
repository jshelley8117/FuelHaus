package model

import "time"

type Product struct {
	ProductId       string    `firestore:"-" json:"productId"`
	Name            string    `firestore:"name" json:"name" validate:"required"`
	Description     string    `firestore:"description,omitempty" json:"description"`
	Category        string    `firestore:"category" json:"category"`
	Price           float32   `firestore:"price" json:"price" validate:"required,min=0"`
	CreatedAt       time.Time `firestore:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time `firestore:"updatedAt" json:"updatedAt"`
	IsProductActive bool      `firestore:"isProductActive" json:"isProductActive"`
}

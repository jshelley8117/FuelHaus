package client

import (
	"context"
	"log"

	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"google.golang.org/api/iterator"
)

type ProductClient struct{}

func NewProductClient() *ProductClient {
	return &ProductClient{}
}

func (pc *ProductClient) CreateProduct(ctx context.Context, firebaseService resource.FirebaseServices, product model.Product) error {
	log.Println("Entered Client: CreateProduct")

	firestoreClient := firebaseService.Firestore
	_, _, err := firestoreClient.Collection("products").Add(ctx, product)
	if err != nil {
		log.Println("Client Error: Failed to create new product in firestore")
		return err
	}
	return nil
}

func (pc *ProductClient) GetAllProducts(ctx context.Context, firebaseService resource.FirebaseServices) ([]model.Product, error) {
	log.Println("Entered Client: GetAllProducts")
	firestoreClient := firebaseService.Firestore
	docIter := firestoreClient.Collection("products").Documents(ctx)
	defer docIter.Stop()
	var products []model.Product
	for {
		doc, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var product model.Product
		if err := doc.DataTo(&product); err != nil {
			log.Printf("Client Error: Failed to map Firestore document to Product struct: %v", err)
			return nil, err
		}
		product.ProductId = doc.Ref.ID
		products = append(products, product)
	}
	return products, nil
}

func (pc *ProductClient) GetProductById(ctx context.Context, firebaseService resource.FirebaseServices, id string) (model.Product, error) {
	log.Println("Entered Client: GetProductById")
	firestoreClient := firebaseService.Firestore
	doc, err := firestoreClient.Collection("products").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Client Error: Failed to get product with ID %s: %v", id, err)
		return model.Product{}, err
	}
	if !doc.Exists() {
		log.Printf("Client Error: Product with ID %s does not exist", id)
		return model.Product{}, err
	}
	var product model.Product
	if err := doc.DataTo(&product); err != nil {
		log.Printf("Client Error: Failed to map Firestore document to Product struct (ID: %s): %v", id, err)
		return model.Product{}, err
	}
	product.ProductId = doc.Ref.ID
	return product, nil
}

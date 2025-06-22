package service

import (
	"context"
	"log"
	"time"

	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req model.Product) error
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	UpdateProduct(ctx context.Context, req model.Product) error
	GetProductById(ctx context.Context, id string) (model.Product, error)
}

type ProductService struct {
	ProductClient   *client.ProductClient
	FirebaseService resource.FirebaseServices
}

func NewProductService(productClient *client.ProductClient, firebaseService resource.FirebaseServices) *ProductService {
	return &ProductService{
		ProductClient:   productClient,
		FirebaseService: firebaseService,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, req model.Product) error {
	log.Println("Entered Service: CreateProduct")
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	req.IsProductActive = true
	if err := ps.ProductClient.CreateProduct(ctx, ps.FirebaseService, req); err != nil {
		log.Println("Service Error: Failed to create Product")
		return err
	}
	return nil
}

func (ps *ProductService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	log.Println("Entered Service: GetAllProducts")
	products, err := ps.ProductClient.GetAllProducts(ctx, ps.FirebaseService)
	if err != nil {
		log.Println("Service Error: Failed to retreive all products")
		return []model.Product{}, err
	}
	return products, nil
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id string) error {

	return nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, req model.Product) error {

	return nil
}

func (ps *ProductService) GetProductById(ctx context.Context, id string) (model.Product, error) {
	log.Println("Entered Service: GetProductById")
	product, err := ps.ProductClient.GetProductById(ctx, ps.FirebaseService, id)
	if err != nil {
		log.Println("Service Error: Failed to retrieve product by ID")
		return model.Product{}, err
	}
	return product, nil
}

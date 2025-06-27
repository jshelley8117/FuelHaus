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
	CreateProduct(ctx context.Context, req model.ProductRequest) error
	GetAllProducts(ctx context.Context) ([]model.ProductResponse, error)
	DeleteProductById(ctx context.Context, id string) error
	UpdateProduct(ctx context.Context, req model.Product) error
	GetProductById(ctx context.Context, id string) (model.ProductResponse, error)
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

func (ps *ProductService) CreateProduct(ctx context.Context, req model.ProductRequest) error {
	log.Println("Entered Service: CreateProduct")
	product := model.Product{
		Name:            req.Name,
		Description:     req.Description,
		Category:        req.Category,
		Price:           model.DollarAmountToCents(req.Price),
		IsProductActive: true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := ps.ProductClient.CreateProduct(ctx, ps.FirebaseService, product); err != nil {
		log.Println("Service Error: Failed to create Product")
		return err
	}
	return nil
}

func (ps *ProductService) GetAllProducts(ctx context.Context) ([]model.ProductResponse, error) {
	log.Println("Entered Service: GetAllProducts")
	products, err := ps.ProductClient.GetAllProducts(ctx, ps.FirebaseService)
	if err != nil {
		log.Println("Service Error: Failed to retreive all products")
		return []model.ProductResponse{}, err
	}
	productsResponse := make([]model.ProductResponse, len(products))
	for i, product := range products {
		productsResponse[i] = mapProductToProductResponse(product)
	}
	return productsResponse, nil
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id string) error {

	return nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, req model.Product) error {

	return nil
}

func (ps *ProductService) GetProductById(ctx context.Context, id string) (model.ProductResponse, error) {
	log.Println("Entered Service: GetProductById")
	product, err := ps.ProductClient.GetProductById(ctx, ps.FirebaseService, id)
	productResponse := mapProductToProductResponse(product)
	if err != nil {
		log.Println("Service Error: Failed to retrieve product by ID")
		return model.ProductResponse{}, err
	}
	return productResponse, nil
}

func (ps *ProductService) DeleteProductById(ctx context.Context, id string) error {
	log.Println("Entered Service: DeleteProductById")
	if err := ps.ProductClient.DeleteProductById(ctx, ps.FirebaseService, id); err != nil {
		return err
	}
	return nil
}

func mapProductToProductResponse(p model.Product) model.ProductResponse {
	return model.ProductResponse{
		ProductId:       p.ProductId,
		Name:            p.Name,
		Description:     p.Description,
		Category:        p.Category,
		Price:           model.CentsToDollarAmount(p.Price),
		IsProductActive: p.IsProductActive,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

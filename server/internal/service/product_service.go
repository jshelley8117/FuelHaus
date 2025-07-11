package service

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/resource"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req model.ProductRequest) error
	GetAllProducts(ctx context.Context) ([]model.ProductResponse, error)
	DeleteProductById(ctx context.Context, id string) error
	UpdateProductById(ctx context.Context, id string, req model.ProductUpdateRequest) error
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

func (ps *ProductService) UpdateProductById(ctx context.Context, id string, req model.ProductUpdateRequest) error {
	log.Println("Entered Service: UpdateProductById")
	var updates []firestore.Update
	if req.Name != nil {
		updates = append(updates, firestore.Update{
			Path:  "name",
			Value: req.Name,
		})
	}
	if req.Description != nil {
		updates = append(updates, firestore.Update{
			Path:  "description",
			Value: req.Description,
		})
	}
	if req.Category != nil {
		updates = append(updates, firestore.Update{
			Path:  "category",
			Value: req.Category,
		})
	}
	if req.Price != nil {
		updates = append(updates, firestore.Update{
			Path:  "priceInCents",
			Value: model.DollarAmountToCents(*req.Price),
		})
	}
	if req.IsProductActive != nil {
		updates = append(updates, firestore.Update{
			Path:  "isProductActive",
			Value: req.IsProductActive,
		})
	}
	updates = append(updates, firestore.Update{
		Path:  "updatedAt",
		Value: time.Now(),
	})
	if err := ps.ProductClient.UpdateProductById(ctx, ps.FirebaseService, id, updates); err != nil {
		return err
	}
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

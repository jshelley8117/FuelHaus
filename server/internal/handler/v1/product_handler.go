package v1

import (
	"log"
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/lib"
	"github.com/jshelley8117/FuelHaus/internal/model"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

type ProductHandler struct {
	ProductService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	var product model.ProductRequest
	if err := lib.DecodeAndValidateRequest(r, &product); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: err.Error()})
		return
	}
	if err := ph.ProductService.CreateProduct(ctx, product); err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.WriteJSONResponse(w, http.StatusCreated, lib.HandlerResponse{Message: "Product created successfully"})
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	products, err := ph.ProductService.GetAllProducts(ctx)
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{
		Message: "Products fetched successfully",
		Data:    products,
	})
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	id := r.PathValue("id")
	product, err := ph.ProductService.GetProductById(ctx, id)
	if err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{
		Message: "Product fetched successfully",
		Data:    product,
	})
}

func (pg *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling HTTP Request:\nMethod: %v\nPath: %v", r.Method, r.URL.Path)
	ctx := r.Context()
	id := r.PathValue("id")
	if err := pg.ProductService.DeleteProductById(ctx, id); err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.WriteJSONResponse(w, http.StatusOK, lib.HandlerResponse{Message: "Product deleted successfully"})
}

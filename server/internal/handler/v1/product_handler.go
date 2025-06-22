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
	var product model.Product
	if err := lib.DecodeAndValidateRequest(r, product); err != nil {
		lib.WriteJSONResponse(w, http.StatusBadRequest, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.SanitizeInput(&product)
	if err := ph.ProductService.CreateProduct(ctx, product); err != nil {
		lib.WriteJSONResponse(w, http.StatusInternalServerError, lib.HandlerResponse{Message: err.Error()})
		return
	}
	lib.WriteJSONResponse(w, http.StatusCreated, lib.HandlerResponse{Message: "Product Created Successfully"})
}

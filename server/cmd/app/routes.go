package main

import (
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/client"
	v1 "github.com/jshelley8117/FuelHaus/internal/handler/v1"
	v2 "github.com/jshelley8117/FuelHaus/internal/handler/v2"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

func SetupRoutes(mux *http.ServeMux, firebaseServices resource.FirebaseServices) {

	// users API
	userClient := client.NewUserClient()
	userService := service.NewUserService(userClient, firebaseServices)
	v1UserHandler := v1.NewUserHandler(userService)
	v2UserHandler := v2.NewUserHandler(userService)

	// v1 users routes
	mux.Handle("/api/v1/users", v1UserHandler)

	// v2 users routes
	mux.HandleFunc("GET /api/v2/users", v2UserHandler.GetAllUsers)
	mux.HandleFunc("GET /api/v2/users/{email}", v2UserHandler.GetUserByEmail)
	// mux.HandleFunc("POST /api/v2/users", v2UserHandler.CreateUser)
	mux.HandleFunc("DELETE /api/v2/users/{uid}", v2UserHandler.DeleteUser)
	mux.HandleFunc("PUT /api/v2/users/{uid}", v2UserHandler.UpdateUser)

	// auth API
	authClient := client.NewAuthClient()
	authService := service.NewAuthService(userService, firebaseServices, authClient)
	v1AuthHandler := v1.NewAuthHandler(authService)
	v2AuthHandler := v2.NewAuthHandler(authService)

	// v1 auth routes
	mux.Handle("/api/v1/auth/", v1AuthHandler)

	// v2 auth routes
	mux.HandleFunc("POST /api/v2/auth/register", v2AuthHandler.Register)
	mux.HandleFunc("POST /api/v2/auth/login", v2AuthHandler.Login)

	// products API
	productClient := client.NewProductClient()
	productService := service.NewProductService(productClient, firebaseServices)
	productHandler := v1.NewProductHandler(productService)

	// v1 product routes
	mux.HandleFunc("POST /api/v1/products", productHandler.HandleCreateProduct)
	mux.HandleFunc("GET /api/v1/products", productHandler.HandleGetAllProducts)
	mux.HandleFunc("GET /api/v1/products/{id}", productHandler.HandleGetProductById)
	mux.HandleFunc("DELETE /api/v1/products/{id}", productHandler.HandleDeleteProductById)
	mux.HandleFunc("PATCH /api/v1/products/{id}", productHandler.HandleUpdateProductById)
}

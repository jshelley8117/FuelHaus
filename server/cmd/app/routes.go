package main

import (
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/client"
	v1 "github.com/jshelley8117/FuelHaus/internal/handler/v1"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

func SetupRoutes(mux *http.ServeMux, firebaseServices resource.FirebaseServices) {

	// users API
	userClient := client.NewUserClient()
	userService := service.NewUserService(userClient, firebaseServices)
	v1UserHandler := v1.NewUserHandler(&userService)

	// v1 users routes
	mux.Handle("/users", v1UserHandler)

	// v2 users routes

	// auth API
	authClient := client.NewAuthClient()
	authService := service.NewAuthService(&userService, firebaseServices, authClient)
	v1AuthHandler := v1.NewAuthHandler(&authService)

	// v1 auth routes
	mux.Handle("/auth/", v1AuthHandler)
}

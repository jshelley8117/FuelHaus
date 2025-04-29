package main

import (
	"net/http"

	"github.com/jshelley8117/FuelHaus/internal/client"
	"github.com/jshelley8117/FuelHaus/internal/handler"
	"github.com/jshelley8117/FuelHaus/internal/resource"
	"github.com/jshelley8117/FuelHaus/internal/service"
)

func SetupRoutes(mux *http.ServeMux, firebaseServices resource.FirebaseServices) {

	userClient := client.NewUserClient()
	userService := service.NewUserService(userClient, firebaseServices)
	userHandler := handler.NewUserHandler(&userService)

	mux.Handle("/users", userHandler)
}

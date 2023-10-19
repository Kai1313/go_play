package router

import (
	"go_play/middleware"
	"go_play/repositories"

	"github.com/gorilla/mux"
)

func InventoryRouter(repository repositories.InventoryRepository, router *mux.Router) {
	InventoryHandler := middleware.NewInventoryRepository(&repository)

	// Use the CORS middleware here
	router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/api/inventory", InventoryHandler.GetAllInventoryHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/inventory/last", InventoryHandler.GetLastInventoryHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/inventory", InventoryHandler.CreateInventoryHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/inventory/{id}", InventoryHandler.ShowInventoryHandler).Methods("GET", "OPTIONS")
}

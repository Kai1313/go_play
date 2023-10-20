package router

import (
	"go_play/middleware"
	"go_play/repositories"

	"github.com/gorilla/mux"
)

func GenerateNumberRouter(repository repositories.GenerateNumberRepository, router *mux.Router) {
	GenerateNumberHandler := middleware.NewGenerateNumberRepository(&repository)

	// Use the CORS middleware here
	router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/api/generateNumber", GenerateNumberHandler.GetAllGenerateNumberHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/generateNumber", GenerateNumberHandler.CreateGenerateNumberHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/generateNumber/{id}", GenerateNumberHandler.ShowGenerateNumberHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/generateNumber/last/{id}", GenerateNumberHandler.ShowLastGenerateNumberHandler).Methods("GET", "OPTIONS")
}

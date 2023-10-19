package main

import (
	"fmt"
	"go_play/models"
	"go_play/repositories"
	"go_play/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct the DSN from environment variables
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.Inventory{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	inventoryRepository := repositories.NewInventory(db)

	routers := mux.NewRouter()
	router.InventoryRouter(*inventoryRepository, routers)

	fmt.Printf("Main is running")
	log.Fatal(http.ListenAndServe(":8080", routers))
}

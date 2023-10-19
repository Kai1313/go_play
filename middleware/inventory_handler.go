package middleware

import (
	"encoding/json"
	"go_play/models"
	"go_play/repositories"
	"go_play/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InventoryHandler struct {
	repository *repositories.InventoryRepository
}

func NewInventoryRepository(repository *repositories.InventoryRepository) *InventoryHandler {
	return &InventoryHandler{
		repository: repository,
	}
}

func (h *InventoryHandler) ShowInventoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	inventoryJson, err := services.GetInventory(int(id), *h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
	}

	res := Response{
		Success: true,
		Data:    inventoryJson,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *InventoryHandler) CreateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var inventory models.InventoryJson
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		log.Fatalf("Unable to decode the requested body. %v", err)
	}

	inventoryDB := models.Inventory{
		InventoryCode: inventory.InventoryCode,
		InventoryName: inventory.InventoryName,
	}

	inventoryJson, err := services.CreateInventory(&inventoryDB, *h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
	}

	res := Response{
		Success: true,
		Data:    inventoryJson,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *InventoryHandler) GetAllInventoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	inventoryJsonList, err := services.GetAllInventory(*h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Success: true,
		Data:    inventoryJsonList,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *InventoryHandler) GetLastInventoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	generatedNumber, err := services.GenerateInventoryCode(*h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Success: true,
		Data:    generatedNumber,
	}
	json.NewEncoder(w).Encode(res)
}
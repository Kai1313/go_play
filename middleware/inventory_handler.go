package middleware

import (
	"bytes"
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

	// Make an HTTP GET request to GetLastNumber API
	// resp, err := http.Get("http://localhost:8080/api/generateNumber/last/T02")
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// defer resp.Body.Close()

	// // Decode the response from GetLastNumber
	// if resp.StatusCode != http.StatusOK {
	//     http.Error(w, "Failed to get last number", resp.StatusCode)
	//     return
	// }

	// var lastNumber Response
	// if err := json.NewDecoder(resp.Body).Decode(&lastNumber); err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// lastNumberValue := lastNumber.Data.(string)

	// Define the payload for the POST request
	payload := struct {
		ModuleID  string `json:"module_id"`
		Source    string `json:"source"`
		Warehouse string `json:"warehouse"`
	}{
		ModuleID:  inventory.ModuleId,  // Set the desired module_id
		Source:    inventory.Source, // Set the source value
		Warehouse: inventory.Warehouse, // Set the warehouse value
	}

	// Serialize the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to serialize the request payload", http.StatusInternalServerError)
		return
	}

	// Make an HTTP POST request to the GetLastNumber API
	resp, err := http.Post("http://localhost:8080/api/generateNumber/last", "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the response from GetLastNumber
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get last number", resp.StatusCode)
		return
	}

	var lastNumber Response
	if err := json.NewDecoder(resp.Body).Decode(&lastNumber); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastNumberValue := lastNumber.Data.(string)

	inventoryDB := models.Inventory{
		InventoryCode: lastNumberValue,
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

package services

import (
	"fmt"
	"go_play/models"
	"go_play/repositories"
	"strconv"
	"strings"
)

func CreateInventory(inventory *models.Inventory, repository repositories.InventoryRepository) (*models.InventoryJson, error) {
	inventoryDB, err := repository.Create(inventory)
	if err != nil {
		return nil, err
	}

	inventoryId := int(inventoryDB.InventoryId)
	inventoryJson := models.InventoryJson{
		InventoryId:   &inventoryId,
		InventoryCode: inventoryDB.InventoryCode,
		InventoryName: inventoryDB.InventoryName,
	}

	return &inventoryJson, nil
}

func GetInventory(id int, repository repositories.InventoryRepository) (*models.InventoryJson, error) {
	inventoryDB, err := repository.Show(id)
	if err != nil {
		return nil, err
	}

	inventoryId := int(inventoryDB.InventoryId)
	inventoryJson := models.InventoryJson{
		InventoryId:   &inventoryId,
		InventoryCode: inventoryDB.InventoryCode,
		InventoryName: inventoryDB.InventoryName,
	}

	return &inventoryJson, nil
}

func GetAllInventory(repository repositories.InventoryRepository) ([]models.InventoryJson, error) {
	inventoryList, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	var inventoryJsonList []models.InventoryJson
	for _, inventoryDB := range inventoryList {
		inventoryId := int(inventoryDB.InventoryId)
		inventoryJson := models.InventoryJson{
			InventoryId:   &inventoryId,
			InventoryCode: inventoryDB.InventoryCode,
			InventoryName: inventoryDB.InventoryName,
		}
		inventoryJsonList = append(inventoryJsonList, inventoryJson)
	}

	return inventoryJsonList, nil
}

// GenerateInventoryCode generates a unique inventory code based on the existing codes in the repository.
func GenerateInventoryCode(repository repositories.InventoryRepository) (string, error) {
	// Retrieve the last used inventory code from the repository
	lastInventory, err := repository.GetLast()
	if err != nil {
		return "", err
	}

	if lastInventory == nil {
		// If no inventory records exist, start with an initial code
		return "INV.00001", nil
	}

	// Parse the last used inventory code
	lastCode := lastInventory.InventoryCode

	// Increment the code
	newCode, err := incrementInventoryCode(lastCode)
	if err != nil {
		return "", err
	}

	return newCode, nil
}

// incrementInventoryCode increments the inventory code.
func incrementInventoryCode(code string) (string, error) {
	// Parse the code to extract the numeric part
	codeParts := strings.Split(code, ".")
	if len(codeParts) != 2 {
		return "", fmt.Errorf("Invalid inventory code format")
	}

	numericPart := codeParts[1]
	numericValue, err := strconv.Atoi(numericPart)
	if err != nil {
		return "", err
	}

	// Increment the numeric value
	numericValue++

	// Format the new code
	newCode := fmt.Sprintf("INV.%05d", numericValue)
	return newCode, nil
}

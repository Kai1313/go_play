package repositories

import (
	"fmt"
	"go_play/models"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventory(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) GetAll() ([]models.Inventory, error) {
	var inventoryList []models.Inventory
	result := r.db.Find(&inventoryList)

	if result.Error != nil {
		return nil, result.Error
	}

	return inventoryList, nil
}

func (r *InventoryRepository) Show(id int) (*models.Inventory, error) {
	var inventory models.Inventory
	result := r.db.Find(&inventory, id)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Inventory not found")
	}

	return &inventory, nil
}

func (r *InventoryRepository) Create(inventory *models.Inventory) (*models.Inventory, error) {
	err := r.db.Save(inventory).Error
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (r *InventoryRepository) GetLast() (*models.Inventory, error) {
	var lastInventory models.Inventory
    // Order the results in descending order of inventory_id and select the first record.
    result := r.db.Order("inventory_code desc").First(&lastInventory)

    if result.Error != nil {
        return nil, result.Error
    }

    return &lastInventory, nil
}
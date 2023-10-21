package models

type Inventory struct {
	InventoryId   uint `gorm:"PrimaryKey"`
	InventoryCode string
	InventoryName string
}

type InventoryJson struct {
	InventoryId   *int   `json:"inventory_id"`
	InventoryCode string `json:"inventory_code"`
	InventoryName string `json:"inventory_name"`
	ModuleId      string `json:"module_id"`
	Source        string `json:"source"`
	Warehouse     string `json:"warehouse"`
}

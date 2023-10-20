package models

type GenerateNumber struct {
	GenerateId uint   `gorm:"PrimaryKey"`
	ModuleId   string `gorm:"type: varchar(50)"`
	DefaultNo  string `gorm:"type: varchar(254)"`
	TableData  string `gorm:"type: varchar(254)"`
	FieldData  string `gorm:"type: varchar(254)"`
}

type GenerateNumberJson struct {
	GenerateId *int   `json:"generate_id"`
	ModuleId   string `json:"module_id"`
	DefaultNo  string `json:"default_no"`
	TableData  string `json:"table_data"`
	FieldData  string `json:"field_data"`
}

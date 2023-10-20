package repositories

import (
	"fmt"
	"go_play/models"

	"gorm.io/gorm"
)

type GenerateNumberRepository struct {
	db *gorm.DB
}

func NewGenerateNumber(db *gorm.DB) *GenerateNumberRepository {
	return &GenerateNumberRepository{
		db: db,
	}
}

func (r *GenerateNumberRepository) GetAll() ([]models.GenerateNumber, error) {
	var genNumberList []models.GenerateNumber
	result := r.db.Find(&genNumberList)

	if result.Error != nil {
		return nil, result.Error
	}

	return genNumberList, nil
}

func (r *GenerateNumberRepository) Show(id int) (*models.GenerateNumber, error) {
	var genNumber models.GenerateNumber
	result := r.db.Find(&genNumber, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Number generator not found")
	}

	return &genNumber, nil
}

func (r *GenerateNumberRepository) Create(genNumber *models.GenerateNumber) (*models.GenerateNumber, error) {
	err := r.db.Save(genNumber).Error
	if err != nil {
		return nil, err
	}
	return genNumber, nil
}

func (r *GenerateNumberRepository) GetLastNumber(moduleId string) (*models.GenerateNumber, error) {
	var genNumber models.GenerateNumber
	result := r.db.Where("module_id = ?", moduleId).First(&genNumber)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Number generator not found")
	}

    // Accessing the value of the tabledata column
    tableDataValue := genNumber.TableData
    fieldDataValue := genNumber.FieldData
    fmt.Println("Table Data Value:", tableDataValue, ", Field Data Value: ", fieldDataValue)

	// Now you can use the `tableDataValue` to construct a raw SQL query:
    query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s DESC LIMIT 1", fieldDataValue, tableDataValue, fieldDataValue)

    // Execute the raw query
    rows, err := r.db.Raw(query).Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	// Print the rows
	var codeNumber string // Change these types to match your table's columns
    for rows.Next() {
        // Scan the row into the variables
        err := rows.Scan(&codeNumber)
        if err != nil {
            return nil, err
        }

        // Print the row values
        fmt.Printf("codeNumber: %s\n", codeNumber)
    }

    // Check for errors from iterating over rows
    if err := rows.Err(); err != nil {
        return nil, err
    }
	// Order the results in descending order of inventory_id and select the first record.
	// result := r.db.Order("inventory_code desc").First(&lastInventory)

	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	return &genNumber, nil
}

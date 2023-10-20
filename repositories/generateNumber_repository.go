package repositories

import (
	"fmt"
	"go_play/models"
	"regexp"
	"strconv"

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

func (r *GenerateNumberRepository) GetLastNumber(code string, genNumber *models.GenerateNumber) (string, error) {
	var codeNumber string
	tableDataValue := genNumber.TableData
	fieldDataValue := genNumber.FieldData
	defaultValue := genNumber.DefaultNo
	fmt.Println("Table Data Value:", tableDataValue, ", Field Data Value: ", fieldDataValue)

	// Regular expression pattern to match the number within [AUTOx]
	autoPattern := `\[AUTO(\d+)\]`
	re := regexp.MustCompile(autoPattern)
	matches := re.FindStringSubmatch(defaultValue)
	var extractedNumber = 0
	if len(matches) == 2 {
		extracted, err := strconv.Atoi(matches[1]) // Convert the extracted string to an integer
		if err != nil {
			return "", err
		}
		extractedNumber = extracted
		fmt.Printf("Input: %s, Extracted Number: %0d\n", defaultValue, extractedNumber)
	}

	likeParam := code[:len(code)-extractedNumber] + "%"
	fmt.Println("Like query :", likeParam)

	// Now you can use the `tableDataValue` to construct a raw SQL query:
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIKE ? ORDER BY %s DESC LIMIT 1", fieldDataValue, tableDataValue, fieldDataValue, fieldDataValue)
	// Execute the raw query
	rows, err := r.db.Raw(query, likeParam).Rows()
	if err != nil {
		return "Error query", err
	}
	defer rows.Close()

	for rows.Next() {
		// Scan the row into the variables
		err := rows.Scan(&codeNumber)
		if err != nil {
			return "Error scan", err
		}
	}
	if len(codeNumber) == 0 || codeNumber == "" || codeNumber == "[]" {
		codeNumber = code
	}
	fmt.Printf("code number : %s\n", codeNumber)
	return codeNumber, nil
}

func (r *GenerateNumberRepository) GetRecord(moduleId string) (*models.GenerateNumber, error) {
	var genNumber models.GenerateNumber
	result := r.db.Where("module_id = ?", moduleId).First(&genNumber)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Generate number record not found")
	}

	return &genNumber, nil
}

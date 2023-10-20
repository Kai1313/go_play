package services

import (
	"fmt"
	"go_play/models"
	"go_play/repositories"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CreateGenerateNumber(genNumber *models.GenerateNumber, repository repositories.GenerateNumberRepository) (*models.GenerateNumberJson, error) {
	genNumberDB, err := repository.Create(genNumber)
	if err != nil {
		return nil, err
	}

	genNumberId := int(genNumberDB.GenerateId)
	genNumberJson := models.GenerateNumberJson{
		GenerateId: &genNumberId,
		ModuleId:   genNumberDB.ModuleId,
		DefaultNo:  genNumberDB.DefaultNo,
		TableData:  genNumberDB.TableData,
		FieldData:  genNumberDB.FieldData,
	}

	return &genNumberJson, nil
}

func GetGenerateNumber(id int, repository repositories.GenerateNumberRepository) (*models.GenerateNumberJson, error) {
	genNumberDB, err := repository.Show(id)
	if err != nil {
		return nil, err
	}

	genNumberId := int(genNumberDB.GenerateId)
	genNumberJson := models.GenerateNumberJson{
		GenerateId: &genNumberId,
		ModuleId:   genNumberDB.ModuleId,
		DefaultNo:  genNumberDB.DefaultNo,
		TableData:  genNumberDB.TableData,
		FieldData:  genNumberDB.FieldData,
	}

	return &genNumberJson, nil
}

func GetAllGenerateNumber(repository repositories.GenerateNumberRepository) ([]models.GenerateNumberJson, error) {
	genNumberList, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	var genNumberJsonList []models.GenerateNumberJson
	for _, genNumberDB := range genNumberList {
		genNumberId := int(genNumberDB.GenerateId)
		genNumberJson := models.GenerateNumberJson{
			GenerateId: &genNumberId,
			ModuleId:   genNumberDB.ModuleId,
			DefaultNo:  genNumberDB.DefaultNo,
			TableData:  genNumberDB.TableData,
			FieldData:  genNumberDB.FieldData,
		}
		genNumberJsonList = append(genNumberJsonList, genNumberJson)
	}

	return genNumberJsonList, nil
}

func GetLastNumber(moduleId string, repository repositories.GenerateNumberRepository) (*string, error) {
	genNumberDB, err := repository.GetRecord(moduleId)
	if err != nil {
		return nil, err
	}

	numberPattern := genNumberDB.DefaultNo
	fmt.Printf("codeNumber: %s\n", numberPattern)
	// Increment the code
	newCode, err := AutoGenerateNumber(numberPattern)
	if err != nil {
		return nil, err
	}
	fmt.Printf("newCode: %s\n", newCode)

	getLast, err := repository.GetLastNumber(newCode, genNumberDB)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Last number: %s\n", getLast)

	getNewCode, err := IncrementGenerateNumber(getLast, numberPattern)
	fmt.Printf("New Code : %s\n", getNewCode)

	return &getNewCode, nil
}

func AutoGenerateNumber(pattern string) (string, error) {
	patternParts := strings.Split(pattern, ".")
	if len(patternParts) == 0 {
		return "", fmt.Errorf("Pattern is null")
	}

	var resultParts []string
	for _, part := range patternParts {
		replacedPart, err := AutoGenerateReplacer(part)
		if err != nil {
			return "", err
		}
		resultParts = append(resultParts, replacedPart)
	}
	result := strings.Join(resultParts, ".")

	return result, nil
}

func AutoGenerateReplacer(code string) (string, error) {
	// Regular expression pattern to match "[AUTO" followed by digits and "]"
	autoPattern := `\[AUTO(\d+)\]`
	re := regexp.MustCompile(autoPattern)
	switch code {
	case "[MM]":
		currentMonth := time.Now().Format("01")
		return currentMonth, nil
	case "[YY]":
		currentYear := time.Now().Format("06")
		return currentYear, nil
	}

	if re.MatchString(code) {
		// Define a replacer function to replace [AUTOx] with zeros
		replacer := func(match string) string {
			digits := re.FindStringSubmatch(match)
			if len(digits) == 2 {
				n, err := strconv.Atoi(digits[1])
				if err != nil {
					return match
				}
				return fmt.Sprintf("%0*d", n, 0)
			}
			return match
		}

		// Replace [AUTOx] with zeros
		return re.ReplaceAllStringFunc(code, replacer), nil
	}
	fmt.Printf(code)

	return code, nil
}

func IncrementGenerateNumber(code string, pattern string) (string, error) {
	// Parse the code to extract the numeric part
	codeParts := strings.Split(code, ".")
	if len(codeParts) < 2 {
		return "", fmt.Errorf("Invalid code format")
	}

	numericPart := codeParts[len(codeParts)-1]
	numericValue, err := strconv.Atoi(numericPart)
	if err != nil {
		return "", err
	}

	// Increment the numeric value
	numericValue++

	// Regular expression pattern to match "[AUTO" followed by digits and "]"
	patternParts := strings.Split(pattern, ".")
	autoPattern := `\[AUTO(\d+)\]`
	re := regexp.MustCompile(autoPattern)
	matches := re.FindStringSubmatch(patternParts[len(patternParts)-1])
	var extractedNumber int
	if len(matches) == 2 {
		extracted, err := strconv.Atoi(matches[1]) // Convert the extracted string to an integer
		if err != nil {
			return "", err
		}
		extractedNumber = extracted
	}

	// Format the new code
	// newCode := fmt.Sprintf("%0*d", extractedNumber, numericValue)
	newCode := strings.Replace(code, numericPart, fmt.Sprintf("%0*d", extractedNumber, numericValue), 1)
	return newCode, nil
}

package services

import (
	"go_play/models"
	"go_play/repositories"
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

func GetLastNumber(moduleId string, repository repositories.GenerateNumberRepository) (*models.GenerateNumberJson, error) {
	genNumberDB, err := repository.GetLastNumber(moduleId)
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
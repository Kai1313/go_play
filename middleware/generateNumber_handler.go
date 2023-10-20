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

type GenerateNumberHandler struct {
	repository *repositories.GenerateNumberRepository
}

func NewGenerateNumberRepository(repository *repositories.GenerateNumberRepository) *GenerateNumberHandler {
	return &GenerateNumberHandler{
		repository: repository,
	}
}

func (h *GenerateNumberHandler) ShowGenerateNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int. %v", err)
	}

	genNumberJson, err := services.GetGenerateNumber(int(id), *h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
	}

	res := Response{
		Success: true,
		Data:    genNumberJson,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *GenerateNumberHandler) CreateGenerateNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var genNumber models.GenerateNumberJson
	err := json.NewDecoder(r.Body).Decode(&genNumber)
	if err != nil {
		log.Fatalf("Unable to decode the requested body. %v", err)
	}

	genNumberDB := models.GenerateNumber{
		ModuleId: genNumber.ModuleId,
		DefaultNo: genNumber.DefaultNo,
		TableData: genNumber.TableData,
		FieldData: genNumber.FieldData,
	}

	genNumberJson, err := services.CreateGenerateNumber(&genNumberDB, *h.repository)
	if err != nil {
		res := Response{
			Success: false,
			Message: err.Error(),
		}

		json.NewEncoder(w).Encode(res)
	}

	res := Response{
		Success: true,
		Data:    genNumberJson,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *GenerateNumberHandler) GetAllGenerateNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	genNumberJsonList, err := services.GetAllGenerateNumber(*h.repository)
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
		Data:    genNumberJsonList,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *GenerateNumberHandler) ShowLastGenerateNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	moduleId := mux.Vars(r)["id"]

	genNumberJson, err := services.GetLastNumber(moduleId, *h.repository)
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
		Data:    genNumberJson,
	}
	json.NewEncoder(w).Encode(res)
}
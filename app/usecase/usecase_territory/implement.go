package usecase_territory

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
)

type UseCaseImpl struct {
	TerritoryRepository repository.TerritoryRepository
	Validate            *validator.Validate
}

func NewUseCase(
	territoryRepo repository.TerritoryRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		Validate:            validate,
		TerritoryRepository: territoryRepo,
	}
}

func (controller *UseCaseImpl) GetProvince(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	dataResponse := controller.TerritoryRepository.FindSpesificDataProvince(r.Context(), map[string]interface{}{})

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetCity(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	dataResponse := controller.TerritoryRepository.FindSpesificDataCity(r.Context(), map[string]interface{}{
		"province_id": vars.Get("province_id"),
	})

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetSubdistrict(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	dataResponse := controller.TerritoryRepository.FindSpesificDataSubdistrict(r.Context(), map[string]interface{}{
		"regency_id": vars.Get("regency_id"),
	})

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetVillage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	dataResponse := controller.TerritoryRepository.FindSpesificDataVillage(r.Context(), map[string]interface{}{
		"district_id": vars.Get("district_id"),
	})

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

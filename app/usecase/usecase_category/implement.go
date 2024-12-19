package usecase_category

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
)

type UseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewUseCase(categoryRepo repository.CategoryRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		CategoryRepository: categoryRepo,
	}
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["categoryId"]
	dataResponse, err := controller.CategoryRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	dataResponse := controller.CategoryRepository.FindAll(r.Context())
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

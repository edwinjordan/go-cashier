package usecase_major

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
)

type UseCaseImpl struct {
	MajorRepository repository.MajorRepository
	Validate        *validator.Validate
}

func NewUseCase(majorRepo repository.MajorRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:        validate,
		MajorRepository: majorRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Major{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	dataResponse := controller.MajorRepository.Create(r.Context(), dataRequest)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	dataRequest := entity.Major{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	_, err = controller.MajorRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataRequest.MajorId = id
	dataResponse := controller.MajorRepository.Update(r.Context(), dataRequest, id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	_, err := controller.MajorRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.MajorRepository.Delete(r.Context(), id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["majorId"]
	dataResponse, err := controller.MajorRepository.FindById(r.Context(), id)
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

	dataResponse := controller.MajorRepository.FindAll(r.Context())
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

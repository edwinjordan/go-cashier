package usecase_product

import (
	"fmt"
	"net/http"
	"strconv"

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
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewUseCase(ProductRepo repository.ProductRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:          validate,
		ProductRepository: ProductRepo,
	}
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["productId"]
	dataResponse, err := controller.ProductRepository.FindById(r.Context(), id)
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
	query := r.URL.Query()
	Qlimit := query.Get("limit")
	Qoffset := query.Get("offset")
	search := query.Get("search")

	if Qlimit == "" {
		Qlimit = "10"
	}

	if Qoffset == "" {
		Qoffset = "0"
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset, _ := strconv.Atoi(Qoffset)

	nextOffset := limit + offset

	conf := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"search": search,
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")
	where := entity.Product{
		ProductCategoryId: query.Get("category_id"),
	}
	dataResponse := controller.ProductRepository.FindAll(r.Context(), where, conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

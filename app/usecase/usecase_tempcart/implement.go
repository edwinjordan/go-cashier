package usecase_tempcart

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
	TempCartRepository repository.TempCartRepository
	VarianRepository   repository.VarianRepository
	Validate           *validator.Validate
}

func NewUseCase(tempcartRepo repository.TempCartRepository, productRepo repository.VarianRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		TempCartRepository: tempcartRepo,
		VarianRepository:   productRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.TempCart{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	/* cek apakah masih ada sisa di product itu */
	varian, _ := controller.VarianRepository.FindById(r.Context(), dataRequest.TempCartProductVarianId)
	if varian.ProductVarianQtyLeft == 0 {
		panic(exceptions.NewBadRequestError("Tidak dapat menambah jumlah lagi karena sudah habis dipesan"))
	}

	/* cek apakah data sudah ada sebelumnya */
	tempCart := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: dataRequest.TempCartProductVarianId,
		TempCartUserId:          dataRequest.TempCartUserId,
	})

	if tempCart != nil {
		controller.TempCartRepository.Delete(r.Context(), tempCart[0].TempCartId)
		dataRequest.TempCartQty = dataRequest.TempCartQty + tempCart[0].TempCartQty
	}
	dataRequest.TempCartOrderId = ""
	controller.TempCartRepository.Create(r.Context(), dataRequest)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    varian,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productVarianId := vars["productVarianId"]
	userId := vars["userId"]
	dataRequest := entity.TempCart{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* ambil data temp cart */
	cart := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: productVarianId,
		TempCartUserId:          userId,
	})

	if cart == nil {
		panic(exceptions.NewNotFoundError("Data tidak ditemukan"))
	}

	productVarian, _ := controller.VarianRepository.FindById(r.Context(), productVarianId)
	/* check apakah masih ada quantity tersisa */
	if (productVarian.ProductVarianQtyLeft + cart[0].TempCartQty) < dataRequest.TempCartQty {
		panic(exceptions.NewBadRequestError("Tidak dapat mengubah data karena jumlah tersisa kurang"))
	}

	/* update data temporary */
	dataRequest.TempCartId = cart[0].TempCartId
	dataRequest.TempCartProductVarianId = cart[0].TempCartProductVarianId
	dataRequest.TempCartUserId = cart[0].TempCartUserId
	controller.TempCartRepository.Update(r.Context(), dataRequest, cart[0].TempCartId)

	/* ambil lagi data temporary terbaru */
	data, _ := controller.VarianRepository.FindById(r.Context(), productVarianId)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    data,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	producVarianId := vars["productVarianId"]
	userId := vars["userId"]

	data := controller.TempCartRepository.FindSpesificData(r.Context(), entity.TempCart{
		TempCartProductVarianId: producVarianId,
		TempCartUserId:          userId,
	})

	if data != nil {
		controller.TempCartRepository.Delete(r.Context(), data[0].TempCartId)
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

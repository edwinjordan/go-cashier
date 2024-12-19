package usecase_customer_address

import (
	"html"
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
	CustomerAddressRepository repository.CustomerAddressRepository
	Validate                  *validator.Validate
}

func NewUseCase(addressRepo repository.CustomerAddressRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:                  validate,
		CustomerAddressRepository: addressRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.CustomerAddress{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* check apakah ada data sebelumnya */
	checkAddress := controller.CustomerAddressRepository.FindSpesificData(r.Context(), entity.CustomerAddress{
		AddressCustomerId: dataRequest.AddressCustomerId,
	})

	if checkAddress != nil {
		/* jika lebih dari 6 maka tidak bisa menambah lagi */
		if len(checkAddress) == 6 {
			panic(exceptions.NewBadRequestError("Tidak bisa menambah alamat lagi"))
		}

		/* check apakah data baru diset sebagai alamat utama */
		if dataRequest.AddressMain == 1 {
			/* jika iya maka ubah semua menjadi 0 */
			controller.CustomerAddressRepository.Update(r.Context(), entity.CustomerAddress{
				AddressMain: 0,
			}, "address_main", entity.CustomerAddress{
				AddressCustomerId: dataRequest.AddressCustomerId,
			})
		}
	} else {
		/* jika belum ada data sama sekali maka otomatis buat menjadi alamat utama */
		dataRequest.AddressMain = 1
	}
	dataRequest.AddressName = html.EscapeString(dataRequest.AddressName)
	dataRequest.AddressText = html.EscapeString(dataRequest.AddressText)
	dataResponse := controller.CustomerAddressRepository.Create(r.Context(), dataRequest)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["addressId"]
	dataRequest := entity.CustomerAddress{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	address, err := controller.CustomerAddressRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	if dataRequest.AddressMain == 1 {
		/* jika iya maka ubah semua menjadi 0 */
		controller.CustomerAddressRepository.Update(r.Context(), entity.CustomerAddress{
			AddressMain: 0,
		}, "address_main", entity.CustomerAddress{
			AddressCustomerId: dataRequest.AddressCustomerId,
		})
	}
	dataRequest.AddressId = id
	dataRequest.AddressCreateAt = address.AddressCreateAt
	dataRequest.AddressName = html.EscapeString(dataRequest.AddressName)
	dataRequest.AddressText = html.EscapeString(dataRequest.AddressText)
	dataResponse := controller.CustomerAddressRepository.Update(r.Context(), dataRequest, "*", entity.CustomerAddress{
		AddressId: id,
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["addressId"]
	address, err := controller.CustomerAddressRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	/* check apakah alamat utama yang dihapus */
	if address.AddressMain == 1 {
		panic(exceptions.NewBadRequestError("Tidak dapat menghapus alamat utama"))
	}
	controller.CustomerAddressRepository.Delete(r.Context(), id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["addressId"]
	dataResponse, err := controller.CustomerAddressRepository.FindById(r.Context(), id)
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
	vars := r.URL.Query()
	dataResponse := controller.CustomerAddressRepository.FindSpesificData(r.Context(), entity.CustomerAddress{
		AddressCustomerId: vars.Get("customer_id"),
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

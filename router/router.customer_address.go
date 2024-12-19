package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_customer_address"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_address_repository"
	"gorm.io/gorm"
)

func CustomerAddressRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	addressRepository := customer_address_repository.New(db)
	addressController := usecase_customer_address.NewUseCase(addressRepository, validate)

	router.HandleFunc("/api/address", addressController.FindAll).Methods("GET")
	router.HandleFunc("/api/address/{addressId}", addressController.FindById).Methods("GET")
	router.HandleFunc("/api/address", addressController.Create).Methods("POST")
	router.HandleFunc("/api/address/{addressId}", addressController.Update).Methods("PUT")
	router.HandleFunc("/api/address/{addressId}", addressController.Delete).Methods("DELETE")

}

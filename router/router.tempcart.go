package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_tempcart"
	"github.com/jolebo/e-canteen-cashier-api/repository/tempcart_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/varian_repository"
	"gorm.io/gorm"
)

func TempCartRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	tempCartRepo := tempcart_repository.New(db)
	varianRepository := varian_repository.New(db)
	tempCartController := usecase_tempcart.NewUseCase(tempCartRepo, varianRepository, validate)

	router.HandleFunc("/api/tempcart", tempCartController.Create).Methods("POST")
	router.HandleFunc("/api/tempcart/{productVarianId}/{userId}", tempCartController.Update).Methods("PUT")
	router.HandleFunc("/api/tempcart/{productVarianId}/{userId}", tempCartController.Delete).Methods("DELETE")
}

package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_varian"
	"github.com/jolebo/e-canteen-cashier-api/repository/varian_repository"
	"gorm.io/gorm"
)

func VarianRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	varianRepository := varian_repository.New(db)
	varianController := usecase_varian.NewUseCase(varianRepository, validate)
	router.HandleFunc("/api/varian/{varianId}", varianController.FindById).Methods("GET")
	router.HandleFunc("/api/varian", varianController.FindAll).Methods("GET")
}

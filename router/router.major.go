package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_major"
	"github.com/jolebo/e-canteen-cashier-api/repository/major_repository"
	"gorm.io/gorm"
)

func MajorRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	majorRepository := major_repository.New(db)
	majorController := usecase_major.NewUseCase(majorRepository, validate)
	router.HandleFunc("/api/major/{majorId}", majorController.FindById).Methods("GET")
	router.HandleFunc("/api/major", majorController.FindAll).Methods("GET")
}

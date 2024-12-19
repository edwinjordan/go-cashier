package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_category"
	"github.com/jolebo/e-canteen-cashier-api/repository/category_repository"
	"gorm.io/gorm"
)

func CategoryRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	categoryRepository := category_repository.New(db)
	categoryController := usecase_category.NewUseCase(categoryRepository, validate)
	router.HandleFunc("/api/category", categoryController.FindAll).Methods("GET")
	router.HandleFunc("/api/category/{categoryId}", categoryController.FindAll).Methods("GET")
}

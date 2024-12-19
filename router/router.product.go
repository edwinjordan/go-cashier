package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_product"
	"github.com/jolebo/e-canteen-cashier-api/repository/product_repository"
	"gorm.io/gorm"
)

func ProductRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	productRepository := product_repository.New(db)
	productController := usecase_product.NewUseCase(productRepository, validate)
	router.HandleFunc("/api/products", productController.FindAll).Methods("GET")
}

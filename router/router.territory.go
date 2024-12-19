package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_territory"
	"github.com/jolebo/e-canteen-cashier-api/repository/territory_repository"
	"gorm.io/gorm"
)

func TerritoryRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	repositoryRepository := territory_repository.New(db)
	repositoryController := usecase_territory.NewUseCase(repositoryRepository, validate)
	router.HandleFunc("/api/province", repositoryController.GetProvince).Methods("GET")
	router.HandleFunc("/api/city", repositoryController.GetCity).Methods("GET")
	router.HandleFunc("/api/subdistrict", repositoryController.GetSubdistrict).Methods("GET")
	router.HandleFunc("/api/village", repositoryController.GetVillage).Methods("GET")
}

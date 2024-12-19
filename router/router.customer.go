package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_customer"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/otp_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/tempcart_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/user_repository"
	"gorm.io/gorm"
)

func CustomerRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	customerRepository := customer_repository.New(db)
	otpRepository := otp_repository.New(db)
	userLogRepository := user_repository.NewLog(db)
	tempCartRepo := tempcart_repository.New(db)
	customerController := usecase_customer.NewUseCase(customerRepository, otpRepository, tempCartRepo, userLogRepository, validate)
	router.HandleFunc("/api/customer/login", customerController.DoLogin).Methods("POST")
	router.HandleFunc("/api/customer/addLog", customerController.AddLog).Methods("POST")
	router.HandleFunc("/api/customer/logout", customerController.DoLogout).Methods("POST")
	router.HandleFunc("/api/customer/change_password", customerController.ChangePassword).Methods("POST")

	router.HandleFunc("/api/customer", customerController.FindAll).Methods("GET")
	router.HandleFunc("/api/customer/{customerId}", customerController.FindById).Methods("GET")
	router.HandleFunc("/api/customer", customerController.Register).Methods("POST")
	router.HandleFunc("/api/customer/{customerId}", customerController.Update).Methods("PUT")
	router.HandleFunc("/api/customer/{customerId}", customerController.Delete).Methods("DELETE")
	router.HandleFunc("/api/customer/verifyOtp", customerController.VerifyOtp).Methods("POST")
	router.HandleFunc("/api/customer/sentOTPResetPassword", customerController.SendOTPResetPassword).Methods("POST")

}

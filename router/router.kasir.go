package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_transaction"
	"github.com/jolebo/e-canteen-cashier-api/app/usecase/usecase_user"
	"github.com/jolebo/e-canteen-cashier-api/repository/order_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/stock_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/tempcart_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/transaction_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/user_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/varian_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/version_repository"
	"gorm.io/gorm"
)

func KasirRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	tempCartRepo := tempcart_repository.New(db)
	userRepository := user_repository.New(db)
	userLogRepository := user_repository.NewLog(db)
	versionRepository := version_repository.New(db)
	userController := usecase_user.NewUseCase(userRepository, userLogRepository, tempCartRepo, versionRepository, validate)
	router.HandleFunc("/api/kasir/login", userController.DoLogin).Methods("POST")
	router.HandleFunc("/api/kasir/logout", userController.DoLogout).Methods("PUT")
	router.HandleFunc("/api/kasir/version", userController.GetVersionAdmin).Methods("GET")
	router.HandleFunc("/api/shop/version", userController.GetVersionShop).Methods("GET")
	router.HandleFunc("/api/check_maintenance_mode/{confCode}", userController.CheckMaintenanceMode).Methods("GET")

	/* tansaction */
	transRepo := transaction_repository.NewTrans(db)
	transDetailRepo := transaction_repository.NewTransDetail(db)
	stockBoothRepo := stock_repository.NewBooth(db)
	varianRepo := varian_repository.New(db)
	orderRepository := order_repository.NewOrder(db)
	orderDetailRepository := order_repository.NewOrderDetail(db)
	transController := usecase_transaction.NewUseCase(transRepo, transDetailRepo, tempCartRepo, stockBoothRepo, varianRepo, orderRepository, orderDetailRepository, validate)
	router.HandleFunc("/api/kasir/transaction", transController.Create).Methods("POST")
	router.HandleFunc("/api/transaction", transController.FindAll).Methods("GET")
	router.HandleFunc("/api/kasir/transaction/{transId}", transController.FindById).Methods("GET")
	router.HandleFunc("/api/kasir/transaction_detail", transController.GetTransDetail).Methods("GET")
	router.HandleFunc("/api/kasir/transaction_summary", transController.GetTransactionSummary).Methods("GET")
}

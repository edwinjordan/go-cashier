package main

import (
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/middleware"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"github.com/jolebo/e-canteen-cashier-api/pkg/mysql"
	"github.com/jolebo/e-canteen-cashier-api/router"
	"github.com/rs/cors"
)

func main() {
	validate := validator.New()
	db := mysql.DBConnectGorm()
	route := mux.NewRouter()

	/* setting cors */
	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders: []string{
			"*",
		},
	})
	/* load middleware */
	// route.Use(middleware.Recovery)
	route.Use(middleware.Authentication)

	/* load router */
	router.KasirRouter(db, validate, route)
	router.VarianRouter(db, validate, route)
	router.TempCartRouter(db, validate, route)
	router.ProductRouter(db, validate, route)
	router.CategoryRouter(db, validate, route)
	router.CustomerRouter(db, validate, route)
	router.CustomerAddressRouter(db, validate, route)
	router.MajorRouter(db, validate, route)
	router.TerritoryRouter(db, validate, route)
	router.OrderRouter(db, validate, route)

	server := http.Server{
		Addr:    config.GetEnv("HOST_ADDR"),
		Handler: corsOpt.Handler(route),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}

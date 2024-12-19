package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.GetEnv("DB_USERNAME"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"), config.GetEnv("DB_NAME"))
	db, err := sql.Open(config.GetEnv("DB_DRIVER"), connectionString)
	helpers.PanicIfError(err)

	idleCon, _ := strconv.Atoi(config.GetEnv("DB_MAXIDLECON"))
	openCon, _ := strconv.Atoi(config.GetEnv("DB_MAXOPENCON"))

	db.SetMaxIdleConns(idleCon)
	db.SetMaxOpenConns(openCon)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func DBConnectGorm() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.GetEnv("DB_USERNAME"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"), config.GetEnv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	helpers.PanicIfError(err)

	return db
}

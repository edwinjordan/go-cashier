package helpers

import (
	"fmt"

	"github.com/jolebo/e-canteen-cashier-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetFCMToken(userId string) []string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.GetEnv("DB_USERNAME"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"), config.GetEnv("DB_NAME"))
	db, _ := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	data := []map[string]interface{}{}
	db.Raw("SELECT log_user_token FROM tb_user_token WHERE log_user_user_id = ? AND log_user_logout_date IS NULL", userId).Find(&data)

	res := []string{}
	for _, v := range data {
		res = append(res, v["log_user_token"].(string))
	}
	return res
}

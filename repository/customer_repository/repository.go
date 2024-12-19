package customer_repository

import (
	"context"
	"errors"
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{
		DB: db,
	}
}

func (repo *CustomerRepositoryImpl) Create(ctx context.Context, customer entity.Customer) entity.CustomerResponse {
	customerData := &Customer{}
	customerData = customerData.FromEntity(&customer)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&customerData).Error
	helpers.PanicIfError(err)

	return *customerData.ToEntity()
}

func (repo *CustomerRepositoryImpl) Update(ctx context.Context, selectField interface{}, customer entity.Customer, customerId string) entity.CustomerResponse {
	customerData := &Customer{}
	customerData = customerData.FromEntity(&customer)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("customer_id = ?", customerId).Select(selectField).Updates(&customerData).Error
	helpers.PanicIfError(err)
	return *customerData.ToEntity()
}

func (repo *CustomerRepositoryImpl) Delete(ctx context.Context, customerId string) {
	customer := &Customer{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("customer_id = ?", customerId).Delete(&customer).Error
	helpers.PanicIfError(err)
}

func (repo *CustomerRepositoryImpl) FindById(ctx context.Context, customerId string) (entity.CustomerResponse, error) {
	customerData := &Customer{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("customer_id = ?", customerId).
		Preload("Address").
		Preload("Major").
		First(&customerData).Error
	if err != nil {
		return *customerData.ToEntity(), errors.New("data pelanggan tidak ditemukan")
	}
	return *customerData.ToEntity(), nil
}

func (repo *CustomerRepositoryImpl) FindAll(ctx context.Context, config map[string]interface{}) []entity.CustomerResponse {
	customer := []Customer{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	var tempData []entity.CustomerResponse

	whereLike := ""
	if config["search"].(string) != "" {
		search := html.EscapeString(config["search"].(string))
		whereLike = "(customer_name LIKE '%" + search + "%' OR customer_phonenumber LIKE '%" + search + "%')"
	}

	/* ambil data customer yang dipilih */
	whereNot := ""
	if whereLike == "" && config["customer"].(string) != "" {
		whereNot = config["customer"].(string)
		if config["offset"].(int) == 0 {
			cust, err := repo.FindById(ctx, config["customer"].(string))
			if err == nil {
				tempData = append(tempData, cust)
			}
		}
	}

	err := tx.WithContext(ctx).
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Preload("Major").
		Preload("Address").
		Where("customer_id != ? ", whereNot).
		Where(whereLike).
		Find(&customer).Error
	helpers.PanicIfError(err)

	for _, v := range customer {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *CustomerRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Customer) []entity.CustomerResponse {
	customer := []Customer{}
	customerWhere := &Customer{}
	customerWhere = customerWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(customerWhere).Preload("Address").Preload("Major").Find(&customer).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerResponse
	for _, v := range customer {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}

func (repo *CustomerRepositoryImpl) GenCustCode(ctx context.Context) string {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	invoice := map[string]interface{}{}

	month := fmt.Sprint(int(time.Now().Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	year := fmt.Sprint(int(time.Now().Year()) % 1e2)

	date := month + year

	tx.WithContext(ctx).Table("tb_customer").Select("IFNULL(customer_code,'') customer_code").Where("customer_code LIKE ?", "%"+date+"%").Order("customer_code DESC").Find(invoice)
	inv := ""
	if invoice["customer_code"] == nil {
		inv = "ESC-" + date + "-000"
	} else {
		inv = invoice["customer_code"].(string)
	}
	sort := inv[len(inv)-3:]
	newInv := inv[:len(inv)-3]
	str, _ := strconv.Atoi(sort)
	str += 1
	if str < 10 {
		sort = "00" + fmt.Sprint(str)
	} else if str < 100 {
		sort = "0" + fmt.Sprint(str)
	} else {
		sort = fmt.Sprint(str)
	}

	return newInv + sort
}

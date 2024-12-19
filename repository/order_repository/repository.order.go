package order_repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) repository.CustomerOrderRepository {
	return &CustomerOrderRepositoryImpl{
		DB: db,
	}
}

func (repo *CustomerOrderRepositoryImpl) Create(ctx context.Context, order entity.CustomerOrder) entity.CustomerOrder {
	orderData := &CustomerOrder{}
	orderData = orderData.FromEntity(&order)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&orderData).Error
	helpers.PanicIfError(err)

	return *orderData.ToEntity()
}

func (repo *CustomerOrderRepositoryImpl) Update(ctx context.Context, order entity.CustomerOrder, selectField interface{}, where entity.CustomerOrder) entity.CustomerOrder {
	orderData := &CustomerOrder{}
	orderData = orderData.FromEntity(&order)

	whereData := &CustomerOrder{}
	whereData = whereData.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Select(selectField).Save(&orderData).Error
	helpers.PanicIfError(err)
	return *orderData.ToEntity()
}

func (repo *CustomerOrderRepositoryImpl) Delete(ctx context.Context, where entity.CustomerOrder) {
	order := &CustomerOrder{}
	whereData := order.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Delete(&order).Error
	helpers.PanicIfError(err)
}

func (repo *CustomerOrderRepositoryImpl) FindById(ctx context.Context, orderId string) (entity.CustomerOrder, error) {
	orderData := &CustomerOrder{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("order_id = ?", orderId).
		Preload("Address").
		Preload("Customer").
		Preload("OrderDetail").
		First(&orderData).Error
	if err != nil {
		return *orderData.ToEntity(), errors.New("data order tidak ditemukan")
	}
	return *orderData.ToEntity(), nil
}

func (repo *CustomerOrderRepositoryImpl) FindAll(ctx context.Context) []entity.CustomerOrder {
	order := []CustomerOrder{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrder
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *CustomerOrderRepositoryImpl) FindSpesificData(ctx context.Context, where entity.CustomerOrder, config map[string]interface{}) []entity.CustomerOrder {
	order := []CustomerOrder{}
	orderWhere := &CustomerOrder{}
	orderWhere = orderWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderString := "order_create_at DESC"

	if orderWhere.OrderStatus == 1 || orderWhere.OrderStatus == 3 {
		orderString = "order_finished_datetime DESC"
	} else if orderWhere.OrderStatus == 2 {
		orderString = "order_processed_datetime DESC"
	}

	err := tx.WithContext(ctx).
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Order(orderString).
		Where(orderWhere).
		Preload("Address").
		Preload("Customer").
		Preload("OrderDetail").
		Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrder
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}

func (repo *CustomerOrderRepositoryImpl) GenInvoice(ctx context.Context) string {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	invoice := map[string]interface{}{}

	month := fmt.Sprint(int(time.Now().Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := fmt.Sprint(time.Now().Day())
	if len(day) == 1 {
		day = "0" + day
	}
	year := fmt.Sprint(int(time.Now().Year()) % 1e2)

	date := day + month + year

	tx.WithContext(ctx).Table("tb_customer_order").Select("IFNULL(order_inv_number,'') order_inv_number").Where("order_inv_number LIKE ?", "%"+date+"%").Order("order_inv_number DESC").Find(invoice)
	inv := ""
	if invoice["order_inv_number"] == nil {
		inv = "ORN-" + date + "-000"
	} else {
		inv = invoice["order_inv_number"].(string)
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

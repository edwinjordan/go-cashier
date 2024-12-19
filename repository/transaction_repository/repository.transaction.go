package transaction_repository

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

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTrans(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

func (repo *TransactionRepositoryImpl) Create(ctx context.Context, transaksi entity.Transaction) entity.Transaction {
	transData := &Transaction{}
	transData = transData.FromEntity(&transaksi)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&transData).Error
	helpers.PanicIfError(err)

	return *transData.ToEntity()
}

func (repo *TransactionRepositoryImpl) Update(ctx context.Context, transaksi entity.Transaction, transaksiId string) entity.Transaction {
	transData := &Transaction{}
	transData = transData.FromEntity(&transaksi)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("trans_id = ?", transaksiId).Save(&transData).Error
	helpers.PanicIfError(err)
	return *transData.ToEntity()
}

func (repo *TransactionRepositoryImpl) Delete(ctx context.Context, transaksiId string) {
	transaksi := &Transaction{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("trans_id = ?", transaksiId).Delete(&transaksi).Error
	helpers.PanicIfError(err)
}

func (repo *TransactionRepositoryImpl) FindById(ctx context.Context, transaksiId string) (entity.Transaction, error) {
	transData := &Transaction{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("trans_id = ?", transaksiId).
		Preload("TransDetail").
		First(&transData).Error
	if err != nil {
		return *transData.ToEntity(), errors.New("data hadiah tidak ditemukan")
	}
	return *transData.ToEntity(), nil
}

func (repo *TransactionRepositoryImpl) FindAll(ctx context.Context) []entity.Transaction {
	transaksi := []Transaction{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Joins("left join tb_customer on customer_id = trans_customer_id").
		Preload("User").
		Preload("Customer").
		Preload("TransDetail").
		Find(&transaksi).Error
	helpers.PanicIfError(err)

	var tempData []entity.Transaction
	for _, v := range transaksi {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *TransactionRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Transaction, config map[string]interface{}) []entity.Transaction {
	transaksi := []Transaction{}
	transaksiWhere := &Transaction{}
	transaksiWhere = transaksiWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	whereString := ""
	if config["typeofDate"] == "today" || config["typeofDate"] == "yesterday" {
		whereString = "DATE_FORMAT(trans_create_at,'%Y-%m-%d') = '" + config["date"].(map[string]interface{})["startDate"].(string) + "'"
		// tx.WithContext(ctx).Where("DATE_FORMAT(trans_create_at,'%Y-%m-%d') = ?", config["date"].(map[string]interface{})["startDate"])
	}

	if config["typeofDate"] == "this_month" {
		whereString = "DATE_FORMAT(trans_create_at,'%Y-%m') = '" + config["date"].(map[string]interface{})["startDate"].(string) + "'"
		// tx.WithContext(ctx).Where("DATE_FORMAT(trans_create_at,'%Y-%m') = ?", config["date"].(map[string]interface{})["startDate"])
	}

	if config["typeofDate"] == "7" || config["typeofDate"] == "30" {
		whereString = "DATE_FORMAT(trans_create_at,'%Y-%m-%d')  BETWEEN '" + config["date"].(map[string]interface{})["startDate"].(string) + "' AND '" + config["date"].(map[string]interface{})["endDate"].(string) + "'"
		// tx.WithContext(ctx).Where("DATE_FORMAT(trans_create_at,'%Y-%m-%d') BETWEEN ? AND ?", config["date"].(map[string]interface{})["dateStart"], config["date"].(map[string]interface{})["endDate"])
	}
	err := tx.WithContext(ctx).
		Joins("left join tb_customer on customer_id = trans_customer_id").
		Where(whereString).
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Preload("Customer").
		Preload("TransDetail").
		Preload("User").
		Order("trans_invoice DESC").
		Where(transaksiWhere).
		Find(&transaksi).Error
	helpers.PanicIfError(err)

	var tempData []entity.Transaction
	for _, v := range transaksi {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *TransactionRepositoryImpl) GetTransactionSummary(ctx context.Context, where map[string]interface{}) map[string]interface{} {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	data := map[string]interface{}{}
	/* get transaction summary */
	dataSum := map[string]interface{}{}

	whereString := ""
	if where["typeofDate"] == "today" || where["typeofDate"] == "yesterday" {
		whereString = "AND (DATE_FORMAT(trans_create_at,'%Y-%m-%d') = '" + where["date"].(map[string]interface{})["startDate"].(string) + "')"
	}

	if where["typeofDate"] == "this_month" {
		whereString = "AND (DATE_FORMAT(trans_create_at,'%Y-%m') = '" + where["date"].(map[string]interface{})["startDate"].(string) + "')"
	}

	if where["typeofDate"] == "7" || where["typeofDate"] == "30" {
		whereString = "AND (DATE_FORMAT(trans_create_at,'%Y-%m-%d')  BETWEEN '" + where["date"].(map[string]interface{})["startDate"].(string) + "' AND '" + where["date"].(map[string]interface{})["endDate"].(string) + "')"
	}

	tx.WithContext(ctx).Raw("SELECT SUM( trans_subtotal ) trans_subtotal,SUM( trans_total ) trans_total, COUNT(*) total_trans, SUM( trans_discount) trans_discount FROM tb_transaction  WHERE trans_user_id = ? "+whereString, where["userId"].(string)).Find(&dataSum)

	/* get product top selling */
	dataProduct := []map[string]interface{}{}
	tx.WithContext(ctx).Raw("SELECT SUM( trans_detail_qty ) trans_detail_qty, SUM( trans_detail_subtotal ) trans_detail_subtotal, CONCAT(product_name,'(',varian_name,')') product_name FROM v_tb_trans_detail JOIN tb_transaction ON trans_id=trans_detail_parent_id WHERE 0 = 0 " + whereString + " GROUP BY product_name ORDER BY SUM(trans_detail_qty) DESC LIMIT 5").Find(&dataProduct)
	data["summary"] = dataSum
	data["topsell"] = dataProduct
	return data
}

func (repo *TransactionRepositoryImpl) GenInvoice(ctx context.Context) string {
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

	tx.WithContext(ctx).Table("tb_transaction").Select("IFNULL(trans_invoice,'') trans_invoice").Where("trans_invoice LIKE ?", "%"+date+"%").Order("trans_invoice DESC").Find(invoice)
	inv := ""
	if invoice["trans_invoice"] == nil {
		inv = "INV-" + date + "-000"
	} else {
		inv = invoice["trans_invoice"].(string)
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

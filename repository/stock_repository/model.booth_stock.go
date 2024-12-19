package stock_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type StockBooth struct {
	ProductStokId              string    `gorm:"column:product_stok_id"`
	ProductStokProductVarianId string    `gorm:"column:product_stok_product_varian_id"`
	ProductStokFirstQty        int       `gorm:"column:product_stok_first_qty"`
	ProductStokQty             int       `gorm:"column:product_stok_qty"`
	ProductStokLastQty         int       `gorm:"column:product_stok_last_qty"`
	ProductStokJenis           string    `gorm:"column:product_stok_jenis"`
	ProductStokDatetime        time.Time `gorm:"column:product_stok_datetime"`
	ProductStokPegawaiId       string    `gorm:"column:product_stok_pegawai_id"`
}

func (StockBooth) TableName() string {
	return "tb_product_stock_booth"
}

func (model *StockBooth) BeforeCreate(tx *gorm.DB) (err error) {
	model.ProductStokId = helpers.GenUUID()
	model.ProductStokDatetime = helpers.CreateDateTime()
	return
}

func (StockBooth) FromEntity(e *entity.StockBooth) *StockBooth {
	return &StockBooth{
		ProductStokId:              e.ProductStokId,
		ProductStokProductVarianId: e.ProductStokProductVarianId,
		ProductStokFirstQty:        e.ProductStokFirstQty,
		ProductStokQty:             e.ProductStokQty,
		ProductStokLastQty:         e.ProductStokLastQty,
		ProductStokJenis:           e.ProductStokJenis,
		ProductStokDatetime:        e.ProductStokDatetime,
		ProductStokPegawaiId:       e.ProductStokPegawaiId,
	}
}

func (model *StockBooth) ToEntity() *entity.StockBooth {
	modelData := &entity.StockBooth{
		ProductStokId:              model.ProductStokId,
		ProductStokProductVarianId: model.ProductStokProductVarianId,
		ProductStokFirstQty:        model.ProductStokFirstQty,
		ProductStokQty:             model.ProductStokQty,
		ProductStokLastQty:         model.ProductStokLastQty,
		ProductStokJenis:           model.ProductStokJenis,
		ProductStokDatetime:        model.ProductStokDatetime,
		ProductStokPegawaiId:       model.ProductStokPegawaiId,
	}
	return modelData
}

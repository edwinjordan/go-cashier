package pegawai_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type Pegawai struct {
	PegawaiId          string    `gorm:"column:pegawai_id"`
	PegawaiCode        string    `gorm:"column:pegawai_code"`
	PegawaiName        string    `gorm:"column:pegawai_name"`
	PegawaiGender      string    `gorm:"column:pegawai_gender"`
	PegawaiPhonenumber string    `gorm:"column:pegawai_phonenumber"`
	PegawaiCreateAt    time.Time `gorm:"column:pegawai_create_at"`
	PegawaiUpdateAt    time.Time `gorm:"column:pegawai_update_at"`
	PegawaiDeleteAt    time.Time `gorm:"column:pegawai_delete_at"`
}

func (Pegawai) TableName() string {
	return "ms_pegawai"
}

func (Pegawai) FromEntity(e *entity.Pegawai) *Pegawai {
	return &Pegawai{
		PegawaiId:          e.PegawaiId,
		PegawaiCode:        e.PegawaiCode,
		PegawaiName:        e.PegawaiName,
		PegawaiGender:      e.PegawaiGender,
		PegawaiPhonenumber: e.PegawaiPhonenumber,
		PegawaiCreateAt:    e.PegawaiCreateAt,
		PegawaiUpdateAt:    e.PegawaiUpdateAt,
		PegawaiDeleteAt:    e.PegawaiDeleteAt,
	}
}

func (model *Pegawai) ToEntity() *entity.Pegawai {
	modelData := &entity.Pegawai{
		PegawaiId:          model.PegawaiId,
		PegawaiCode:        model.PegawaiCode,
		PegawaiName:        model.PegawaiName,
		PegawaiGender:      model.PegawaiGender,
		PegawaiPhonenumber: model.PegawaiPhonenumber,
		PegawaiCreateAt:    model.PegawaiCreateAt,
		PegawaiUpdateAt:    model.PegawaiUpdateAt,
		PegawaiDeleteAt:    model.PegawaiDeleteAt,
	}

	return modelData
}

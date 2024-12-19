package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type PegawaiRepository interface {
	Create(ctx context.Context, pegawai entity.Pegawai) entity.Pegawai
	Update(ctx context.Context, pegawai entity.Pegawai, pegawaiId string) entity.Pegawai
	Delete(ctx context.Context, pegawaiId string)
	FindById(ctx context.Context, pegawaiId string) (entity.Pegawai, error)
	FindAll(ctx context.Context) []entity.Pegawai
	FindSpesificData(ctx context.Context, where entity.Pegawai) []entity.Pegawai
}

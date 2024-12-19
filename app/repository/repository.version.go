package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type VersionRepository interface {
	GetVersionAdmin(ctx context.Context) entity.VersionAdmin
	GetVersionShop(ctx context.Context) entity.VersionShop
}

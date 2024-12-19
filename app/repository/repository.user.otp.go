package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type UserOTPRepository interface {
	Create(ctx context.Context, otp entity.UserOTP) entity.UserOTP
	Update(ctx context.Context, otp entity.UserOTP, otpId string) entity.UserOTP
	Delete(ctx context.Context, otpId string)
	FindById(ctx context.Context, otpId string) (entity.UserOTP, error)
	FindAll(ctx context.Context) []entity.UserOTP
	FindSpesificData(ctx context.Context, where entity.UserOTP) []entity.UserOTP
}

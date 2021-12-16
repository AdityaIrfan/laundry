package otprps

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/ports"
	"github.com/xlzd/gotp"
)

type otpDefault struct {
	redis *redis.Client
}

func NewOtpDefault(redis *redis.Client) ports.OtpRepository {
	return &otpDefault{
		redis: redis,
	}
}

func (instance *otpDefault) GenerateOTP(ctx context.Context, uuid string, expired uint) (*domain.DefaultOTP, error) {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	totp.Now()          // current otp '123456'
	totp.At(1524486261) // otp of timestamp 1524486261 '123456'

	// OTP verified for a given timestamp
	/* totp.Verify('492039', 1524486261)  // true
	totp.Verify('492039', 1520000000)  // false */
	return nil, nil
}

func (instance *otpDefault) GetOTPByUUID(ctx context.Context, uuid string) (*domain.DefaultOTP, error) {
	return nil, nil
}

func (instacne *otpDefault) DeleteOTP(ctx context.Context, uuid string) error {
	return nil
}

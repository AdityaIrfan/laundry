package ports

import (
	"context"

	"github.com/laundry/internal/core/domain"
)

type (
	AuthRepository interface {
		GenerateSession(ctx context.Context, uuid string, expired uint) (*domain.GenerateSessionRespose, error)
		ValidateSession(ctx context.Context, sessionToken string) (*domain.ValidateSession, error)
		GenerateJWT(ctx context.Context, user domain.User) (*domain.GenerateJWTResponse, error)
		GeneratePassword(ctx context.Context, password string) (string, error)
		ValidateRefreshToken(ctx context.Context, refreshToken string) (*domain.RefreshTokenDataClaim, error)
		GenerateAccessToken(ctx context.Context, user domain.User) (*domain.GenerateJWTResponse, error)
		DeleteRefreshToken(ctx context.Context, jti string) error
	}

	UserRepository interface {
		GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
		GetUserByPhone(ctx context.Context, email string) (*domain.User, error)
		Create(ctx context.Context, userModel domain.User) (*domain.User, error)
		GetUserByUUID(ctx context.Context, uuid string) (*domain.User, error)
		Update(ctx context.Context, userModel domain.User) (*domain.User, error)
	}

	OtpRepository interface {
		GenerateOTP(ctx context.Context, uuid string, expired uint) (*domain.DefaultOTP, error)
		GetOTPByUUID(ctx context.Context, uuid string) (*domain.DefaultOTP, error)
		DeleteOTP(ctx context.Context, uuid string) error
	}
)

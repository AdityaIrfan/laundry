package ports

import (
	"context"

	"github.com/laundry/internal/core/domain"
)

type (
	RegisterService interface {
		BeforeRegisterWithEmail(ctx context.Context, request domain.BeforeRegisterWithEmail) (*domain.BeforeRegisterResponse, error)
		BeforeRegisterWithPhone(ctx context.Context, request domain.BeforeRegisterWithPhone) (*domain.BeforeRegisterResponse, error)
		ConfirmationRegister(ctx context.Context, request domain.ConfirmationRegisterRequest) (*domain.GenerateSessionRespose, error)
		DoRegister(ctx context.Context, request domain.DoRegisterRequest) (*domain.DoRegisterResponse, error)
		ResendCode(ctx context.Context, request domain.ResendCodeRequest) (*domain.DefaultOTP, error)
	}

	LoginService interface {
		DoLogin(ctx context.Context, request domain.DoLoginRequest) (*domain.DoLoginResponse, error)
		DoRefreshToken(ctx context.Context, request domain.DoRefreshTokenRequest) (*domain.DoRefreshTokenResponse, error)
		DoLogout(ctx context.Context, request domain.DoLogoutRequest) (*string, error)
	}
)

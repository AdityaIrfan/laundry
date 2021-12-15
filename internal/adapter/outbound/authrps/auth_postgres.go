package authrps

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/ports"
	"github.com/laundry/internal/core/transformer"
	viperPkg "github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	postgres *gorm.DB
}

func NewAuthPostgres(pg *gorm.DB) ports.AuthRepository {
	return &AuthPostgres{
		postgres: pg,
	}
}

func (instance *AuthPostgres) GenerateJWT(ctx context.Context, user domain.User) (*transformer.GenerateJWTTransformer, error) {
	var (
		// access
		ISSUER_ACCESS = "laundy-JWT-Access"
		// refresh
		ISSUER_REFRESH     = "laundry-JWT-refresh"
		EXPIRESAT          = time.Now().Add(time.Duration(1) * time.Hour)
		JWT_SIGNING_METHOD = jwt.SigningMethodHS256
		JWT_SIGNATURE_KEY  = viperPkg.GetString("jwt.signature")
	)

	// create access claims
	accessClaims := transformer.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ISSUER_ACCESS,
			ExpiresAt: EXPIRESAT.Unix(),
			Id:        uuid.New().String(),
		},
		UUID:        user.UUID,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
	accessToken := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		accessClaims,
	)
	signedAccessToken, err := accessToken.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return nil, err
	}

	// create refresh claims
	refreshClaims := transformer.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: ISSUER_REFRESH,
			Id:     uuid.New().String(),
		},
		UUID: user.UUID,
	}
	refreshToken := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		refreshClaims,
	)
	signedRefreshToken, err := refreshToken.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return nil, err
	}

	return transformer.ToGenereteJWTTransformer(signedAccessToken, signedRefreshToken, accessClaims.ExpiresAt), nil
}

func (instance *AuthPostgres) GeneratePassword(ctx context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

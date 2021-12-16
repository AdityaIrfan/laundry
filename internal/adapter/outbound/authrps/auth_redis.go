package authrps

import (
	"context"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/ports"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var (
	// jwt
	ISSUER_ACCESS      = "laundy-JWT-Access"
	ISSUER_REFRESH     = "laundry-JWT-refresh"
	EXPIRESAT          = time.Now().Add(time.Duration(1) * time.Hour)
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY  = viper.GetString("jwt.signature")

	// error return
	errNotValidToken          = "invalid identity"
	SuccessDeleteRefreshToken = "successfully deleted token"
)

type authRepo struct {
	redis *redis.Client
}

func NewAuthPostgres(redis *redis.Client) ports.AuthRepository {
	return &authRepo{
		redis: redis,
	}
}

func (instance *authRepo) GenerateJWT(ctx context.Context, user domain.User) (*domain.GenerateJWTResponse, error) {
	// create access claims
	accessClaims := domain.AccessTokenClaims{
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
	refreshClaims := domain.RefreshTokenClaims{
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

	// insert to redis
	err = instance.redis.Set(ctx, refreshClaims.StandardClaims.Id, refreshClaims.UUID, 0).Err()
	if err != nil {
		return nil, err
	}

	return domain.ToGenereteJWTTransformer(signedAccessToken, signedRefreshToken, accessClaims.ExpiresAt), nil
}

func (instance *authRepo) GeneratePassword(ctx context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (instance *authRepo) GenerateSession(ctx context.Context, uuid string, expired uint) (*domain.GenerateSessionRespose, error) {
	session := base64.StdEncoding.EncodeToString([]byte(uuid))
	inSecond, err := time.ParseDuration(strconv.FormatUint(uint64(expired), 10) + "s")
	if err != nil {
		return nil, err
	}
	err = instance.redis.SetEX(ctx, session, uuid, inSecond).Err()
	if err != nil {
		return nil, err
	}
	return &domain.GenerateSessionRespose{
		SessionToken: session,
	}, nil
}

func (instance *authRepo) ValidateSession(ctx context.Context, sessionToken string) (*domain.ValidateSession, error) {
	val, err := instance.redis.Get(ctx, sessionToken).Result()
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, nil
	}
	return &domain.ValidateSession{
		UUID: val,
	}, nil
}

func (instance *authRepo) ValidateRefreshToken(ctx context.Context, refreshToken string) (*domain.RefreshTokenDataClaim, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(errNotValidToken)
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New(errNotValidToken)
		}
		return []byte(viper.GetString("jwt.refresh_signature")), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	var (
		jti *string
	)
	getJti := claims["jti"].(string)
	jti = &getJti
	if jti == nil {
		return nil, errors.New(errNotValidToken)
	}
	val, err := instance.redis.Get(ctx, *jti).Result()
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, errors.New(errNotValidToken)
	}
	return &domain.RefreshTokenDataClaim{
		UUID: val,
		JTI:  *jti,
	}, nil
}

func (instance *authRepo) GenerateAccessToken(ctx context.Context, user domain.User) (*domain.GenerateJWTResponse, error) {
	accessClaims := domain.AccessTokenClaims{
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
	signedAccessToken, err := accessToken.SignedString([]byte(viper.GetString("jwt.signature")))
	if err != nil {
		return nil, err
	}
	return &domain.GenerateJWTResponse{
		AccessToken: signedAccessToken,
		ExpiresIn:   accessClaims.ExpiresAt,
	}, nil
}

func (instance *authRepo) DeleteRefreshToken(ctx context.Context, jti string) error {
	err := instance.redis.Del(ctx, jti).Err()
	if err != nil {
		return err
	}
	return nil
}

package domain

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaims struct {
	jwt.StandardClaims
	UUID        string  `json:"uuid"`
	Username    string  `json:"username"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UUID string `json:"email"`
}

type RefreshTokenDataClaim struct {
	UUID string
	JTI  string
}

type GenerateJWTResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func ToGenereteJWTTransformer(signedAcessToken string, signedRefreshToken string, accessClaimExpiresAt int64) *GenerateJWTResponse {
	return &GenerateJWTResponse{
		AccessToken:  signedAcessToken,
		RefreshToken: signedRefreshToken,
		ExpiresIn:    accessClaimExpiresAt,
	}
}

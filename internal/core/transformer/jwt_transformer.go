package transformer

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

type GenerateJWTTransformer struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func ToGenereteJWTTransformer(signedAcessToken string, signedRefreshToken string, accessClaimExpiresAt int64) *GenerateJWTTransformer {
	return &GenerateJWTTransformer{
		AccessToken:  signedAcessToken,
		RefreshToken: signedRefreshToken,
		ExpiresIn:    accessClaimExpiresAt,
	}
}

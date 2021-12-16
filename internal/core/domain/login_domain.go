package domain

type DoLoginRequest struct {
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Username *string `json:"username"`
	Password string  `json:"password"`
}

type DoLoginResponse struct {
	User         *User  `json:"user,omitempty"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type DoRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type DoRefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type RefreshTokenDataClaims struct {
	UUID string
	JTI  string
}

type DoLogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

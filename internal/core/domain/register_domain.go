package domain

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BeforeRegisterWithEmail struct {
	Email string `json:"email"`
}

type BeforeRegisterWithPhone struct {
	Phone string `json:"phone"`
}

type BeforeRegisterResponse struct {
	ID    uint64 `json:"id"`
	UUID  string `json:"uuid"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	OTP   string `json:"otp,omitempty"`
}

type ConfirmationRegisterRequest struct {
	UUID string `json:"uuid"`
	OTP  string `json:"otp"`
}

type DoRegisterRequest struct {
	SessionToken    string `json:"session_token"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (dr *DoRegisterRequest) ValidateDoRegisterRequest() error {
	return validation.ValidateStruct(&dr,
		validation.Field(&dr.Name, validation.Required),
		validation.Field(&dr.Password, validation.Required, validation.Length(6, 20)),
		validation.Field(&dr.ConfirmPassword, validation.Required),
	)
}

func (dr *DoRegisterRequest) IsPasswordMatch() bool {
	return dr.Password == dr.ConfirmPassword
}

type DoRegisterResponse struct {
	User         *UserTransformer `json:"user,omitempty"`
	RefreshToken string           `json:"refresh_token"`
	AccessToken  string           `json:"access_token"`
	ExpiresIn    int64            `json:"expires_in"`
}

type ResendCodeRequest struct {
	UUID string `json:"uuid"`
}

type DefaultOTP struct {
	UUID string `json:"uuid"`
	OTP  string `json:"otp"`
}

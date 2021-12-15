package domain

type BeforeRegisterWithEmail struct {
	Email string `json:"email"`
}

type BeforeRegisterWithPhone struct {
	Phone string `json:"phone"`
}

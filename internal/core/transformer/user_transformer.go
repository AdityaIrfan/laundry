package transformer

type UserTransformer struct {
	UUID        string  `json:"uuid"`
	Username    string  `json:"username"`
	Description string  `json:"description"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Password    string  `json:"password"`
}

func (c Connections) ToUserTransformer() *UserTransformer {
	return &UserTransformer{
		UUID:        c.User.UUID,
		Username:    c.User.Username,
		Description: c.User.Description,
		PhoneNumber: c.User.PhoneNumber,
		Email:       c.User.Email,
		Password:    c.User.Password,
	}
}

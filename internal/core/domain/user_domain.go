package domain

type User struct {
	UUID        string
	Username    string
	Description string
	PhoneNumber *string
	Email       *string
	Password    string
}

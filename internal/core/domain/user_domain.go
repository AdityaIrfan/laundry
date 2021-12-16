package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserStatus string

const (
	UserNew       UserStatus = "new"
	UserConfirmed UserStatus = "confirmed"
	UserVerified  UserStatus = "verified"
)

type User struct {
	ID          uint64
	UUID        string
	Username    string
	Description string
	PhoneNumber *string
	Email       *string
	Status      UserStatus
	Password    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserTransformer struct {
	UUID        string     `json:"uuid"`
	Username    string     `json:"username"`
	Description string     `json:"description"`
	PhoneNumber *string    `json:"phone_number"`
	Email       *string    `json:"email"`
	Status      UserStatus `json:"status"`
}

func (u *User) ToUserTransformer() *UserTransformer {
	return &UserTransformer{
		UUID:        u.UUID,
		Username:    u.Username,
		Description: u.Description,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
}

func (u *User) IsEmpty() bool {
	return u == nil
}

func (u *User) IsStatusVerified() bool {
	return u.Status == UserVerified
}

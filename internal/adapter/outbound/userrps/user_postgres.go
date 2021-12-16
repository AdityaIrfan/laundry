package userrps

import (
	"context"

	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/ports"
	"gorm.io/gorm"
)

type userRepo struct {
	postgres *gorm.DB
}

func NewUserPostgres(postgres *gorm.DB) ports.UserRepository {
	return &userRepo{
		postgres: postgres,
	}
}

func (instance *userRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User

	err := instance.postgres.Where("email = ?", email).Order("updated_at desc").Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}

func (instance *userRepo) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user *domain.User

	err := instance.postgres.Where("phone = ?", phone).Order("updated_at desc").Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}

	return user, nil
}

func (instance *userRepo) Create(ctx context.Context, userModel domain.User) (*domain.User, error) {
	err := instance.postgres.Create(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (instance *userRepo) Update(ctx context.Context, userModel domain.User) (*domain.User, error) {
	err := instance.postgres.Save(&userModel).Error
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (instance *userRepo) GetUserByUUID(ctx context.Context, uuid string) (*domain.User, error) {
	var userModel *domain.User

	err := instance.postgres.Where("uuid = ?", uuid).Find(&userModel).Error
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

package registersvc

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/ports"
	responseErr "github.com/laundry/internal/error"
)

var (
	ErrDefault     = "Jaringan sibuk, coba beberapa saat lagi"
	ExistEmail     = "Email telah digunakan"
	ErrUserCreated = "Gagal mendaftarkan user"
)

type registerService struct {
	authRepo ports.AuthRepository
	userRepo ports.UserRepository
}

func NewRegisterService(authRepo ports.AuthRepository, userRepo ports.UserRepository) ports.RegisterService {
	return &registerService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (instance *registerService) BeforeRegisterWithEmail(ctx context.Context, request domain.BeforeRegisterWithEmail) (*domain.BeforeRegisterResponse, error) {
	// check user is empty
	user, err := instance.userRepo.GetUserByEmail(context.Background(), request.Email)
	if err != nil {
		return nil, responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(ErrDefault))
	}

	//email already used or haven't verified
	if !user.IsEmpty() && user.IsStatusVerified() {
		return nil, responseErr.New(fiber.StatusUnprocessableEntity, responseErr.WithMessage(ExistEmail))
	}

	if user.IsEmpty() {
		// create user
		user, err := instance.userRepo.Create(context.Background(), domain.User{
			UUID:   uuid.New().String(),
			Email:  &request.Email,
			Status: domain.UserNew,
		})
		if err != nil {
			return nil, responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(ErrUserCreated))
		}

		// Generate OTP
		otp, err := instance.authRepo
	}
}

func (instance *registerService) BeforeRegisterWithPhone(ctx context.Context, request domain.BeforeRegisterWithPhone) (*domain.BeforeRegisterResponse, error) {
	return nil, nil
}

func (instance *registerService) ConfirmationRegister(ctx context.Context, request domain.ConfirmationRegisterRequest) (*domain.GenerateSessionRespose, error) {
	return nil, nil
}

func (instance *registerService) DoRegister(ctx context.Context, request domain.DoRegisterRequest) (*domain.DoRegisterResponse, error) {
	return nil, nil
}

func (instance *registerService) ResendCode(ctx context.Context, request domain.ResendCodeRequest) (*domain.DefaultOTP, error) {
	return nil, nil
}

package ports

import (
	"context"

	"github.com/laundry/internal/core/domain"
	"github.com/laundry/internal/core/transformer"
)

type (
	AuthRepository interface {
		GenerateJWT(ctx context.Context, user domain.User) (*transformer.GenerateJWTTransformer, error)
		GeneratePassword(ctx context.Context, password string) (string, error)
	}
)

package ports

import "context"

type (
	RegisterService interface {
		BeforeRegisterWithEmail(ctx context.Context)
	}
)

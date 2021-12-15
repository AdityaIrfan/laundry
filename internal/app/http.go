package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handlers struct {
	Postgres *gorm.DB
	R        *fiber.App
	// Logger   *zap.Logger
}

func (h *Handlers) SetupRouter() {
	// initialize repository

	// initialize bussiness

	// initialize handlers
}

package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handlers struct {
	Postgres *gorm.DB
	R        *fiber.App
	// Logger   *zap.Logger
	Redis *redis.Client
}

func (h *Handlers) SetupRouter() {
	// initialize repository

	// initialize bussiness

	// initialize handlers
}

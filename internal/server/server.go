package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	baseApp "github.com/laundry/internal/app"
	"github.com/laundry/pkg/postgres"
	"github.com/laundry/pkg/viper"
	viperPkg "github.com/spf13/viper"
)

func Run() {
	//load config
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")
	config := &viper.EnvConfig{
		FileName: "config",
		FileType: "yaml",
		Path:     basePath,
	}
	if err := config.ReadConfig(); err != nil {
		log.Fatal("failed to read config.yaml: ", err)
	}

	// load connection postgres
	pg, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := pg.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// initialize logger
	// zap, err := logger.initialize()

	//load fiber
	app := fiber.New(fiber.Config{
		IdleTimeout: 5,
	})
	app.Use(
		recover.New(),
		compress.New(),
		etag.New(),
		cors.New(),
		fiberlog.New(),
	)

	rh := &baseApp.Handlers{
		Postgres: pg,
		R:        app,
		// Logger:   zap,
	}
	rh.SetupRouter()

	// listen from different goroutine
	go func() {
		if err := app.Listen(":" + viperPkg.GetString("server.port")); err != nil {
			log.Panicf("failed listen into port %v", err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	var _ = <-c // This blocks the main thread until an interrupt is received
	log.Println("gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	sqlDB.Close()
	fmt.Println("services was successful shutdown.")
}

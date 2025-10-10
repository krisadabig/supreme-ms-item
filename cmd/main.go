package main

import (
	"fmt"
	nethttp "net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	gormio "gorm.io/gorm"

	"github.com/krisadabig/supreme-ms-item/config"
	"github.com/krisadabig/supreme-ms-item/internal/adapters/primary/http"
	"github.com/krisadabig/supreme-ms-item/internal/adapters/secondary/logger"
	"github.com/krisadabig/supreme-ms-item/internal/adapters/secondary/storage/gorm"
	"github.com/krisadabig/supreme-ms-item/internal/core/services"
)

func main() {
	// Initialize logger first
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	appEnv := os.Getenv("APP_ENV")

	logLevel := zerolog.DebugLevel
	log := logger.New(
		logger.WithLevel(logLevel),
		logger.WithOutput(os.Stdout),
		// logger.WithPrettyConsole(),
	)

	// Log application startup
	log.Info("starting application initialization")
	log.With("environment", appEnv).
		With("log_level", logLevel.String()).
		Info("loading configuration")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	// Initialize Echo
	e := echo.New()
	e.Use(http.CORSMiddleware(cfg.Server.AllowedOrigins))
	e.Use(http.Logger(log))
	apiV1 := e.Group("/api/v1")

	// Initialize database connection
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_mode=session&prepare_statement_mode=simple",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		// cfg.Database.SSLMode,
	)

	db, err := gormio.Open(postgres.Open(dsn), &gormio.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Initialize application components
	itemRepo := gorm.NewGormItemRepository(db)
	itemService := services.NewItemService(itemRepo, log)
	itemHandler := http.NewItemHandler(itemService, log)

	// Setup routes
	itemHandler.RegisterRoutes(apiV1)

	// health check
	e.GET("/ping", func(c echo.Context) error {
		return c.String(nethttp.StatusOK, "pong")
	})

	// Log successful initialization
	log.With("version", "1.0.0").Info("application initialized successfully")

	// Start server
	addr := cfg.Server.Port
	log.With("address", addr).Info("starting http server")

	if err := e.Start(addr); err != nil {
		log.Fatal("Failed to start server", err)
	}

}

package bootstrap

import (
	"os"

	"item/internal/config"
	"item/internal/logger"
	"item/internal/middleware"
	"item/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/supabase-community/supabase-go"
)

func Init() *echo.Echo {
	// Initialize logger first
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	appEnv := os.Getenv("APP_ENV")

	logLevel := zerolog.InfoLevel
	pretty := true
	if appEnv == "production" {
		pretty = false
		logLevel = zerolog.InfoLevel
	} else {
		logLevel = zerolog.DebugLevel
	}
	logger.Init(logLevel, pretty)

	// Log application startup
	logger.Info("starting application initialization")
	logger.GetGlobalLogger().
		With("environment", appEnv).
		With("log_level", logLevel.String()).
		Info("bootstrap configuration")

	// Load config
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", err)
	}

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.CORSMiddleware(cfg.Server.AllowedOrigins))

	// Initialize Supabase client
	sb, err := supabase.NewClient(cfg.Supabase.URL, cfg.Supabase.AnonKey, &supabase.ClientOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + cfg.Supabase.AnonKey,
			"apikey":        cfg.Supabase.AnonKey,
		},
	})
	if err != nil {
		logger.Fatal("Failed to create Supabase client", err)
	}

	// Initialize application components
	itemRepo := InitRepositories(sb)
	itemService := InitServices(itemRepo, logger.GetGlobalLogger())
	itemHandler := InitHandlers(e, itemService, logger.GetGlobalLogger())

	// Setup routes
	router.SetupRoutes(e, itemHandler)

	// Log successful initialization
	logger.GetGlobalLogger().With("version", "1.0.0").Info("application initialized successfully")

	// Start server
	addr := cfg.Server.Port
	logger.GetGlobalLogger().With("address", addr).Info("starting http server")

	if err := e.Start(addr); err != nil {
		logger.Fatal("Failed to start server", err)
	}

	return e
}

package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port           string   `mapstructure:"port"`
		AllowedOrigins []string `mapstructure:"allowed_origins"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

func Load() (*Config, error) {
	// Get environment (default to "dev" if not set)
	env := strings.ToLower(getEnv("APP_ENV", "dev"))

	// Set config file name based on environment
	configName := "config"
	if env != "dev" && env != "development" {
		configName = "config." + env
	}

	viper.SetConfigName(configName) // Name of config file (without extension)
	viper.SetConfigType("yaml")     // Or "toml", "json"
	viper.AddConfigPath(".")        // Look in current dir; add more paths if needed

	// Load config file (optional fallback)
	if err := viper.ReadInConfig(); err != nil {
		// If environment-specific config doesn't exist, try default config
		if configName != "config" {
			viper.SetConfigName("config")
			if err := viper.ReadInConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Override with env vars (prefixed with APP_ for safety)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Convert . to _ for env
	viper.BindEnv("server.port", "PORT")                   // Allow plain PORT too
	viper.BindEnv("server.allowed_origins", "ALLOWED_ORIGINS")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.dbname", "DB_NAME")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("Failed to unmarshal config:", err)
		return nil, err
	}

	// Handle comma-separated env overrides for allowed origins
	if raw := viper.GetString("server.allowed_origins"); raw != "" {
		parts := strings.Split(raw, ",")
		cfg.Server.AllowedOrigins = cfg.Server.AllowedOrigins[:0]
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				cfg.Server.AllowedOrigins = append(cfg.Server.AllowedOrigins, trimmed)
			}
		}
	}

	return &cfg, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

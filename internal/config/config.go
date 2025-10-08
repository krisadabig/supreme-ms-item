// internal/config/config.go
package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Supabase struct {
		URL     string `mapstructure:"url"`
		AnonKey string `mapstructure:"anon_key"`
	} `mapstructure:"supabase"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Or "toml", "json"
	viper.AddConfigPath(".")      // Look in current dir; add more paths if needed

	// Load config file (optional fallback)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Load .env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// bind value from .env to viper
	config := viper.AllSettings()
	if config["port"] != nil {
		viper.Set("server.port", config["port"])
	}
	if config["supabase_url"] != nil {
		viper.Set("supabase.url", config["supabase_url"])
	}
	if config["supabase_anon_key"] != nil {
		viper.Set("supabase.anon_key", config["supabase_anon_key"])
	}

	// Override with env vars (prefixed with APP_ for safety)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Convert . to _ for env
	viper.BindEnv("server.port", "PORT")                   // Allow plain PORT too
	viper.BindEnv("supabase.url", "SUPABASE_URL")
	viper.BindEnv("supabase.anon_key", "SUPABASE_ANON_KEY")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("Failed to unmarshal config:", err)
		return nil, err
	}

	return &cfg, nil
}

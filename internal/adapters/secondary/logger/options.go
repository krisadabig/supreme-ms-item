package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// config holds the configuration for the logger.
type config struct {
	level      zerolog.Level
	output     io.Writer
	timeFormat string
}

// Option defines a function that configures the logger.
type Option func(*config)

// WithLevel sets the logging level.
func WithLevel(level zerolog.Level) Option {
	return func(c *config) {
		c.level = level
	}
}

// WithPrettyConsole sets the output to a human-readable console format.
func WithPrettyConsole() Option {
	return func(c *config) {
		c.output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}
}

// WithOutput sets a custom output writer for the logs.
func WithOutput(writer io.Writer) Option {
	return func(c *config) {
		c.output = writer
	}
}

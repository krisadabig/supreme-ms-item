package logger

import (
	"context"
	"os"
	"time"

	"github.com/krisadabig/supreme-ms-item/internal/constants"
	"github.com/krisadabig/supreme-ms-item/internal/core/ports"

	"github.com/rs/zerolog"
)

// zerologAdapter implements the ports.Logger interface using zerolog.
type zerologAdapter struct {
	logger zerolog.Logger
}

// New creates a new logger instance that implements ports.Logger.
// It is configured through options for better flexibility.
func New(opts ...Option) ports.Logger {
	// Sensible defaults
	cfg := config{
		level:      zerolog.InfoLevel,
		output:     os.Stdout,
		timeFormat: time.RFC3339,
	}

	// Apply custom options
	for _, opt := range opts {
		opt(&cfg)
	}

	logger := zerolog.New(cfg.output).
		Level(cfg.level).
		With().
		Timestamp().
		Logger()

	return &zerologAdapter{logger: logger}
}

func (l *zerologAdapter) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *zerologAdapter) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *zerologAdapter) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *zerologAdapter) Error(msg string, err error) {
	event := l.logger.Error()
	if err != nil {
		event = event.Err(err)
	}
	event.Msg(msg)
}

func (l *zerologAdapter) Fatal(msg string, err error) {
	l.logger.Fatal().Err(err).Msg(msg)
}

func (l *zerologAdapter) With(key string, value any) ports.Logger {
	newLogger := l.logger.With().Interface(key, value).Logger()
	return &zerologAdapter{logger: newLogger}
}

func (l *zerologAdapter) WithFields(fields ...ports.Field) ports.Logger {
	contextBuilder := l.logger.With()
	for _, f := range fields {
		contextBuilder = contextBuilder.Interface(f.Key, f.Value)
	}
	return &zerologAdapter{logger: contextBuilder.Logger()}
}

// WithContext extracts values from a context and returns a child logger.
// This is useful for adding request-specific data like correlation IDs.
func (l *zerologAdapter) WithContext(ctx context.Context) ports.Logger {
	if ctx == nil {
		return l
	}

	// Example: Extract a correlation ID from context
	// You would set this in a middleware.
	if correlationID, ok := ctx.Value(constants.ContextCorrelationID).(string); ok && correlationID != "" {
		newLogger := l.logger.With().Str("request_id", correlationID).Logger()
		return &zerologAdapter{logger: newLogger}
	}

	return l
}

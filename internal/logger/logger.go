package logger

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"item/domain"
	"item/internal/constants"

	"github.com/rs/zerolog"
)

// zerologLogger implements the domain.Logger interface using zerolog
type zerologLogger struct {
	logger zerolog.Logger
}

// New creates a new logger instance that implements domain.Logger
func New(level zerolog.Level, pretty bool) domain.Logger {
	var output io.Writer = os.Stdout

	if pretty {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FieldsOrder: []string{
				"request_id",
				"method",
				"path",
				"remote_ip",
				"status",
				"latency_ms",
				"response_size",
				"response_body",
			},
		}
	}

	logger := zerolog.New(output).
		Level(level).
		With().
		Timestamp().
		Logger()

	return &zerologLogger{logger: logger}
}

func (l *zerologLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *zerologLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *zerologLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *zerologLogger) Error(msg string, err error) {
	event := l.logger.Error()
	if err != nil {
		event = event.Err(err)
	}
	event.Msg(msg)
}

func (l *zerologLogger) With(key string, value any) domain.Logger {
	ctx := l.logger.With().Interface(key, value)
	return &zerologLogger{logger: ctx.Logger()}
}

func (l *zerologLogger) WithFields(fields ...domain.Field) domain.Logger {
	contextBuilder := l.logger.With()
	for _, f := range fields {
		contextBuilder = contextBuilder.Interface(f.Key, f.Value)
	}
	return &zerologLogger{logger: contextBuilder.Logger()}
}

func (l *zerologLogger) WithContext(ctx context.Context) domain.Logger {
	if ctx == nil {
		return l
	}

	logger := l.logger
	if correlationID, ok := ctx.Value(constants.ContextCorrelationID).(string); ok && correlationID != "" {
		logger = logger.With().Str("request_id", correlationID).Logger()
	}

	return &zerologLogger{logger: logger}
}

// Global logger instance
var global domain.Logger

// Init initializes the global logger
func Init(level zerolog.Level, pretty bool) {
	global = New(level, pretty)
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() domain.Logger {
	return global
}

// SetGlobalLogger sets the global logger instance
func SetGlobalLogger(logger domain.Logger) {
	global = logger
}

// Global helper functions

// Debug logs a debug message using the global logger
func Debug(msg string) {
	if global == nil {
		return
	}
	global.Debug(msg)
}

// Info logs an info message using the global logger
func Info(msg string) {
	if global == nil {
		return
	}
	global.Info(msg)
}

// Warn logs a warning message using the global logger
func Warn(msg string) {
	if global == nil {
		return
	}
	global.Warn(msg)
}

// Error logs an error message using the global logger
func Error(msg string, err error) {
	if global == nil {
		return
	}
	global.Error(msg, err)
}

// Fatal logs a fatal message and exits using the global logger
func Fatal(msg string, err error) {
	if global == nil {
		os.Exit(1)
	}
	if err == nil {
		err = errors.New("fatal error")
	}
	global.Error(msg, err)
	os.Exit(1)
}

package domain

import "context"

// Logger defines the interface for logging operations
type Logger interface {
	// Debug logs a debug message with the currently attached fields
	Debug(msg string)

	// Info logs an info message with the currently attached fields
	Info(msg string)

	// Warn logs a warning message with the currently attached fields
	Warn(msg string)

	// Error logs an error message with the currently attached fields
	Error(msg string, err error)

	// With attaches a single structured field and returns a child logger
	With(key string, value any) Logger

	// WithFields attaches multiple structured fields and returns a child logger
	WithFields(fields ...Field) Logger

	// WithContext returns a new logger with context values merged in
	WithContext(ctx context.Context) Logger
}

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value any
}

// Field creates a new Field
func NewField(key string, value any) Field {
	return Field{Key: key, Value: value}
}

// String creates a string field
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int creates an int field
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Int64 creates an int64 field
func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

// Error creates an error field
func Error(err error) Field {
	return Field{Key: "error", Value: err}
}

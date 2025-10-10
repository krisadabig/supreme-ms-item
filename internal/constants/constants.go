package constants

// Common HTTP headers
const (
	HeaderCorrelationID = "X-Correlation-ID"
	HeaderUserID        = "X-User-ID"
)

// Context keys
type contextKey string

const (
	// ContextCorrelationID is the key used to store the correlation ID in the context
	ContextCorrelationID contextKey = "correlation_id"
	// ContextKeyLogger is the key used to store the logger in the context
	ContextKeyLogger contextKey = "logger"
)

// Error messages
const (
	ErrInvalidRequest   = "invalid request"
	ErrInternalServer   = "internal server error"
	ErrResourceNotFound = "resource not found"
)

// API routes
const (
	APIV1 = "/api/v1"
)

// Database related constants
const (
	DBTimeout = 10 // seconds
)

package logger

import (
	"context"

	"item/domain"
	"item/internal/constants"
)

// ContextWithRequestID stores a request ID in the context for downstream logging.
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	if requestID == "" {
		return ctx
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, constants.ContextCorrelationID, requestID)
}

// RequestIDFromContext retrieves the request ID from the context if available.
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if requestID, ok := ctx.Value(constants.ContextCorrelationID).(string); ok {
		return requestID
	}
	return ""
}

// ContextWithLogger stores the provided logger in the context.
func ContextWithLogger(ctx context.Context, log domain.Logger) context.Context {
	if log == nil {
		return ctx
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, constants.ContextKeyLogger, log)
}

// FromContext tries to fetch a logger from the context, falling back to the global logger.
func FromContext(ctx context.Context) domain.Logger {
	if ctx == nil {
		return global
	}
	if log, ok := ctx.Value(constants.ContextKeyLogger).(domain.Logger); ok && log != nil {
		return log
	}
	if global == nil {
		return nil
	}
	return global.WithContext(ctx)
}

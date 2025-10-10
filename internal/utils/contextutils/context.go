package contextutils

import (
	"context"

	"github.com/krisadabig/supreme-ms-item/internal/constants"
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

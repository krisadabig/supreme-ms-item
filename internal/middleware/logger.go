package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"item/internal/constants"
	"item/internal/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// responseBodyWriter wraps http.ResponseWriter to capture the response body.
type responseBodyWriter struct {
	io.Writer
	http.ResponseWriter
	body *bytes.Buffer
}

func (rbw *responseBodyWriter) Write(b []byte) (int, error) {
	rbw.body.Write(b)
	return rbw.Writer.Write(b)
}

// Logger returns a middleware that logs HTTP requests and responses.
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			start := time.Now()

			// Get or generate request ID
			requestID := req.Header.Get(constants.HeaderCorrelationID)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			ctx := logger.ContextWithRequestID(req.Context(), requestID)
			baseLogger := logger.FromContext(ctx)
			if baseLogger == nil {
				baseLogger = logger.GetGlobalLogger()
			}
			requestLogger := baseLogger.
				With("request_id", requestID).
				With("method", req.Method).
				With("path", req.URL.Path).
				With("remote_ip", c.RealIP())

			ctx = logger.ContextWithLogger(ctx, requestLogger)
			req = req.WithContext(ctx)
			c.SetRequest(req)

			// Capture response body for logging
			res := c.Response()
			rbw := &responseBodyWriter{
				Writer:         res.Writer,
				body:           &bytes.Buffer{},
				ResponseWriter: res.Writer,
			}
			res.Writer = rbw

			// Ensure request ID is echoed back to the client
			res.Header().Set(constants.HeaderCorrelationID, requestID)

			err := next(c)

			latency := time.Since(start)
			status := res.Status
			if status == 0 {
				status = http.StatusOK
			}

			logEntry := requestLogger.
				With("status", status).
				With("latency_ms", latency.Milliseconds()).
				With("response_size", res.Size).
				With("response_body", rbw.body.String())

			if query := req.URL.RawQuery; query != "" {
				logEntry = logEntry.With("query", query)
			}

			if err != nil {
				if echoErr, ok := err.(*echo.HTTPError); ok {
					logEntry = logEntry.
						With("error_code", echoErr.Code).
						With("error", echoErr.Error())
				} else {
					logEntry = logEntry.With("error", err.Error())
				}
			}

			logEntry.Info("")

			return err
		}
	}
}

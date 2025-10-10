package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware returns a configured CORS middleware instance.
// When allowedOrigins is empty, it falls back to a wildcard.
func CORSMiddleware(allowedOrigins []string) echo.MiddlewareFunc {
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"*"}
	}

	config := echoMiddleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderAuthorization,
			"X-Requested-With",
			"X-Correlation-Id",
			"X-User-Id",
		},
		ExposeHeaders: []string{
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
	}

	return echoMiddleware.CORSWithConfig(config)
}

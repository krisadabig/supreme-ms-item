package http

// import (
// 	"item/internal/core/adapters/primary/handlers"
// 	"item/internal/middleware"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func SetupRoutes(e *echo.Echo, itemHandler *handlers.ItemHandler) {
// 	// Add request logging middleware
// 	e.Use(middleware.Logger())
// 	apiV1 := e.Group("/api/v1")

// 	itemGroup := apiV1.Group("/items")
// 	itemGroup.POST("", itemHandler.CreateItem)
// 	itemGroup.PUT("", itemHandler.UpdateItem)
// 	itemGroup.DELETE("", itemHandler.DeleteItem)
// 	itemGroup.GET("", itemHandler.GetItems)
// 	itemGroup.GET("/:id", itemHandler.GetItem)
// 	itemGroup.GET("/user/:user_id", itemHandler.GetItemsByUserID)

// 	e.GET("/ping", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "pong")
// 	})
// }

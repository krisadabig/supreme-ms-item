package bootstrap

import (
	"item/adapters/handlers"
	"item/application"
	"item/domain"

	"github.com/labstack/echo/v4"
)

func InitHandlers(e *echo.Echo, itemService *application.ItemService, logger domain.Logger) *handlers.ItemHandler {
	itemHandler := handlers.NewItemHandler(itemService, logger)

	return itemHandler
}

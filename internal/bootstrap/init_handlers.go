package bootstrap

import (
	"item/adapters/handlers"
	"item/application"

	"github.com/labstack/echo/v4"
)

func InitHandlers(e *echo.Echo, itemService *application.ItemService) *handlers.ItemHandler {
	itemHandler := handlers.NewItemHandler(itemService)

	return itemHandler
}

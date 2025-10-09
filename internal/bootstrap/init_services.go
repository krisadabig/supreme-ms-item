package bootstrap

import (
	"item/application"
	"item/domain"
)

func InitServices(itemRepo domain.ItemRepository, logger domain.Logger) *application.ItemService {
	itemService := application.NewItemService(itemRepo, logger)

	return itemService
}

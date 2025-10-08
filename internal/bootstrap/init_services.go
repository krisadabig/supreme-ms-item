package bootstrap

import (
	"item/application"
	"item/domain"
)

func InitServices(itemRepo domain.ItemRepository) *application.ItemService {
	itemService := application.NewItemService(itemRepo)

	return itemService
}

package application

import (
	"context"
	"fmt"
	"item/domain"
)

type ItemService struct {
	repo   domain.ItemRepository
	logger domain.Logger
}

func NewItemService(repo domain.ItemRepository, logger domain.Logger) *ItemService {
	return &ItemService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ItemService) Create(ctx context.Context, item *domain.Item) error {
	log := s.logger.WithContext(ctx).With("operation", "create_item")

	if err := item.Validate(); err != nil {
		log.Error("validation failed", err)
		return fmt.Errorf("validation failed: %w", err)
	}

	title := ""
	if item.Title != nil {
		title = *item.Title
	}

	log.With("user_id", item.UserID).
		With("title", title).
		Info("creating item")

	if err := s.repo.Create(item); err != nil {
		log.With("user_id", item.UserID).
			With("title", title).
			Error("failed to create item", err)
		return fmt.Errorf("failed to create item: %w", err)
	}

	log.With("item_id", item.ID).
		With("user_id", item.UserID).
		Info("item created successfully")
	return nil
}

func (s *ItemService) Update(ctx context.Context, item *domain.Item) error {
	log := s.logger.WithContext(ctx).
		With("operation", "update_item").
		With("item_id", item.ID).
		With("user_id", item.UserID)

	if err := item.Validate(); err != nil {
		log.Error("validation failed", err)
		return fmt.Errorf("validation failed: %w", err)
	}

	if item.ID == 0 {
		log.Error("cannot update item with id 0", nil)
		return domain.ErrInvalidItem
	}

	log.Info("updating item")
	err := s.repo.Update(item)
	if err != nil {
		log.Error("failed to update item", err)
		return fmt.Errorf("failed to update item: %w", err)
	}

	log.Info("item updated successfully")
	return nil
}

func (s *ItemService) Delete(ctx context.Context, item *domain.Item) error {
	log := s.logger.WithContext(ctx).
		With("operation", "delete_item").
		With("item_id", item.ID).
		With("user_id", item.UserID)

	if item.ID == 0 {
		log.Error("cannot delete item with id 0", nil)
		return domain.ErrInvalidItem
	}

	// Check if item exists
	if _, err := s.repo.GetByID(item.ID); err != nil {
		log.Error("Item not found for deletion", err)
		return fmt.Errorf("item not found: %w", err)
	}

	log.Info("deleting item")
	err := s.repo.Delete(item)
	if err != nil {
		log.Error("failed to delete item", err)
		return fmt.Errorf("failed to delete item: %w", err)
	}

	log.Info("item deleted successfully")
	return nil
}

func (s *ItemService) GetAll(ctx context.Context) ([]domain.Item, error) {
	log := s.logger.WithContext(ctx).With("operation", "get_all_items")

	log.Debug("fetching all items")

	items, err := s.repo.GetAll()
	if err != nil {
		log.Error("failed to fetch items", err)
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}

	log.With("count", len(items)).Debug("successfully fetched items")
	return items, nil
}

func (s *ItemService) GetByID(ctx context.Context, id int64) (*domain.Item, error) {
	log := s.logger.WithContext(ctx).
		With("operation", "get_item_by_id").
		With("item_id", id)

	log.Debug("fetching item by id")
	item, err := s.repo.GetByID(id)
	if err != nil {
		log.Error("failed to fetch item by id", err)
		return nil, fmt.Errorf("failed to fetch item by ID: %w", err)
	}

	log.Debug("successfully fetched item by id")
	return item, nil
}

func (s *ItemService) GetByUserID(ctx context.Context, userID string) ([]domain.Item, error) {
	log := s.logger.WithContext(ctx).
		With("operation", "get_items_by_user_id").
		With("user_id", userID)

	log.Debug("fetching items by user id")
	items, err := s.repo.GetByUserID(userID)
	if err != nil {
		log.Error("failed to fetch items by user id", err)
		return nil, fmt.Errorf("failed to fetch items by user ID: %w", err)
	}

	log.With("count", len(items)).Debug("successfully fetched items by user id")
	return items, nil
}

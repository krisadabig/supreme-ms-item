package application

import (
	"context"
	"item/domain"
)

type ItemService struct {
	repo domain.ItemRepository
}

func NewItemService(repo domain.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) Create(ctx context.Context, item *domain.Item) error {
	if err := item.Validate(); err != nil {
		return err
	}
	return s.repo.Create(item)
}

func (s *ItemService) Update(ctx context.Context, item *domain.Item) error {
	if err := item.Validate(); err != nil {
		return err
	}
	if item.ID == 0 {
		return domain.ErrInvalidItem
	}
	if _, err := s.repo.GetByID(item.ID); err != nil {
		return err
	}
	return s.repo.Update(item)
}

func (s *ItemService) Delete(ctx context.Context, item *domain.Item) error {
	if err := item.Validate(); err != nil {
		return err
	}
	if item.ID == 0 {
		return domain.ErrInvalidItem
	}
	if _, err := s.repo.GetByID(item.ID); err != nil {
		return err
	}

	return s.repo.Delete(item)
}

func (s *ItemService) GetAll(ctx context.Context) ([]domain.Item, error) {
	return s.repo.GetAll()
}

func (s *ItemService) GetByID(ctx context.Context, id int64) (*domain.Item, error) {
	return s.repo.GetByID(id)
}

func (s *ItemService) GetByUserID(ctx context.Context, userID string) ([]domain.Item, error) {
	return s.repo.GetByUserID(userID)
}

package ports

import (
	"context"

	"github.com/krisadabig/supreme-ms-item/internal/core/domain"
)

type ItemService interface {
	Create(ctx context.Context, item *domain.Item) error
	Update(ctx context.Context, item *domain.Item) error
	Delete(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context) ([]domain.Item, error)
	GetByID(ctx context.Context, id int64) (*domain.Item, error)
	GetByUserID(ctx context.Context, userID string) ([]domain.Item, error)
}

type ItemRepository interface {
	Create(item *domain.Item) error
	Update(item *domain.Item) error
	Delete(item *domain.Item) error
	GetAll() ([]domain.Item, error)
	GetByID(id int64) (*domain.Item, error)
	GetByUserID(userID string) ([]domain.Item, error)
}

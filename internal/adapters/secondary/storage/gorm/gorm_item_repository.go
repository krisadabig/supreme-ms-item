package gorm

import (
	"github.com/krisadabig/supreme-ms-item/internal/core/domain"
	"gorm.io/gorm"
)

type GormItemRepository struct {
	db *gorm.DB
}

func NewGormItemRepository(db *gorm.DB) *GormItemRepository {
	return &GormItemRepository{
		db: db,
	}
}

func (r *GormItemRepository) Create(item *domain.Item) error {
	return r.db.Create(item).Error
}

func (r *GormItemRepository) Update(item *domain.Item) error {
	return r.db.Save(item).Error
}

func (r *GormItemRepository) Delete(item *domain.Item) error {
	return r.db.Delete(item).Error
}

func (r *GormItemRepository) GetAll() ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *GormItemRepository) GetByID(id int64) (*domain.Item, error) {
	var item domain.Item
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GormItemRepository) GetByUserID(userID string) ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

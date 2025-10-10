package domain

import (
	"errors"
	"time"
)

type Item struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       *string   `json:"title" gorm:"not null"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
	UserID      string    `json:"user_id" gorm:"not null;index"`
}

var (
	ErrItemNotFound = errors.New("item not found")
	ErrItemExists   = errors.New("item already exists")
	ErrInvalidItem  = errors.New("invalid item")
)

func (i *Item) Validate() error {
	if i.Title == nil || *i.Title == "" {
		return ErrInvalidItem
	}
	return nil
}

package domain

import (
	"errors"
	"time"
)

type Item struct {
	ID          int64     `json:"id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UserID      string    `json:"user_id"`
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

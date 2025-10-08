package domain

type ItemRepository interface {
	Create(item *Item) error
	Update(item *Item) error
	Delete(item *Item) error
	GetAll() ([]Item, error)
	GetByID(id int64) (*Item, error)
	GetByUserID(userID string) ([]Item, error)
}

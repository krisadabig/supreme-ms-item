package repositories

import (
	"item/domain"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

type SupabaseItemRepository struct {
	sb *supabase.Client
}

func NewSupabaseItemRepository(sb *supabase.Client) *SupabaseItemRepository {
	return &SupabaseItemRepository{
		sb: sb,
	}
}

func (r *SupabaseItemRepository) Create(item *domain.Item) error {
	_, _, err := r.sb.From("items").Insert(
		map[string]any{
			"title":       item.Title,
			"description": item.Description,
			"user_id":     item.UserID,
		},
		false,
		"",
		"",
		"",
	).Execute()

	return err
}

func (r *SupabaseItemRepository) Update(item *domain.Item) error {
	_, _, err := r.sb.From("items").Update(
		map[string]any{
			"title":       item.Title,
			"description": item.Description,
			"user_id":     item.UserID,
		},
		"",
		"",
	).Eq("id", strconv.FormatInt(item.ID, 10)).Execute()

	return err
}

func (r *SupabaseItemRepository) Delete(item *domain.Item) error {
	_, _, err := r.sb.From("items").Delete(
		"",
		"",
	).Eq("id", strconv.FormatInt(item.ID, 10)).Execute()
	return err
}

func (r *SupabaseItemRepository) GetAll() ([]domain.Item, error) {
	var items []domain.Item
	_, err := r.sb.From("items").Select("*", "exact", false).ExecuteTo(&items)
	return items, err
}

func (r *SupabaseItemRepository) GetByID(id int64) (*domain.Item, error) {
	var item domain.Item
	_, err := r.sb.From("items").Select("*", "exact", false).Eq("id", strconv.FormatInt(id, 10)).ExecuteTo(&item)
	return &item, err
}

func (r *SupabaseItemRepository) GetByUserID(userID string) ([]domain.Item, error) {
	var items []domain.Item
	_, err := r.sb.From("items").Select("*", "exact", false).Eq("user_id", userID).ExecuteTo(&items)
	return items, err
}

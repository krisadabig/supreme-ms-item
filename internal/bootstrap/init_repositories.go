package bootstrap

import (
	"item/adapters/repositories"
	"item/domain"

	"github.com/supabase-community/supabase-go"
)

func InitRepositories(sb *supabase.Client) domain.ItemRepository {
	itemRepo := repositories.NewSupabaseItemRepository(sb)

	return itemRepo
}

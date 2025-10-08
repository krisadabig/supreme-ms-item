package bootstrap

import (
	"item/internal/config"
	"item/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/supabase-community/supabase-go"
)

func Init() *echo.Echo {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	sb, err := supabase.NewClient(cfg.Supabase.URL, cfg.Supabase.AnonKey, &supabase.ClientOptions{
		Headers: map[string]string{
			"Authorization": "Bearer " + cfg.Supabase.AnonKey,
			"apikey":        cfg.Supabase.AnonKey,
		},
	})
	if err != nil {
		panic(err)
	}

	itemRepo := InitRepositories(sb)
	itemService := InitServices(itemRepo)
	itemHandler := InitHandlers(e, itemService)

	router.SetupRoutes(e, itemHandler)

	e.Logger.Fatal(e.Start(cfg.Server.Port))

	return e
}

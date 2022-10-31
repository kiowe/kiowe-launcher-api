package app

import (
	"github.com/kiowe/kiowe-launcher-api/internal/service"
	"github.com/kiowe/kiowe-launcher-api/internal/storage"
	"github.com/kiowe/kiowe-launcher-api/internal/storage/psql"
	"github.com/kiowe/kiowe-launcher-api/internal/transport/rest"
	"github.com/kiowe/kiowe-launcher-api/internal/transport/rest/handler"
	"log"
)

type App struct {
}

func NewApp() (App, error) {
	return App{}, nil
}

func (a *App) Run() {
	log.Println("Database connection initializing...")
	pool, err := psql.NewPostgres()
	if err != nil {
		log.Printf("[ERR]: Failed to initialize database connection: %s\n", err.Error())
	}

	log.Println("Storages initializing...")
	storages := storage.New(pool)

	log.Println("Services initializing...")
	services := service.New(service.Deps{
		GameShopListStorage: storages.GameShopListStorage,
	})

	log.Println("Handlers initializing...")
	restHandlers := handler.New(handler.Deps{
		GameShopListService: services.GameShopListService,
	})

	app := restHandlers.Init()

	srv := rest.NewServer("localhost:8081", app, pool)
	srv.StartWithGracefulShutdown()
}

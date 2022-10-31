package storage

import "github.com/jackc/pgx/v4/pgxpool"

type Storage struct {
	GameShopListStorage *GameShopListStorage
}

func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		GameShopListStorage: NewGameShopListStorage(pool),
	}
}

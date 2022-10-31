package storage

import "github.com/jackc/pgx/v4/pgxpool"

type GameShopListStorage struct {
	pool *pgxpool.Pool
}

func NewGameShopListStorage(pool *pgxpool.Pool) *GameShopListStorage {
	return &GameShopListStorage{pool: pool}
}

func (s *GameShopListStorage) GetOne(id int) error {
	return nil
}

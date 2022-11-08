package storage

import "github.com/jackc/pgx/v4/pgxpool"

type Storage struct {
	DevSignupStorage    *DevSignupStorage
	GameShopListStorage *GameShopListStorage
}

func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		DevSignupStorage:    NewDevSignupStorage(pool),
		GameShopListStorage: NewGameShopListStorage(pool),
	}
}

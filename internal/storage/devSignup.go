package storage

import "github.com/jackc/pgx/v4/pgxpool"

type DevSignupStorage struct {
	pool *pgxpool.Pool
}

func NewDevSignupStorage(pool *pgxpool.Pool) *DevSignupStorage {
	return &DevSignupStorage{pool: pool}
}

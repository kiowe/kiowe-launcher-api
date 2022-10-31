package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func NewPostgres() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"kiowe",
		"8982qw",
		"localhost",
		"5432",
		"postgres",
	)

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Printf("[ERR]: DB connection error. %v", err)
		return nil, err
	}

	return pool, nil
}

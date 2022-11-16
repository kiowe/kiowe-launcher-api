package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/pkg/utils"
	"log"
)

type DevSignupStorage struct {
	pool *pgxpool.Pool
}

func NewDevSignupStorage(pool *pgxpool.Pool) *DevSignupStorage {
	return &DevSignupStorage{pool: pool}
}

func (s *DevSignupStorage) GetByLogin(login string) (bool, error) {
	sql := `SELECT EXISTS(SELECT * FROM dev_pub_account WHERE login = $1)`

	isExist := false

	if err := s.pool.QueryRow(context.Background(), sql, login).Scan(
		&isExist,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]1: %v", err)
			return isExist, err
		}
		log.Printf("[QUERY ERROR]1: %v", err)
		return isExist, err
	}

	return isExist, nil
}

func (s *DevSignupStorage) Create(dto *core.DevPubAccountDTO) (int, error) {
	sql := `INSERT INTO dev_pub_account(login, password, email, name, description) VALUES($1, $2, $3, $4, $5) RETURNING id`

	var id int

	if err := s.pool.QueryRow(context.Background(), sql,
		&dto.Login,
		&dto.Password,
		&dto.Email,
		&dto.Name,
		&dto.Description).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return 0, err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return 0, err
	}

	return id, nil
}

func (s *DevSignupStorage) GetPwByLogin(login string) (*core.DevPubAccPw, error) {
	sql := `SELECT id, password FROM dev_pub_account WHERE login = $1`

	devPubAccPw := core.DevPubAccPw{}

	if err := s.pool.QueryRow(context.Background(), sql, &login).Scan(&devPubAccPw.Id, &devPubAccPw.Password); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return nil, err
	}

	return &devPubAccPw, nil
}

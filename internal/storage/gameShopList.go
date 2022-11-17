package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/pkg/utils"
	"log"
)

type GameShopListStorage struct {
	pool *pgxpool.Pool
}

func NewGameShopListStorage(pool *pgxpool.Pool) *GameShopListStorage {
	return &GameShopListStorage{pool: pool}
}

func (s *GameShopListStorage) GetOne(id int) (*core.Game, error) {
	sql := `SELECT id, name, price, id_developers, 
       id_publishers, id_categories, system_requirements, 
       age_limit, description, release_date, version, rating FROM games
       WHERE id = $1`

	game := core.Game{}

	if err := s.pool.QueryRow(context.Background(), sql, id).Scan(
		&game.Id,
		&game.Name,
		&game.Price,
		&game.IdDevelopers,
		&game.IdPublishers,
		&game.IdCategories,
		&game.SystemReq,
		&game.AgeLimit,
		&game.Description,
		&game.ReleaseDate,
		&game.Version,
		&game.Rating,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return nil, err
	}

	return &game, nil
}

func (s *GameShopListStorage) GetAll(filter *core.GameFilter) ([]*core.Game, error) {
	sql := `SELECT id, name, price, id_developers, 
		id_publishers, id_categories, system_requirements, 
		age_limit, description, release_date, version, rating FROM games`

	games := make([]*core.Game, 0)

	rows, err := s.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		game := core.Game{}

		if err := rows.Scan(
			&game.Id,
			&game.Name,
			&game.Price,
			&game.IdDevelopers,
			&game.IdPublishers,
			&game.IdCategories,
			&game.SystemReq,
			&game.AgeLimit,
			&game.Description,
			&game.ReleaseDate,
			&game.Version,
			&game.Rating,
		); err != nil {
			return nil, err
		}

		games = append(games, &game)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameShopListStorage) Add(dto *core.CreateGame) error {
	sql := `INSERT INTO games(name, price, id_developers, id_publishers, id_categories, 
                  system_requirements, age_limit, description, release_date, version, rating)
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	if _, err := s.pool.Exec(context.Background(), sql,
		&dto.Name,
		&dto.Price,
		&dto.IdDevelopers,
		&dto.IdPublishers,
		&dto.IdCategories,
		&dto.SystemReq,
		&dto.AgeLimit,
		&dto.Description,
		&dto.ReleaseDate,
		&dto.Version,
		&dto.Rating,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return err
	}

	return nil
}

func (s *GameShopListStorage) Delete(id int) error {
	sql := `DELETE FROM games WHERE id = $1`

	if _, err := s.pool.Exec(context.Background(), sql, id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return err
	}

	return nil
}

func (s *GameShopListStorage) Update(id int, new *core.UpdateGameDTO) (*core.Game, error) {
	hgh := utils.NewTsk()

	if new.Name != nil {
		hgh.AddColumn("name", new.Name)
	}
	if new.Price != nil {
		hgh.AddColumn("price", new.Price)
	}
	if new.IdDevelopers != nil {
		hgh.AddColumn("id_developers", new.IdDevelopers)
	}
	if new.IdPublishers != nil {
		hgh.AddColumn("id_publishers", new.IdPublishers)
	}
	if new.IdCategories != nil {
		hgh.AddColumn("id_categories", new.IdCategories)
	}
	if new.SystemReq != nil {
		hgh.AddColumn("system_req", new.SystemReq)
	}
	if new.AgeLimit != nil {
		hgh.AddColumn("age_limit", new.AgeLimit)
	}
	if new.Description != nil {
		hgh.AddColumn("description", new.Description)
	}
	if new.ReleaseDate != nil {
		hgh.AddColumn("release_date", new.ReleaseDate)
	}
	if new.Version != nil {
		hgh.AddColumn("version", new.Version)
	}
	if new.Rating != nil {
		hgh.AddColumn("rating", new.Rating)
	}

	q := `UPDATE games SET ` + hgh.JoinColEl("id", id) + `RETURNING *`

	game := core.Game{}

	if err := s.pool.QueryRow(context.Background(), q, hgh.GetValue()...).Scan(
		&game.Id,
		&game.Name,
		&game.Price,
		&game.IdDevelopers,
		&game.IdPublishers,
		&game.IdCategories,
		&game.SystemReq,
		&game.AgeLimit,
		&game.Description,
		&game.ReleaseDate,
		&game.Version,
		&game.Rating,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return nil, err
	}

	return &game, nil
}

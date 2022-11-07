package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/pkg/utils"
	"log"
	"strings"
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

	if err := s.pool.QueryRow(context.Background(), sql, id).Scan(&game.Id, &game.Name,
		&game.Price, &game.IdDevelopers, &game.IdPublishers, &game.IdCategories, &game.SystemReq,
		&game.AgeLimit, &game.Description, &game.ReleaseDate, &game.Version, &game.Rating); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}
		log.Printf("[QUERY ERROR]: %v", err)
		return nil, err
	}

	return &game, nil
}

func (s *GameShopListStorage) GetAll() ([]*core.Game, error) {
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

		if err := rows.Scan(&game.Id, &game.Name, &game.Price, &game.IdDevelopers,
			&game.IdPublishers, &game.IdCategories, &game.SystemReq, &game.AgeLimit,
			&game.Description, &game.ReleaseDate, &game.Version, &game.Rating); err != nil {
			return nil, err
		}

		games = append(games, &game)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (s *GameShopListStorage) Add(dto *core.CreateGameDTO) error {
	sql := `INSERT INTO games(name, price, id_developers, id_publishers, id_categories, 
                  system_requirements, age_limit, description, release_date, version, rating)
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	if _, err := s.pool.Exec(context.Background(), sql, &dto.Name,
		&dto.Price, &dto.IdDevelopers, &dto.IdPublishers, &dto.IdCategories, &dto.SystemReq,
		&dto.AgeLimit, &dto.Description, &dto.ReleaseDate, &dto.Version, &dto.Rating); err != nil {
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

type tsk struct {
	column  []string
	element []interface{}
}

func newTsk() *tsk {
	return &tsk{
		column:  make([]string, 0),
		element: make([]interface{}, 0),
	}
}

func (tsk *tsk) addColumn(name string, value interface{}) {
	tsk.column = append(tsk.column, fmt.Sprintf("%s = $%d", name, len(tsk.column)+1))
	tsk.element = append(tsk.element, value)
}

func (tsk *tsk) joinColEl() string {

	return strings.Join(tsk.column, ", ")
}

func (tsk *tsk) getValue() []interface{} {
	return tsk.element
}

func (s *GameShopListStorage) Update(id int, new *core.UpdateGameDTO) (*core.Game, error) {
	hgh := newTsk()

	if new.Name != nil {
		hgh.addColumn("name", new.Name)
	}
	if new.Price != nil {
		hgh.addColumn("price", new.Price)
	}
	if new.AgeLimit != nil {
		hgh.addColumn("age_limit", new.AgeLimit)
	}

	fmt.Println(hgh.joinColEl())

	// TODO add where id
	q := `UPDATE games SET ` + hgh.joinColEl() + `RETURNING *`

	game := core.Game{}

	if err := s.pool.QueryRow(context.Background(), q, hgh.getValue()...).Scan(
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

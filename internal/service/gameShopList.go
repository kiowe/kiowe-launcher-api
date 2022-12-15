package service

import (
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/pkg/api/filter"
	"github.com/kiowe/kiowe-launcher-api/pkg/api/pagination"
	"github.com/kiowe/kiowe-launcher-api/pkg/api/sort"
	"math"
)

type GameShopListStorage interface {
	GetOne(id int) (*core.Game, error)
	GetAll(filterOptions []filter.Option, sortOption sort.Option,
		pag pagination.Pagination) ([]*core.Game, error)
	Add(dto *core.CreateGame) error
	Delete(id int) error
	Update(id int, game *core.UpdateGameDTO) (*core.Game, error)
	Count() (int, error)
}

type GameShopListService struct {
	storage GameShopListStorage
}

func NewGameShopListService(s GameShopListStorage) *GameShopListService {
	return &GameShopListService{storage: s}
}

func (s *GameShopListService) GetOne(id int) (*core.Game, error) {
	return s.storage.GetOne(id)
}

func (s *GameShopListService) GetAll(queryParams map[string]string) (*core.GamePage, error) {
	filterOptions := filter.GetFilterOptions(queryParams)
	sortOption := sort.GetSortOptions(queryParams)
	pag := pagination.GetPaginationOptions(queryParams)

	games, err := s.storage.GetAll(filterOptions, sortOption, pag)
	if err != nil {
		return nil, err
	}

	total, err := s.storage.Count()
	if err != nil {
		return nil, err
	}

	return &core.GamePage{
		Games:    games,
		Page:     pag.Page,
		Total:    total,
		LastPage: int(math.Ceil(float64(total / pag.PerPage))),
	}, nil

}

func (s *GameShopListService) Add(dto *core.CreateGame) error {

	// TODO: add validation
	return s.storage.Add(dto)
}

func (s *GameShopListService) Delete(id int) error {
	return s.storage.Delete(id)
}

func (s *GameShopListService) Update(id int, dto *core.UpdateGameDTO) (*core.Game, error) {
	return s.storage.Update(id, dto)
}

package service

import "github.com/kiowe/kiowe-launcher-api/internal/core"

type GameShopListStorage interface {
	GetOne(id int) (*core.Game, error)
	GetAll() ([]*core.Game, error)
	Add(dto *core.CreateGameDTO) error
	Delete(id int) error
	Update(id int, dto *core.UpdateGameDTO) (*core.Game, error)
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

func (s *GameShopListService) GetAll() ([]*core.Game, error) {
	return s.storage.GetAll()
}

func (s *GameShopListService) Add(dto *core.CreateGameDTO) error {

	// TODO: add validation
	return s.storage.Add(dto)
}

func (s *GameShopListService) Delete(id int) error {
	return s.storage.Delete(id)
}

func (s *GameShopListService) Update(id int, dto *core.UpdateGameDTO) (*core.Game, error) {
	return s.storage.Update(id, dto)
}

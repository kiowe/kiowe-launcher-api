package service

import "github.com/kiowe/kiowe-launcher-api/internal/core"

type GameShopListStorage interface {
	GetOne(id int) (*core.Game, error)
	GetAll() ([]*core.Game, error)
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

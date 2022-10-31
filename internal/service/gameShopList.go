package service

type GameShopListStorage interface {
	GetOne(id int) error
}

type GameShopListService struct {
	storage GameShopListStorage
}

func NewGameShopListService(s GameShopListStorage) *GameShopListService {
	return &GameShopListService{storage: s}
}

func (s *GameShopListService) GetOne(id int) error {
	return s.storage.GetOne(id)
}

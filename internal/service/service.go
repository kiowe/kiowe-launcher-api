package service

type Deps struct {
	GameShopListStorage GameShopListStorage
}

type Service struct {
	GameShopListService *GameShopListService
}

func New(deps Deps) *Service {
	return &Service{
		GameShopListService: NewGameShopListService(deps.GameShopListStorage),
	}
}

package service

type Deps struct {
	DevSignupStorage    DevSignupStorage
	GameShopListStorage GameShopListStorage
}

type Service struct {
	DevSignupService    *DevSignupService
	GameShopListService *GameShopListService
}

func New(deps Deps) *Service {
	return &Service{
		DevSignupService:    NewDevSignupService(deps.DevSignupStorage),
		GameShopListService: NewGameShopListService(deps.GameShopListStorage),
	}
}

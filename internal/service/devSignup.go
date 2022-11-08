package service

type DevSignupStorage interface {
}

type DevSignupService struct {
	storage DevSignupStorage
}

func NewDevSignupService(s DevSignupStorage) *DevSignupService {
	return &DevSignupService{storage: s}
}

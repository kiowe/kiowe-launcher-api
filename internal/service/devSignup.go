package service

import (
	"errors"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"golang.org/x/crypto/bcrypt"
)

type DevSignupStorage interface {
	Create(acc *core.DevPubAccountDTO) error
	GetByLogin(login string) (bool, error)
}

type DevSignupService struct {
	storage DevSignupStorage
}

func NewDevSignupService(s DevSignupStorage) *DevSignupService {
	return &DevSignupService{storage: s}
}

func (s *DevSignupService) Signup(dto *core.DevPubAccountDTO) (string, error) {
	acc, err := s.storage.GetByLogin(dto.Login)
	if err != nil {
		return "", nil
	}

	if acc {
		return "", errors.New("account already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	dto.Password = string(hash)

	if err := s.storage.Create(dto); err != nil {
		return "", nil
	}

	return "", nil
}

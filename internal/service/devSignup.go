package service

import (
	"errors"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type DevSignupStorage interface {
	Create(acc *core.DevPubAccountDTO) (int, error)
	GetByLogin(login string) (bool, error)
	GetPwByLogin(login string) (*core.DevPubAccPw, error)
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
		return "", err
	}

	if acc {
		return "", errors.New("account already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	dto.Password = string(hash)
	id, err := s.storage.Create(dto)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateNewAccessToken(id, true)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *DevSignupService) Signin(dto *core.LoginDevPubAccountDTO) (string, error) {
	acc, err := s.storage.GetPwByLogin(dto.Login)
	if err != nil {
		return "", errors.New("login not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(dto.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := middleware.GenerateNewAccessToken(acc.Id, true)
	if err != nil {
		return "", errors.New("cannot generate token")
	}

	return token, nil
}

package handler

import "github.com/gofiber/fiber/v2"

type DevSignupService interface {
}

type DevSignupHandler struct {
	service DevSignupService
}

func NewDevSignupHandler(s DevSignupService) *DevSignupHandler {
	return &DevSignupHandler{service: s}
}

func (h *DevSignupHandler) Signup(c *fiber.Ctx) error {
	return nil
}

func (h *DevSignupHandler) Signin(c *fiber.Ctx) error {
	return nil
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
)

type DevSignupService interface {
	Signup(dto *core.DevPubAccountDTO) (string, error)
	Signin(dto *core.LoginDevPubAccountDTO) (string, error)
}

type DevSignupHandler struct {
	service DevSignupService
}

func NewDevSignupHandler(s DevSignupService) *DevSignupHandler {
	return &DevSignupHandler{service: s}
}

func (h *DevSignupHandler) Signup(c *fiber.Ctx) error {
	devAcc := new(core.DevPubAccountDTO)

	if err := c.BodyParser(devAcc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"SIGNUP::[ERROR]": err.Error(),
		})
	}

	token, err := h.service.Signup(devAcc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"SIGNUP::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *DevSignupHandler) Signin(c *fiber.Ctx) error {
	devAcc := new(core.LoginDevPubAccountDTO)

	if err := c.BodyParser(devAcc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"SIGNIN::[ERROR]": err.Error(),
		})
	}

	if devAcc.Login == "" || devAcc.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"SIGNIN::[ERROR]": "Login or password cannot be empty",
		})
	}

	token, err := h.service.Signin(devAcc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"SIGNIN::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

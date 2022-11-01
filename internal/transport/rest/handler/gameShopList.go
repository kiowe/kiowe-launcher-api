package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
)

type GameShopListService interface {
	GetOne(id int) (*core.Game, error)
	GetAll() ([]*core.Game, error)
	Add(dto *core.CreateGameDTO) error
}

type GameShopListHandler struct {
	service GameShopListService
}

func NewGameShopListHandler(s GameShopListService) *GameShopListHandler {
	return &GameShopListHandler{service: s}
}

func (h *GameShopListHandler) GetOne(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	game, err := h.service.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(game)
}

func (h *GameShopListHandler) GetAll(c *fiber.Ctx) error {
	games, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(games)
}

func (h *GameShopListHandler) Add(c *fiber.Ctx) error {
	game := new(core.CreateGameDTO)

	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.Add(game); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Game was added.",
	})
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
)

type GameShopListService interface {
	GetOne(id int) (*core.Game, error)
	GetAll() ([]*core.Game, error)
	Add(dto *core.CreateGameDTO) error
	Delete(id int) error
	Update(id int, dto *core.UpdateGameDTO) (*core.Game, error)
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
			"GETONE::[ERROR]": err.Error(),
		})
	}

	game, err := h.service.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"GETONE::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(game)
}

func (h *GameShopListHandler) GetAll(c *fiber.Ctx) error {
	games, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"GETALL::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(games)
}

func (h *GameShopListHandler) Add(c *fiber.Ctx) error {
	game := new(core.CreateGameDTO)

	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	if err := h.service.Add(game); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ADD::[OK]": "Game was added.",
	})
}

func (h *GameShopListHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"DELETE::[OK]": "Game was deleted.",
	})
}

func (h *GameShopListHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"UPDATE::[ERROR]": err.Error(),
		})
	}

	game := new(core.UpdateGameDTO)

	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"UPDATE::[ERROR]": err.Error(),
		})
	}

	newGame, err := h.service.Update(id, game)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"UPDATE::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UPDATE::[OK]": "Game was updated",
		"New table:":   newGame,
	})
}

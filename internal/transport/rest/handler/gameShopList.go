package handler

import "github.com/gofiber/fiber/v2"

type GameShopListService interface {
	GetOne(id int) error
}

type GameShopListHandler struct {
	service GameShopListService
}

func NewGameShopListHandler(s GameShopListService) *GameShopListHandler {
	return &GameShopListHandler{service: s}
}

func (h *GameShopListHandler) GetOne(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "sdfsdf"})
}

package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Deps struct {
	GameShopListService GameShopListService
}

type Handler struct {
	app                 *fiber.App
	GameShopListHandler *GameShopListHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		GameShopListHandler: NewGameShopListHandler(deps.GameShopListService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/game/:id", h.GameShopListHandler.GetOne)
	return h.app
}

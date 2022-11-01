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
	game := v1.Group("/game")

	game.Get("/all", h.GameShopListHandler.GetAll)
	game.Get("/:id", h.GameShopListHandler.GetOne)
	game.Post("/add", h.GameShopListHandler.Add)
	game.Patch("/:id", h.GameShopListHandler.Update)
	game.Delete("/:id", h.GameShopListHandler.Delete)
	return h.app
}

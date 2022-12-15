package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/kiowe/kiowe-launcher-api/internal/middleware"
)

type Deps struct {
	DevSignupService    DevSignupService
	GameShopListService GameShopListService
}

type Handler struct {
	app                 *fiber.App
	DevSignupHandler    *DevSignupHandler
	GameShopListHandler *GameShopListHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		GameShopListHandler: NewGameShopListHandler(deps.GameShopListService),
		DevSignupHandler:    NewDevSignupHandler(deps.DevSignupService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	// Developer and publisher
	dev := v1.Group("/dev")
	dev.Post("/signup", h.DevSignupHandler.Signup)
	dev.Post("/signin", h.DevSignupHandler.Signin)

	// Games
	game := v1.Group("/game")
	game.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: middleware.JwtError,
	}))

	game.Get("/all", h.GameShopListHandler.GetAll)
	game.Get("/:id", h.GameShopListHandler.GetOne)
	game.Post("/add", h.GameShopListHandler.Add)
	game.Patch("/:id", h.GameShopListHandler.Update)
	game.Delete("/:id", h.GameShopListHandler.Delete)

	return h.app
}

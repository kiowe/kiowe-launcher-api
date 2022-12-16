package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiowe/kiowe-launcher-api/internal/core"
	"github.com/kiowe/kiowe-launcher-api/pkg/utils"
	"time"
)

type GameShopListService interface {
	GetOne(id int) (*core.Game, error)
	GetAll(queryParams map[string]string) (*core.GamePage, error)
	Add(dto *core.CreateGame) error
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
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"GETONE::[ERROR]": err.Error(),
		})
	}

	if now > claims.Expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"GETONE::[ERROR]": err.Error(),
		})
	}

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
	queryParams := utils.GetQueryParams(c, "name", "price", "id_developers", "id_publishers",
		"id_categories", "system_req", "age_limit", "description", "release_date", "rating", "page", "per_page",
		"sort_by", "sort_order")

	gamePage, err := h.service.GetAll(queryParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"GETALL::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gamePage)
}

func (h *GameShopListHandler) Add(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	if now > claims.Expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	if claims.DevPubAcc != true {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	game := new(core.CreateGameDTO)

	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	newGame := core.CreateGame{
		Name:         game.Name,
		Price:        game.Price,
		IdDevelopers: claims.Id,
		IdPublishers: claims.Id,
		IdCategories: game.IdCategories,
		SystemReq:    game.SystemReq,
		AgeLimit:     game.AgeLimit,
		Description:  game.Description,
		ReleaseDate:  game.ReleaseDate,
		Version:      game.Version,
		Rating:       game.Rating,
	}

	if err := h.service.Add(&newGame); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ADD::[ERROR]": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ADD::[OK]": "Game was added.",
	})
}

func (h *GameShopListHandler) Delete(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	if now > claims.Expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	if claims.DevPubAcc != true {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

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
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	if now > claims.Expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

	if claims.DevPubAcc != true {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"DELETE::[ERROR]": err.Error(),
		})
	}

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

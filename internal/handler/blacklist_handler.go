package handler

import (
	"fww-regulation/internal/model"
	"fww-regulation/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Blacklist struct {
	BlacklistUsecase usecase.BlacklistExecutor
}

type BlacklistHandler interface {
	CheckBlacklist(c *fiber.Ctx) error
}

func NewBlacklistHandler(handler Blacklist) BlacklistHandler {
	return &handler
}

func (handler *Blacklist) CheckBlacklist(c *fiber.Ctx) error {
	var request model.BlacklistRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := handler.BlacklistUsecase.CheckBlacklist(request.KTP)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if !ok {
		return c.JSON(fiber.Map{"status": true})
	}

	return c.JSON(fiber.Map{"status": false})
}

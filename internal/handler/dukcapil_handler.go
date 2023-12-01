package handler

import (
	"fww-regulation/internal/model"
	"fww-regulation/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Dukcapil struct {
	DukcapilUsecase usecase.DukcapilExecutor
}

type DukcapilHandler interface {
	CheckDukcapilByKTP(c *fiber.Ctx) error
}

func NewDukcapilHandler(handler Dukcapil) DukcapilHandler {
	return &handler
}

// @Summary Check KTP status at Dukcapil
// @Description	Check KTP status at Dukcapil
// @Tags dukcapil
// @Accept json
// @Produce	json
// @Param payload body model.DukcapilRequest true "DukcapilRequest"
// @Success 200 {object} model.StatusResponse "OK"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Router /check/dukcapil [get]
func (handler *Dukcapil) CheckDukcapilByKTP(c *fiber.Ctx) error {
	var req model.DukcapilRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	status, err := handler.DukcapilUsecase.CheckDukcapilByKTP(req.Ktp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": status})
}

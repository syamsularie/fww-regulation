package handler

import (
	"fww-regulation/internal/model"
	"fww-regulation/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type PeduliLindungi struct {
	PeduliLindungiUsecase usecase.PeduliLindungiExecutor
}

type PeduliLindungiHandler interface {
	CheckPeduliLindungi(c *fiber.Ctx) error
}

func NewPeduliLindungiHandler(handler PeduliLindungi) PeduliLindungiHandler {
	return &handler
}

// @Summary Check KTP status at Peduli Lindungi
// @Description	Check KTP status at Peduli Lindungi
// @Tags peduli lindungi
// @Accept json
// @Produce	json
// @Param payload body model.PeduliLindungiRequest true "PeduliLindungiRequest"
// @Success 200 {object} model.StatusResponse "OK"
// @Failure 400 {object} model.ErrorResponse "Bad Request"
// @Router /check/pedulilindungi [get]
func (handler *PeduliLindungi) CheckPeduliLindungi(c *fiber.Ctx) error {
	var request model.PeduliLindungiRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	status, err := handler.PeduliLindungiUsecase.CheckPeduliLindungi(request.Ktp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": status})
}

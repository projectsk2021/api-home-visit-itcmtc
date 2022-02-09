package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	services "github.com/kamchai-n/api-student-home-visit/services/visit"
	"github.com/kamchai-n/api-student-home-visit/utils"
)

type visitHandler struct {
	visitService services.VisitService
}

func NewVisitHandler(visitService services.VisitService) VisitHandler {
	return visitHandler{visitService: visitService}
}
func (h visitHandler) ListAllVisit(c *fiber.Ctx) error {
	visit, err := h.visitService.ListAllVisit()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "data_found",
		Data: visit,
	}))
}

func (h visitHandler) ListDetailById(c *fiber.Ctx) error {
	id := c.Params("id")
	visitId := uuid.MustParse(id)

	visit, err := h.visitService.ListDetailById(&visitId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "data_found",
		Data: visit,
	}))
}

func (h visitHandler) SaveVisit(c *fiber.Ctx) error {
	var req domain.RequestVisit
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	visit, err := h.visitService.NewVisit(req)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "created_successfully",
		Data: visit,
	}))
}

func (h visitHandler) EditVisit(c *fiber.Ctx) error {
	var req domain.RequestUpdateVisitForm
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	if err := h.visitService.UpdateVisit(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "update_successfully",
	}))
}

func (h visitHandler) DeleteVisit(c *fiber.Ctx) error {
	id := c.Params("visit_id")
	visitId := uuid.MustParse(id)

	if err := h.visitService.DeleteVisit(&visitId); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "deleted_successfully",
	}))
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	services "github.com/kamchai-n/api-student-home-visit/services/student"
	"github.com/kamchai-n/api-student-home-visit/utils"
)

type studentHandler struct {
	stdService services.StudentService
}

func NewStudentHandler(stdService services.StudentService) StudentHandler {
	return studentHandler{stdService: stdService}
}

func (s studentHandler) SaveStudentFromExcel(c *fiber.Ctx) error {
	dataFile, err := c.FormFile("dataFile")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1301",
		}))
	}

	response, err := s.stdService.NewStudentFromExcel(dataFile, *middlewares.UserClaims.UserId)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "created_successfully",
		Data: response,
	}))
}

func (s studentHandler) SaveStudent(c *fiber.Ctx) error {
	var req domain.StudentForm
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	response, err := s.stdService.NewStudent(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "created_successfully",
		Data: response,
	}))
}

func (s studentHandler) ShowInfoStudent(c *fiber.Ctx) error {
	var req *domain.RequestGetAll
	c.BodyParser(&req)
	// if err := c.BodyParser(&req); err != nil {
	// 	return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
	// 		ErrorCode: "1309",
	// 	}))
	// }

	response, err := s.stdService.ListInfoStudent(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code:  "data_found",
		Count: len(response),
		Data:  response,
	}))
}

func (s studentHandler) ShowInfoStudentById(c *fiber.Ctx) error {
	var req domain.RequestGetOneStudent
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	response, err := s.stdService.ListInfoStudentById(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "data_found",
		Data: response,
	}))
}

func (s studentHandler) DeleteStudentById(c *fiber.Ctx) error {
	var req domain.RequestDeleteStudent
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	if err := s.stdService.RemoveStudentById(req.UserID); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "deleted_successfully",
	}))
}

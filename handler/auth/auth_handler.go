package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	"github.com/kamchai-n/api-student-home-visit/model"
	services "github.com/kamchai-n/api-student-home-visit/services/auth"
	"github.com/kamchai-n/api-student-home-visit/utils"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return authHandler{authService: authService}
}

func (s authHandler) UserRegister(c *fiber.Ctx) error {
	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	_, err := s.authService.NewUser(req)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "created_successfully",
	}))
}

func (s authHandler) UserLogin(c *fiber.Ctx) error {
	var req domain.UserLoginForm
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	token, err := s.authService.UserLogin(req)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code:  "auth_success",
		Token: token,
	}))
}
func (s authHandler) GetAllUser(c *fiber.Ctx) error {
	user, err := s.authService.ShowUserAll()
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "data_found",
		Data: user,
	}))
}

func (s authHandler) UserInfo(c *fiber.Ctx) error {
	user, err := s.authService.ShowUser(*middlewares.UserClaims.UserId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "data_found",
		User: user,
	}))
}

func (s authHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := uuid.MustParse(id)
	// if err != nil {
	// 	return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
	// 		ErrorCode: "1320",
	// 	}))
	// }

	user, err := s.authService.ShowUser(userId)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	if err := s.authService.RemoveUser(user.UserID); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "deleted_successfully",
	}))
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	services "github.com/kamchai-n/api-student-home-visit/services/setting"
	"github.com/kamchai-n/api-student-home-visit/utils"
)

type settingHandler struct {
	settingService services.SettingService
}

func NewSettingHandler(settingService services.SettingService) SettingHandler {
	return settingHandler{settingService: settingService}
}

func (s settingHandler) ShowSetting(c *fiber.Ctx) error {
	user, err := s.settingService.ListSetting(*middlewares.UserClaims.UserId)
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

func (s settingHandler) EditIsFromSdcSetting(c *fiber.Ctx) error {
	var req domain.RequestIsFromSdcUpdateSetting
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: "1309",
		}))
	}

	req.ID = *middlewares.UserClaims.UserId

	if err := s.settingService.UpdateIsFromSdcSetting(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(utils.ErrorResponse(utils.Message{
			ErrorCode: err.Error(),
		}))
	}

	return c.JSON(utils.SuccessResponse(utils.Message{
		Code: "update_successfully",
	}))

}

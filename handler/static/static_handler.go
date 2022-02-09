package handler

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/kamchai-n/api-student-home-visit/services/static"
	"github.com/kamchai-n/api-student-home-visit/utils"
)

type staticHandler struct {
	staticService services.StaticService
}

func NewStaticHandler(staticService services.StaticService) StaticHandler {
	return staticHandler{staticService: staticService}
}

func (s staticHandler) ShowEducation(c *fiber.Ctx) error {
	user, err := s.staticService.ListEducation()
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

package handler

import "github.com/gofiber/fiber/v2"

type SettingHandler interface {
	ShowSetting(c *fiber.Ctx) error
	EditIsFromSdcSetting(c *fiber.Ctx) error
}

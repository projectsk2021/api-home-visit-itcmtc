package handler

import "github.com/gofiber/fiber/v2"

type StaticHandler interface {
	ShowEducation(c *fiber.Ctx) error
}

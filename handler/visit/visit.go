package handler

import "github.com/gofiber/fiber/v2"

type VisitHandler interface {
	ListDetailById(c *fiber.Ctx) error
	ListAllVisit(c *fiber.Ctx) error
	SaveVisit(c *fiber.Ctx) error
	EditVisit(c *fiber.Ctx) error
	DeleteVisit(c *fiber.Ctx) error
}

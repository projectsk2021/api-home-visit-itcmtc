package handler

import "github.com/gofiber/fiber/v2"

type StudentHandler interface {
	SaveStudent(c *fiber.Ctx) error
	SaveStudentFromExcel(c *fiber.Ctx) error
	ShowInfoStudent(c *fiber.Ctx) error
	ShowInfoStudentById(c *fiber.Ctx) error
	DeleteStudentById(c *fiber.Ctx) error
}

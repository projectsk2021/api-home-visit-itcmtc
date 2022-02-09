package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	UserRegister(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
	UserInfo(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

package logic

import (
	"github.com/gofiber/fiber/v2"
)

func AdminTest(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": true,
		"ADMIN":   "LOGGED IN VIA PACKAGERRRR",
	})
}
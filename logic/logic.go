package logic

import (
	"github.com/gofiber/fiber/v2"
	config "example.com/config"
)

func AdminTest(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": true,
		"ADMIN":   "LOGGED IN VIA PACKAGERRRR",
	})
}

func AddAUser(c *fiber.Ctx) error {
	config.AddUser(c.Params("Username"),c.Params("Password"));
	return c.JSON(&fiber.Map{
		"success": true,
		"ADMIN":   "LOGGED IN VIA PACKAGERRRR",
	})
}


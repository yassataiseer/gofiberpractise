package handlers

import (
	"github.com/gofiber/fiber/v2"
	"example.com/logic"
)


func Admin(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": true,
		"ADMIN":   "LOGGED IN",
	})
}
func Register(app *fiber.App) {
  app.Get("/admin",  logic.AdminTest)

}


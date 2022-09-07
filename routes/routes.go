package routes 

import (
    "github.com/gofiber/fiber/v2"
)



func Admin(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": true,
		"ADMIN":   "LOGGED IN",
	})
}
func Register(app *fiber.App) {
  app.Get("/admin",  Admin)

}

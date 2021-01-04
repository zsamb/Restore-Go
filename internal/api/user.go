package api

import "github.com/gofiber/fiber/v2"

func User(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})
}

package routers

import "github.com/gofiber/fiber/v2"

func InitWebRouter(server *fiber.App) {
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
}

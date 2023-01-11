package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/fiber-gorm-mysql.git/config"
)

func main() {

	app := fiber.New()

	config.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello")
		return err
	})

	app.Listen(":3000")

}

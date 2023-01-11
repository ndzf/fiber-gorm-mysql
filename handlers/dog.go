package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ndzf/fiber-gorm-mysql.git/config"
	"github.com/ndzf/fiber-gorm-mysql.git/entities"
)

func GetDogs(c *fiber.Ctx) error {

	var dogs []entities.Dog

	config.Database.Take(&dogs)
	return c.JSON(dogs)
}

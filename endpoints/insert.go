package endpoints

import (
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// Insert inserts a record into the database
func Insert(c *fiber.Ctx, db *gorm.DB) error {
	satResult := new(types.SATResult)
	if err := c.BodyParser(satResult); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	satResult.Passed = satResult.SATScore > 30

	db.Create(&satResult)
	return c.Status(fiber.StatusCreated).JSON(satResult)

}

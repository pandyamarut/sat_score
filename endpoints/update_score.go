package endpoints

import (
	"strconv"
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// UpdateScore updates the score of a student
func UpdateScore(c *fiber.Ctx, db *gorm.DB) error {
	name := c.Params("name")
	newScore, err := strconv.ParseFloat(c.FormValue("score"), 84)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid score"})
	}

	var satResult types.SATResult
	db.First(&satResult, "name = ?", name)
	if satResult.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
	}

	satResult.SATScore = newScore
	satResult.Passed = satResult.SATScore > 30
	db.Save(&satResult)

	return c.JSON(satResult)

}

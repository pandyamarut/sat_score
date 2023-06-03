package endpoints

import (
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// Delete deletes a record from the database
func Delete(c *fiber.Ctx, db *gorm.DB) error {
	name := c.Params("name")

	var satResult types.SATResult
	// Finds the record based on the name and deletes it
	db.Model(&types.SATResult{}).Where("name = ?", name).First(&satResult)
	if satResult.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found", "NAME": &satResult.Name})
	}

	db.Delete(&satResult)
	return c.JSON(fiber.Map{"message": "Record deleted successfully", "NAME": &satResult.Name})

}

package endpoints

import (
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// GetRank returns the rank of a student
func GetRank(c *fiber.Ctx, db *gorm.DB) error {
	name := c.Params("name")
	var count int
	db.Model(&types.SATResult{}).Where("name < ?", name).Count(&count)
	return c.JSON(fiber.Map{"rank": count + 1})

}

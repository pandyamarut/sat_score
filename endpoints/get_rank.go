package endpoints

import (
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// GetRank returns the rank of the entry with the given name
func GetRank(c *fiber.Ctx, db *gorm.DB) error {
	name := c.Params("name")
	var count int
	var satResult types.SATResult
	//searched for the entry with the name and then found the count of the entries with a higher score
	db.Model(&types.SATResult{}).Where("name = ?", name).First(&satResult)
	// counts the number of rows with a higher score relative to the scoe of the name passed.
	db.Model(&types.SATResult{}).Where("sat_score > ?", satResult.SATScore).Count(&count)
	rank := count + 1 // adds one to the count to get the rank
	return c.JSON(fiber.Map{"rank==": rank})
}

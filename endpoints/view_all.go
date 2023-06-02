package endpoints

import (
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// ViewAll returns all the data in the database
func ViewAll(c *fiber.Ctx, db *gorm.DB) error {
	var satResults []types.SATResult
	db.Find(&satResults)
	return c.JSON(satResults)

}

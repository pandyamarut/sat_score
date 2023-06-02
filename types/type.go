package types

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SATResult is the database model
type SATResult struct {
	gorm.Model
	Name     string
	Address  string
	SATScore float64
	Passed   bool
}

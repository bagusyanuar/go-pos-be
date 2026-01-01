package seeder

import "gorm.io/gorm"

func Seed(db *gorm.DB) {
	DeveloperSeeder(db)
}

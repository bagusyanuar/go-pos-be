package seeder

import (
	"errors"
	"log"

	"github.com/bagusyanuar/go-pos-be/internal/shared/entity"
	"gorm.io/gorm"
)

func DeveloperSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("users") {
		log.Println("⛔ Table users not found, seeding cancelled.")
		return
	}

	email := "developer@dev.id"
	username := "developer"
	password := "$2a$13$IkJBTwgSdi/kuNn4ndk4cO0pJ2Ov.qaKXjojrl91l4hKOLtll3WY."

	data := entity.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	var user *entity.User
	err := db.Where("email = ?", email).First(&user).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		if err := db.Create(&data).Error; err != nil {
			log.Printf("❌ failed to create new developer seed: %v", err)
			return
		}
		log.Println("✅ successfully create new developer seed")
	case err != nil:
		log.Printf("❌ failed to execute query developer seed : %v", err)
		return
	default:
		if errUpdate := db.Model(&user).Updates(&data).Error; errUpdate != nil {
			log.Printf("❌ failed to update developer seed : %v", err)
			return
		}
		log.Println("✅ successfully update developer seed")
	}
}

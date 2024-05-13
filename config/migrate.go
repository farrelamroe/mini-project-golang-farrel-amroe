package config

import (
	"gorm.io/gorm"
	"project/models"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		return err
	}
	return nil
}
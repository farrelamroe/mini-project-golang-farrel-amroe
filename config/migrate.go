package config

import (
	"project/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}, &model.Book{}, &model.Categories{}); err != nil {
		return err
	}
	return nil
}

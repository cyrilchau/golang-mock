package repository

import (
	"myapp/internal/user/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	return
}

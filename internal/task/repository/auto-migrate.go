package repository

import (
	"myapp/internal/task/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Task{})
	if err != nil {
		panic(err)
	}

	return
}

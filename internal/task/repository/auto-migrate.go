package repository

import (
	"myapp/internal/worker/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Worker{})
	if err != nil {
		panic(err)
	}

	return
}

package gorm

import (
	models "user-info-service/external/repository/adaptors/postgres/model"

	"gorm.io/gorm"
)

func SyncDB(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}

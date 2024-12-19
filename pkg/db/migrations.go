package db

import (
	"github.com/teakingwang/gin-mysql/internal/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}

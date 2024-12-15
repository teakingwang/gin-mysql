package repository

import (
	"gorm.io/gorm"
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Migrate() error {
	return repo.db.AutoMigrate(&models.User{})
}

func (repo *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	result := repo.db.Create(user)
	return result.Error
}

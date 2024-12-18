package service

import (
	"github.com/teakingwang/gin-mysql/config"
	"github.com/teakingwang/gin-mysql/internal/models"
	"github.com/teakingwang/gin-mysql/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	db, err := gorm.Open(mysql.Open(config.LoadConfig().Database.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepo := repository.NewUserRepo(db)
	if err := userRepo.Migrate(); err != nil {
		panic("failed to migrate database")
	}
	return &UserService{userRepo: userRepo}
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.userRepo.GetAllUsers()
}

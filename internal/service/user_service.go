package service

import (
	"gorm.io/gorm"
	"my-gin-app/config"
	"my-gin-app/internal/repository"
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

func (service *UserService) CreateUser(user *models.User) error {
	return service.userRepo.CreateUser(user)
}

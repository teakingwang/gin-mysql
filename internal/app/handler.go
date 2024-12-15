package app

import (
	"github.com/gin-gonic/gin"
	"my-gin-app/internal/service"
	"net/http"
)

var userService = service.NewUserService()

// GetUserList godoc
// @Summary Get user list
// @Description get user list
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} string "{"message":"success"}"
// @Router /api/v1/users [get]
func GetUserList(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create user
// @Description create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body my-gin-app/internal/models.User true "User object"
// @Success 201 {object} my-gin-app/internal/models.User
// @Router /api/v1/users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-mysql/internal/service"
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

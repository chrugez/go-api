package controller

import (
	"net/http"

	"github.com/chrugez/go-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController{
	return &UserController{
		userService: service.NewUserService(),
	}
}

//uc user controller
//us user service

//controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// name := c.Param("name")
	// name := c.DefaultQuery("name", "chrugez")
	// uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"user1", "user2", "user3"},
	})
}
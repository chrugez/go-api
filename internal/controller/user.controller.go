package controller

import (
	"github.com/chrugez/go-api/internal/service"
	"github.com/chrugez/go-api/pkg/response"
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
	// c.JSON(http.StatusOK, response.ResponseData{
	// 	Code: 20001,
	// 	Message: "success",
	// 	Data: []string{"user1", "user2"},
	// })
	// response.SuccessResponse(c, 20001, []string{"user1", "user2", "user3"})
	response.ErrorResponse(c, 20003, "No need!")
}
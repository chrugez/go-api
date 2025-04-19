package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController{
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("--->My Handler")
	// name := c.Param("name")
	name := c.DefaultQuery("name", "chrugez")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong" + name,
		"uid":     uid,
		"users":   []string{"user1", "user2", "user3"},
	})
}
package user

import (
	"github.com/chrugez/go-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	userController, _ := wire.InitUserRouterHandler()
	//public router
	UserRouterPublic := Router.Group("/user")
	{
		UserRouterPublic.POST("/register", userController.Register)
		UserRouterPublic.POST("/otp")
	}
	//private router
	UserRouterPrivate := Router.Group("/user")
	// UserRouterPrivate.Use(Limiter())
	// UserRouterPrivate.Use(Authen())
	// UserRouterPrivate.Use(Permission())
	{
		UserRouterPrivate.GET("/get_info")
	}
}
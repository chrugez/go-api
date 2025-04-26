package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	//public router
	UserRouterPublic := Router.Group("/user")
	{
		UserRouterPublic.POST("/register")
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
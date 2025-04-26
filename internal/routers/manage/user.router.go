package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	//public router
	// UserRouterPublic := Router.Group("/admin/user")
	// {
	// 	UserRouterPublic.POST("/register")
	// 	UserRouterPublic.POST("/otp")
	// }
	//private router
	UserRouterPrivate := Router.Group("/admin/user")
	// UserRouterPrivate.Use(Limiter())
	// UserRouterPrivate.Use(Authen())
	// UserRouterPrivate.Use(Permission())
	{
		UserRouterPrivate.POST("/active_user")
	}
}
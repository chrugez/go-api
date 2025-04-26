package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup){
	//public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	//private router
	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(Limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.GET("/get_info")
	}
}
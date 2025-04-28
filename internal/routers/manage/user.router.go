package manage

import (
	"github.com/chrugez/go-api/internal/controller"
	"github.com/chrugez/go-api/internal/repo"
	"github.com/chrugez/go-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	//public router
	//this is non-dependency
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)
	UserRouterPublic := Router.Group("/admin/user")

	// WIRE go
	// Dependency Injection (DI)
	{
		UserRouterPublic.POST("/register", userHandlerNonDependency.Register)
		// UserRouterPublic.POST("/otp")
	}
	//private router
	UserRouterPrivate := Router.Group("/admin/user")
	// UserRouterPrivate.Use(Limiter())
	// UserRouterPrivate.Use(Authen())
	// UserRouterPrivate.Use(Permission())
	{
		UserRouterPrivate.POST("/active_user")
	}
}
package routers

import (
	"github.com/chrugez/go-api/internal/routers/manage"
	"github.com/chrugez/go-api/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
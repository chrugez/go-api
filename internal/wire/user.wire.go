//go:build wireinject

package wire

import (
	"github.com/chrugez/go-api/internal/controller"
	"github.com/chrugez/go-api/internal/repo"
	"github.com/chrugez/go-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
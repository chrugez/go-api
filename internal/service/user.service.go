package service

import (
	"github.com/chrugez/go-api/internal/repo"
	"github.com/chrugez/go-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

//INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
	//...
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	//1. Check email exists
	if us.userRepo.GetUserByEmail(email){
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
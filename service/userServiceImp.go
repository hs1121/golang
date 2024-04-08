package service

import (
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
)

// UserService implements the IUserService interface.
type UserServiceImp struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) UserService {
	return &UserServiceImp{
		userRepo: userRepo,
	}
}


func (s *UserServiceImp) CreateUser(user *model.User) error {
	// Here you can add any business logic before saving the user
	return s.userRepo.CreateUser(user)
}
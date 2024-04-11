package service

import (
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
)

type UserServiceImp struct {
	userRepo repository.UserRepository 
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImp{
		userRepo: userRepo,
	}
}

func (s *UserServiceImp) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *UserServiceImp) UpdateUser(user *model.User) (*model.User, error) {
	return s.userRepo.UpdateUser(user)
}

func (s *UserServiceImp) GetUserByID(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserServiceImp) GetUsers() (*[]model.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserServiceImp) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}

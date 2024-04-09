package service

import (
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
)

// UserService implements the IUserService interface.
type UserServiceImp struct {
	userRepo *repository.UserRepository
}

func  NewUserService(userRepo *repository.UserRepository) *UserServiceImp {
	return &UserServiceImp{
		userRepo: userRepo,
	}
}

func (ur *UserServiceImp) CreateUser(user *model.User) (*model.User, error){
	return nil,nil
}
func (ur *UserServiceImp)UpdateUser(user *model.User) (*model.User, error){
	return nil,nil
}
func (ur *UserServiceImp)GetUserByID(id int) (*model.User, error){
	return nil,nil
}
func (ur *UserServiceImp)GetUsers() (*[]model.User, error){
	return nil,nil
}
func (ur *UserServiceImp)DeleteUser(user *model.User) error{
	return nil
}
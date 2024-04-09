package service

import (
	"github.com/hs1121/user/model"
)

// IUserService outlines the methods that the user service will implement.
type UserService interface {
    CreateUser(user *model.User) (*model.User, error)
    UpdateUser(user *model.User) (*model.User, error)
    GetUserByID(id int) (*model.User, error)
    GetUsers() (*[]model.User, error)
    DeleteUser(user *model.User) error
}

package service

import "github.com/hs1121/user/model"

// IUserService outlines the methods that the user service will implement.
type UserService interface {
	CreateUser(user *model.User) error
}

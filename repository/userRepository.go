package repository

import (
	"github.com/hs1121/user/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	DeleteUser(id int) error
	GetAllUsers() (*[]model.User, error)
}

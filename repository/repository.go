package repository

import (
	"github.com/hs1121/user/database"
	"github.com/hs1121/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.Connect(),
	}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id int64) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	return user, result.Error
}

// Add more methods as needed: UpdateUser, DeleteUser, ListUsers, etc.

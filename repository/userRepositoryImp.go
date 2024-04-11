package repository

import (
	"github.com/hs1121/user/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) CreateUser(user *model.User) (*model.User, error) {
	result := repo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *GormUserRepository) UpdateUser(user *model.User) (*model.User, error) {
	result := repo.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *GormUserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *GormUserRepository) DeleteUser(id int) error {
	result := repo.db.Delete(&model.User{}, id)
	return result.Error
}

func (repo *GormUserRepository) GetAllUsers() (*[]model.User, error) {
	var users []model.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

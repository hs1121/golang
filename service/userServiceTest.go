package service

import (
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
)

// UserService implements the IUserService interface.
type UserServiceTest struct {
	repository *repository.UserRepositoryTest
}

func NewUserServiceTest(repository *repository.UserRepositoryTest) *UserServiceTest {
	return &UserServiceTest{
		repository: repository,
	}
}

func (u *UserServiceTest) CreateUser(user *model.User) (*model.User, error) {
	user.ID = u.repository.GetNewId()
	return u.repository.CreateUser(user)
}
func (u *UserServiceTest) UpdateUser(user *model.User) (*model.User, error) {
	return u.repository.UpdateUser(user)
}
func (u *UserServiceTest) GetUserByID(id int) (*model.User, error) {
	return u.repository.GetUserByID(id)
}
func (u *UserServiceTest) GetUsers() (*[]model.User, error) {
	return u.repository.GetAllUsers()
}
func (u *UserServiceTest) DeleteUser(id int) error {
	return u.repository.DeleteUser(id)
}

package repository

import (
	"errors"
	"github.com/hs1121/user/model"
)

type UserRepositoryTest struct {
	database []model.User
	newId    int
}

func NewUserRepositoryTest() *UserRepositoryTest {
	return &UserRepositoryTest{
		database: []model.User{},
		newId:    0,
	}
}

func (ur *UserRepositoryTest) GetNewId() int {
	ur.newId++
	return ur.newId
}

func (ur *UserRepositoryTest) CreateUser(userToAdd *model.User)  (*model.User, error) {
	ur.database = append(ur.database, *userToAdd)

	for index, user := range ur.database {
		if userToAdd.ID == user.ID {
			ur.database[index] = *userToAdd
			return userToAdd,nil
		}
	}	
	return nil,errors.New("unable to add user")
}

func (ur *UserRepositoryTest) UpdateUser(userToUpdate *model.User)  (*model.User, error) {

	for index, user := range ur.database {
		if userToUpdate.ID == user.ID {
			ur.database[index] = *userToUpdate
			return  &ur.database[index], nil
		}
	}
	return nil,errors.New("entry not found")
}

func (ur *UserRepositoryTest) GetUserByID( id int) (*model.User, error) {
	for _, user := range ur.database {
		if id == user.ID {
			return &user,nil
		}
	}
	return nil,errors.New("id not found")
}

func (ur *UserRepositoryTest) DeleteUser(id int) error {
	var indexToRemove int =-1
	for index, user := range ur.database {
		if id == user.ID {
			indexToRemove=index
			break
		}
	}

	if indexToRemove != -1 {
        ur.database = append(ur.database[:indexToRemove], ur.database[indexToRemove+1:]...)
        return nil
    } else {
        return errors.New("index not found")
    }
}

func (ur *UserRepositoryTest) GetAllUsers() (*[]model.User,error) {
	return &ur.database,nil
}

package main

import (
	"fmt"

	"github.com/hs1121/user/database"
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
	"github.com/hs1121/user/service"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(&model.User{})

	repository := repository.NewUserRepository()
	service := service.NewUserService(repository)
}
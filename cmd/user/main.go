package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hs1121/user/controller"
	"github.com/hs1121/user/database"
	"github.com/hs1121/user/model"
	"github.com/hs1121/user/repository"
	"github.com/hs1121/user/service"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {

	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	db := database.Connect(*cfg)
	// AutoMigrate is used to ensure the table exists for the User model.
	db.AutoMigrate(&model.User{})

	router := gin.Default()

	repository := repository.NewGormUserRepository(db)
	var userService service.UserService = service.NewUserService(repository)
	userController := controller.NewUserController(userService)

	router.GET("/user", userController.GetAllUsers)
	router.POST("/user", userController.CreateUser)
	router.GET("/user/:id", userController.GetUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	router.Run()
}

func initConfig() (*model.Config, error) {
	var cfg model.Config

	data, err := os.ReadFile("../../config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

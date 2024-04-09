package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hs1121/user/controller"
	"github.com/hs1121/user/repository"
	"github.com/hs1121/user/service"
)

func main() {
 router := gin.Default()

 repository := repository.NewUserRepositoryTest()
 var userService service.UserService  = service.NewUserServiceTest(repository)
 userController := controller.NewUserController(userService)

  router.GET("/user", userController.GetAllUsers)
 router.POST("/user", userController.CreateUser)
 router.GET("/user/:id", userController.GetUser)
 router.PUT("/user/:id", userController.UpdateUser)
 router.DELETE("/user/:id", userController.DeleteUser)


 router.Run()
}
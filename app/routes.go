package app

import (
	controller "golang-blueprint/app/controllers"
	"golang-blueprint/app/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (server *Server) initializeRoutes(db *gorm.DB) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := controller.NewUserController(userService)

	server.Router = gin.Default()
	api := server.Router.Group("v1")

	api.GET("/", controller.Home)
	api.POST("/register", userController.RegisterUser)
}

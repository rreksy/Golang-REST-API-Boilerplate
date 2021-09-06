package app

import (
	controller "golang-migration/app/controllers"

	"github.com/gin-gonic/gin"
)

func (server *Server) initializeRoutes() {
	server.Router = gin.Default()

	server.Router.GET("/", controller.Home)
}

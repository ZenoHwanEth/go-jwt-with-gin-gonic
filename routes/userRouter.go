package routes

import (
	controller "github.com/ZenoHwanEth/go-jwt-with-gin-gonic/controllers"

	"github.com/ZenoHwanEth/go-jwt-with-gin-gonic/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

package routes

import (
	controller "github.com/barangayretli/kubermario-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.POST("users/signup", controller.Signup())

}

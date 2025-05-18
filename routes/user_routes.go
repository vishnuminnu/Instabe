package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
)

func RegisterUserRoutes(router *gin.Engine){
	r:=router.Group("/api/users")
	r.POST("/signup",controllers.Signup)
	r.POST("/login" ,controllers.Login)
}
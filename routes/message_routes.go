package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
	"github.com/vishnu/instabe/middlewares"
)

func MessageRoutes(r *gin.Engine) {
	msg := r.Group("/api/messages")
	msg.Use(middlewares.JWTAuthMiddleware())
	{
		msg.POST("/send", controllers.SendMessage)
		msg.GET("/chat", controllers.GetChat)
	}
}
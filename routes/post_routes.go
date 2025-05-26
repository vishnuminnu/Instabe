package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
	"github.com/vishnu/instabe/middlewares"
)

func RegisterPostRoutes(router *gin.Engine) {
	r := router.Group("/api/posts")
	r.Use(middlewares.JWTAuthMiddleware())
	r.POST("/create", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
}

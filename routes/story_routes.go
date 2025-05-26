package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
	"github.com/vishnu/instabe/middlewares"
)

func StoryRoutes(r *gin.Engine) {
	storyGroup := r.Group("/api/stories").Use(middlewares.JWTAuthMiddleware())
	{
		storyGroup.POST("/", controllers.CreateStory)
		storyGroup.GET("/:user_id", controllers.GetUserStories)
	}
}
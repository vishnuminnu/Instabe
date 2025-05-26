package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
)

func CommentRoutes(router *gin.Engine) {
	comments := router.Group("/api/comments")
	{
		comments.POST("/", controllers.CreateComment)
		comments.GET("/:postId", controllers.GetCommentsByPostID)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu/instabe/controllers"
	"github.com/vishnu/instabe/middlewares"
)

func RegisterUserRoutes(router *gin.Engine){
	r:=router.Group("/api/users")
	r.POST("/signup",controllers.Signup)
	r.POST("/login" ,controllers.Login)
	r.GET("/", controllers.GetAllUsers)
	r.POST("/:id/follow", middlewares.JWTAuthMiddleware(), controllers.FollowUser)
	r.POST("/:id/unfollow", middlewares.JWTAuthMiddleware(), controllers.UnfollowUser)
	r.GET("/:id/followers", middlewares.JWTAuthMiddleware(), controllers.GetFollowers)
	r.GET("/:id/following", middlewares.JWTAuthMiddleware(), controllers.GetFollowing)

}



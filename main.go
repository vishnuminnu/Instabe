package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/vishnu/instabe/routes"
	"github.com/vishnu/instabe/services"
	"github.com/vishnu/instabe/utils"
)

func main() {
    gotenv.Load()
	utils.ConnectDB()
	services.InitUserService()
	services.InitPostService()
	services.InitCommentService()
	services.InitStoryService()
	services.InitMessageService()
	fmt.Println("hello world")
	r:=gin.Default()
	routes.RegisterUserRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.StoryRoutes(r)
	routes.CommentRoutes(r)
    routes.MessageRoutes(r)
	log.Println("server running on port 8000")
	fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

	
	r.Run(":8000")
}





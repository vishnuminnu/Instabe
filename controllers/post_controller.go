package controllers

import (
	"net/http"

	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/services"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context){
	var post models.Post

	if err := c.ShouldBindJSON(&post);
	err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
if err :=  services.CreatePost(&post);
err!=nil{
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
	return
}
c.JSON(http.StatusCreated,gin.H{"message":"Post created successfully"})

}


func GetPosts(c *gin.Context) {
	posts, err := services.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
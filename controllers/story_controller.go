package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/services"
)

// controllers/story_controller.go
func CreateStory(c *gin.Context) {
	var story models.Story
	if err := c.ShouldBindJSON(&story); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	story.ID = uuid.New()
	story.CreatedAt = time.Now()
	story.ExpiresAt = story.CreatedAt.Add(24 * time.Hour)

	if err := services.CreateStory(story); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create story"})
		return
	}
	c.JSON(http.StatusCreated, story)
}

func GetUserStories(c *gin.Context) {
	userID := c.Param("user_id")
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	stories, err := services.GetStoriesByUser(uuidUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch stories"})
		return
	}
	c.JSON(http.StatusOK, stories)
}
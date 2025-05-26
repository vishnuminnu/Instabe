package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var storyCollection *mongo.Collection

func InitStoryService() {
	storyCollection = utils.DB.Collection("stories")
}

func CreateStory(story models.Story) error {
	_, err := storyCollection.InsertOne(context.TODO(), story)
	return err
}

func GetStoriesByUser(userID uuid.UUID) ([]models.Story, error) {
	ctx := context.TODO()
	var stories []models.Story
	cursor, err := storyCollection.Find(ctx, bson.M{"user_id": userID, "expires_at": bson.M{"$gt": time.Now()}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var s models.Story
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}
	return stories, nil
}

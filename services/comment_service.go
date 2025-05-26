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


var commentCollection *mongo.Collection 
func InitCommentService() {
	commentCollection = utils.DB.Collection("comments")
}

func CreateCommentService(comment models.Comment) error {
	comment.ID = uuid.New()
	comment.CreatedAt = time.Now()
	_, err := commentCollection.InsertOne(context.TODO(), comment)
	return err
}

func GetCommentsByPostIDService(postID uuid.UUID) ([]models.Comment, error) {
	var comments []models.Comment
	cursor, err := commentCollection.Find(context.TODO(), bson.M{"post_id": postID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var comment models.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

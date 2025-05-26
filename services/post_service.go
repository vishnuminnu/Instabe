package services

import (
	"context"
	"time"

	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/utils"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var postCollection *mongo.Collection 
func InitPostService() {
	postCollection = utils.DB.Collection("posts")
}

func CreatePost(post *models.Post)error{
      post.ID = uuid.New()
	  post.CreatedAt = time.Now()
	  _,err := postCollection.InsertOne(context.Background(),post)
	  return err
}

func GetAllPosts()([]models.Post , error){
	var posts []models.Post
	cursor,err := postCollection.Find(context.Background(),bson.M{})
	if err!=nil{
		return nil ,err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()){
		var post models.Post
		if err := cursor.Decode(&post);
		err!=nil{
			return nil,err
		}

		comments, err := GetCommentsByPostIDService(post.ID)
		if err != nil {
			return nil, err
		}
		post.Comments = comments

		posts = append(posts,post)
	}
	return posts,nil
}
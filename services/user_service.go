package services

import (
	"context"
	"time"
    "log"
	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/utils"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserService() {
	userCollection = utils.DB.Collection("users")
}


func CreateUser(user *models.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hash)
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("MongoDB Insert error:", err) // <-- add this
	}
	return err
}


func FindUserByEmail(email string) (*models.User,error){
	var user models.User
	err := userCollection.FindOne(context.Background(),bson.M{"email":email}).Decode(&user)
	return &user,err
}
















func GetAllUsers() ([]models.User, error) {
	var users []models.User
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}




func FollowUserService(targetUserID string, followerID string) error {
	targetID, _ := uuid.Parse(targetUserID)
	followerUUID, _ := uuid.Parse(followerID)
	userCollection := utils.DB.Collection("users")
	ctx := context.Background()

	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": targetID}, bson.M{"$addToSet": bson.M{"followers": followerUUID}})
	if err != nil {
		return err
	}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": followerUUID}, bson.M{"$addToSet": bson.M{"following": targetID}})
	return err
}

func UnfollowUserService(targetUserID string, followerID string) error {
	targetID, _ := uuid.Parse(targetUserID)
	followerUUID, _ := uuid.Parse(followerID)
	userCollection := utils.DB.Collection("users")
	ctx := context.Background()

	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": targetID}, bson.M{"$pull": bson.M{"followers": followerUUID}})
	if err != nil {
		return err
	}
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": followerUUID}, bson.M{"$pull": bson.M{"following": targetID}})
	return err
}

func GetFollowersService(userID string) ([]models.User, error) {
	id, _ := uuid.Parse(userID)
	userCollection := utils.DB.Collection("users")
	ctx := context.Background()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	var followers []models.User
	cursor, err := userCollection.Find(ctx, bson.M{"_id": bson.M{"$in": user.Followers}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &followers)
	return followers, err
}

func GetFollowingService(userID string) ([]models.User, error) {
	id, _ := uuid.Parse(userID)
	userCollection := utils.DB.Collection("users")
	ctx := context.Background()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	var following []models.User
	cursor, err := userCollection.Find(ctx, bson.M{"_id": bson.M{"$in": user.Following}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &following)
	return following, err
}

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
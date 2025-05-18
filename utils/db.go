package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var DB *mongo.Database

func ConnectDB(){
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err!=nil{
		log.Fatal(err)
	} 
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	} 
    fmt.Println("Connected to mongodb")
	DB = client.Database(os.Getenv("DB_NAME"))
}
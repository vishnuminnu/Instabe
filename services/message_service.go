package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vishnu/instabe/models"
	"github.com/vishnu/instabe/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageCollection *mongo.Collection

func InitMessageService() {
	messageCollection = utils.DB.Collection("messages")
}

func CreateMessage(msg models.Message) error {
	msg.ID = uuid.New()
	msg.SentAt = time.Now()
	_, err := messageCollection.InsertOne(context.TODO(), msg)
	return err
}

func GetMessagesBetweenUsers(senderID, receiverID uuid.UUID) ([]models.Message, error) {
	filter := map[string]interface{}{
		"$or": []map[string]interface{}{
			{"sender_id": senderID, "receiver_id": receiverID},
			{"sender_id": receiverID, "receiver_id": senderID},
		},
	}

	cur, err := messageCollection.Find(context.TODO(), filter, options.Find().SetSort(map[string]int{"sent_at": 1}))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var messages []models.Message
	for cur.Next(context.TODO()) {
		var msg models.Message
		if err := cur.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
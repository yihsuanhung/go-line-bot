package model

import (
	"context"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/internal/db"
	"github.com/yihsuanhung/go-line-bot/pkg/bot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MessageId string             `bson:"messageId" json:"messageId"`
	Message   string             `bson:"message" json:"message"`
	UserId    string             `bson:"userId" json:"userId"`
	UserName  string             `bson:"userName" json:"userName"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SendingMessage struct {
	UserId  string `json:"userId" bons:"userId"`
	Message string `json:"message" bons:"message"`
}

func GetAllMessages() ([]Message, error) {
	var msgs []Message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &msgs); err != nil {
		return nil, err
	}
	return msgs, nil
}

func GetMessageByID(id string) (Message, error) {
	var msg Message
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return msg, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&msg)
	return msg, err
}

func CreateMessage(m *Message) error {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection.InsertOne(ctx, m)
	return err
}

func UpdateMessage(msg Message) error {
	msg.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection.ReplaceOne(ctx, bson.M{"_id": msg.ID}, msg)
	return err
}

func DeleteMessage(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func SendMEssage(m *SendingMessage) error {
	_, err := bot.LineBot.PushMessage(m.UserId, linebot.NewTextMessage(m.Message)).Do()
	return err
}

func GetMessagesByUserId(userId string) ([]Message, error) {
	var msgs []Message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection.Find(ctx, bson.D{{Key: "userId", Value: userId}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &msgs); err != nil {
		return nil, err
	}
	return msgs, err
}

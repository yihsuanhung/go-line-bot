package model

import (
	"context"
	"time"

	"github.com/yihsuanhung/go-line-bot/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type User struct {
// 	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
// 	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
// 	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
// 	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
// 	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
// }

type UserMessage struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MessageId string             `bson:"messageId" json:"messageId"`
	Message   string             `bson:"message" json:"message"`
	UserId    string             `bson:"userId" json:"userId"`
	UserName  string             `bson:"userName" json:"userName"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// func GetAllUsers() ([]User, error) {
// 	var users []User
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	cursor, err := Collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = cursor.All(ctx, &users); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func GetUserByID(id string) (User, error) {
// 	var user User
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return user, err
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
// 	return user, err
// }

func CreateUserMessage(m *UserMessage) error {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection.InsertOne(ctx, m)
	return err
}

// func UpdateUser(user User) error {
// 	user.UpdatedAt = time.Now()
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	_, err := Collection.ReplaceOne(ctx, bson.M{"_id": user.ID}, user)
// 	return err
// }

// func DeleteUser(id string) error {
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	_, err = Collection.DeleteOne(ctx, bson.M{"_id": objID})
// 	return err
// }

// LEGACY

// import (
// 	"context"
// 	"fmt"

// 	"github.com/yihsuanhung/go-line-bot/internal/db"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// type User struct {
// 	Name string `bson:"name" json:"name"`
// 	Id   string `bson:"id" json:"id"`
// }

// type Message struct {
// 	Id      string `bson:"id" json:"id"`
// 	Message int    `bson:"message" json:"message"`
// 	User    User   `bson:"user" json:"user"`
// }

// func CreateMessage() {
// 	fmt.Println("新增信息")
// 	db.Collection.InsertOne(context.TODO(), bson.D{{Key: "message", Value: "asdf"}, {Key: "value", Value: 123}})
// }

// func GetMessage() {}

// func UpdateMessage() {}

// func DeleteMessage() {}

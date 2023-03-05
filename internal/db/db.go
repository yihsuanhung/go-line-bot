package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var Collection *mongo.Collection

// Deprecated
func Init() error {
	fmt.Println("connecting...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 連線
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			// return err
			panic(err)
		}
	}()

	// 連線檢查
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		// panic(err)
		return err
	}

	fmt.Println("connected!")

	// 資料庫
	Collection = client.Database("testing").Collection("line")
	// Collection.InsertOne(ctx, bson.D{{Key: "asdf", Value: "asdf"}, {Key: "value", Value: 11111}})
	fmt.Println("MongoDB connected")
	return nil
}

// get collection
// collection := client.Database("testing").Collection("numbers")

// insert data into collection
// collection.InsertOne(ctx, bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 11111}})
// collection.InsertMany(context.Background(), docs)

// collection.ReplaceOne(
// 	context.Background(),
// 	bson.D{ // 设置查询条件, item=paper
// 		{"item", "paper"},
// 	},
// 	bson.D{ // 设置新的文档内容
// 		{"item", "paper"},
// 		{"instock", bson.A{
// 			bson.D{
// 				{"warehouse", "A"},
// 				{"qty", 60},
// 			},
// 			bson.D{
// 				{"warehouse", "B"},
// 				{"qty", 40},
// 			},
// 		}},
// 	},
// )

// id := res.InsertedID
// fmt.Println("新增文档Id=", id)

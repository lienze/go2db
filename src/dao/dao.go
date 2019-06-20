package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB(dbname string) {
	fmt.Println("Hello MongoDB, Connecting...")
	var err error
	dbClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		fmt.Println("[ConnectDB]", err)
	}
	curDB = dbClient.Database(dbname)
	//curColl = curDB.Collection("info")
	//curColl.InsertOne(ctx, bson.M{"Name": "pi", "Age": 17})
}

func SetCurDB(dbname string) bool {
	if !IsConnected() {
		return false
	}
	curDB = dbClient.Database(dbname)
	return true
}

func SetCurColl(collname string) bool {
	if !IsConnected() {
		return false
	}
	if curDB == nil {
		return false
	}
	curColl = curDB.Collection(collname)
	return true
}

func IsConnected() bool {
	if dbClient == nil {
		return false
	} else {
		return true
	}
}

func GetCurDatabase() *mongo.Database {
	return curDB
}

func GetCurColl() *mongo.Collection {
	return curColl
}

func InsertData(data bson.M) bool {
	if curColl != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		curColl.InsertOne(ctx, data)
		return true
	}
	return false
}


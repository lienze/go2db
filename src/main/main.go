package main

import (
	"../dao"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Hello World")
	dao.ConnectDB("mytest")
	dao.SetCurColl("info")
	if !dao.IsConnected() {
		fmt.Println("Connect DB error")
	}
	//testData := bson.M{"Name":"bsonName", "Age":30}
	//dao.InsertData(testData)
	findData := bson.M{"Name":"John"}
	_, ret := dao.QueryData(findData)
	fmt.Println(len(ret))
}

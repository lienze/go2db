package main

import (
	"fmt"
	"github.com/lienze/go2db/dao"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Hello World")
	dao.InitDB("mytest")
	dao.SetCurColl("info")

	//crud Create
	//testData := bson.M{"Name":"bsonName", "Age":30}
	//dao.InsertData(testData)

	//crud Retrieve
	//findData := bson.M{"Name":"John"}
	//ret, retval := dao.QueryData(findData)
	//if ret == true {
	//	if (len(retval) > 0) {
	//		for _, val := range retval {
	//			// do something with each result
	//			fmt.Println(val)
	//		}
	//	}
	//}

	//crud Update
	//filterData := bson.D{{"Name",nil}}
	//newData := bson.D{{"$set",bson.D{{"Name","NULL"}}}}
	//dao.UpdateData(filterData, newData)

	//crud Delete
	filterData := bson.D{{"Name", "newJohn"}}
	dao.DeleteData(filterData)
}

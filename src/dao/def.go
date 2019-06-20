package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client = nil
var curDB *mongo.Database = nil
var curColl *mongo.Collection = nil

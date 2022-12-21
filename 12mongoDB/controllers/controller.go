package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://root:root@cluster0.t2kgtog.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const CollectionName = "WatchList"

var collection *mongo.Collection

func init() {
	//client connection

	clientOption := options.Client().ApplyURI(ConnectionString)

	//connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection sucess")

	collection = client.Database(dbName).Collection(CollectionName)
	fmt.Println("Collecton instance created")
}

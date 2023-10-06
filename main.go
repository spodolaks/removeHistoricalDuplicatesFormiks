package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DuplicateGroup struct {
	ID         bson.M   `bson:"_id"`
	Duplicates []string `bson:"duplicates"`
	Count      int      `bson:"count"`
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://admin:welcome%4031415@20.52.39.84:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Choose the database and collection

	collection := client.Database("formiks_v2").Collection("submissions")

	

	// Define a filter to find documents where parentId is an ObjectId and data.projectNumber equals "ddd"
	filter := bson.D{
		{"parentId", bson.D{{"$type", "objectId"}}},
		{"data.projectNumber", "6110CH220211"},
	}

	// Delete the documents matching the filter
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents\n", deleteResult.DeletedCount)
}


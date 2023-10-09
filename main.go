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

	

	filter := bson.D{
		{"parentId", bson.M{"$eq": nil}},
		{"$or", bson.A{
			bson.M{"data.projectType": "Project"},
			bson.M{"data.projectType": ""},
			bson.M{"data.projectType": bson.M{"$exists": false}},
		}},
	}

	// Define an update to set data.projectType to "Historical Project" for all matched documents
	update := bson.D{
		{"$set", bson.D{
			{"data.projectType", "Historical Project"},
		}},
	}

	// Update all documents matching the filter
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v root documents and modified root %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	filter = bson.D{
		{"$and", bson.A{
			bson.M{"parentId": bson.M{"$ne": nil}},
			bson.M{"data.projectType": "Historical project"},
		}},
	}

	// Define an update to set data.projectType to "Historical Project" for all matched documents
	update = bson.D{
		{"$set", bson.D{
			{"data.projectType", "Historical Project"},
		}},
	}

	// Update all documents matching the filter
	updateResult, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v child documents and modified child %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)


}


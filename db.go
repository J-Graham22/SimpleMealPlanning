package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func setupMongoDBConnection() *mongo.Database {
	uri := os.Getenv("MONGODB_URI") //Needs to be set up as an environment variable

	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not set.")
	}

	dbClient, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	database := dbClient.Database("mealplanning")

	return database
}

func getAllProteins(db *mongo.Database) {
	coll := db.Collection("proteins")

	filter := bson.D{}
	sort := bson.D{}
	opts := options.Find().SetSort(sort)

	proteins, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	//TODO: decode proteins

	return proteins
}

func getMealParts() Proteins {
	content, err := os.ReadFile("./proteins.json")
	if err != nil {
		log.Fatal("Error opening proteins json")
	}

	var proteins Proteins
	err = json.Unmarshal(content, &proteins)
	if err != nil {
		log.Fatal("Error parsing content")
	}

	return proteins
}

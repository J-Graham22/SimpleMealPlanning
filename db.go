package main

import (
	"encoding/json"
	"log"
	"os"

	_ "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func setupMongoDBConnection() *mongo.Client {
	uri := os.Getenv("MONGODB_URI") //Needs to be set up as an environment variable

	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not set.")
	}

	dbClient, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return dbClient
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

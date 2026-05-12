package main

import (
	"os"
	"log"
	"encoding/json"
)

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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type MealPart struct {
	Name   string `json: "name"`
	Recipe string `json: "recipe"`
}

type Proteins struct {
	Proteins []MealPart `json: "proteins"`
}

type Meal struct {
	protein   MealPart
	side      MealPart
	vegetable MealPart
}

func main() {
	fmt.Println("Hello World!")

	content, err := os.ReadFile("./proteins.json")
	if err != nil {
		log.Fatal("Error opening proteins json")
	}

	var proteins Proteins
	err = json.Unmarshal(content, &proteins)
	if err != nil {
		log.Fatal("Error parsing content")
	}

	for _, p := range proteins.Proteins {
		log.Printf("name: %s\n", p.Name)
		log.Printf("recipe: %s\n", p.Recipe)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world\n")
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/a-h/templ"
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

	r := http.NewServeMux()

	component := hello("Justin")
	r.Handle("/", templ.Handler(component))

	server := http.Server{
		Addr: ":3000",
		Handler: r,
	}
	fmt.Println("Listening on :3000")
	server.ListenAndServe()
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	component := hello("justin")
	component.Render(context.Background(), os.Stdout)
	fmt.Fprintf(w, "hello, world\n")
}

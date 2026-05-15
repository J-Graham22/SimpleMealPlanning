package main

import (
	_ "context"
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
	_ "os"
)

func main() {
	fmt.Println("Hello World!")

	proteins := getMealParts()

	for _, p := range proteins.Proteins {
		log.Printf("name: %s\n", p.Name)
		log.Printf("recipe: %s\n", p.Recipe)
	}

	r := http.NewServeMux()

	component := plan()
	r.Handle("/", templ.Handler(component))

	server := http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	fmt.Println("Listening on :3000")
	server.ListenAndServe()
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world\n")
}

func generateWeeklyMealPlan(w http.ResponseWriter, r *http.Request) {

}

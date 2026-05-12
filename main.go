package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/a-h/templ"
)


func main() {
	fmt.Println("Hello World!")

	proteins := getMealParts()

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

func generateWeeklyMealPlan(w http.ResponseWriter, r *http.Request) {

}

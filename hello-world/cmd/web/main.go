package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/et-codes/lab/hello-world/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

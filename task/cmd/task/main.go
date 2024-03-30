package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Task Manager")

	path := "./tasks.db"

	db, err := OpenDB(path)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Database %q opened.\n", path)
}

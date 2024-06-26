package main

import (
	"datnguyen/todo/database"
	"fmt"
	"log"
)

func main() {
	// Initialize the database connection
	database.DatabaseInit()

	println("Database Initialized!")

	todos, err := database.GetAllTodos()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Todos found: %v\n", todos)
	
	todo,err := database.GetTodoById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Todo with id %d : %v", 1, todo)
}

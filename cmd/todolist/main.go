package main

import (
	"log"
	"net/http"

	"github.com/tammarut/todo-list/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/tasks", handler.TaskHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/tammarut/todo-list/internal/handler"
)

type InMemoryTaskStore struct {
}

func (i *InMemoryTaskStore) GetTask(task string) string {
	return "cook"
}

func main() {
	handler := &handler.TaskHandler{&InMemoryTaskStore{}}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/tammarut/todo-list/internal/handler"
)

// InMemoryTaskStore is struct
type InMemoryTaskStore struct {
}

// GetTask will return task
func (i *InMemoryTaskStore) GetTask(task string) string {
	return "cook"
}

// RecordTask will save task in memory
func (i *InMemoryTaskStore) RecordTask(name string) {
}
func main() {
	handler := &handler.TaskHandler{&InMemoryTaskStore{}}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}

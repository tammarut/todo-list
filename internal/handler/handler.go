package handler

import (
	"fmt"
	"net/http"
	"strings"
)

type TaskStore interface {
	GetTask(task string) string
}

type TaskHandler struct {
	Store TaskStore
}

func (t *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	task := strings.TrimPrefix(r.URL.Path, "/api/v1/tasks")
	fmt.Fprint(w, t.Store.GetTask(task))
}

func GetTask(task string) string {
	if task == "" {
		return "sleep"
	}
	return ""
}

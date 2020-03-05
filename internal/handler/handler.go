package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// TaskStore is interface
type TaskStore interface {
	GetTask(task string) string
	RecordTask(name string)
}

// TaskHandler is struct implementing TaskStore interface
type TaskHandler struct {
	Store TaskStore
}

// ServeHTTP is handler from request and switch method
func (t *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.processStore(w)
	case http.MethodGet:
		t.showTask(w, r)
	}

}

func (t *TaskHandler) processStore(w http.ResponseWriter) {
	t.Store.RecordTask("homework")
	w.WriteHeader(http.StatusAccepted)
}

func (t *TaskHandler) showTask(w http.ResponseWriter, r *http.Request) {
	task := strings.TrimPrefix(r.URL.Path, "/api/v1/tasks")

	title := t.Store.GetTask(task)
	if title == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, title)
}

// GetTask will return task
func GetTask(task string) string {
	if task == "" {
		return "sleep"
	}
	return ""
}

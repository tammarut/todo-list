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
	switch r.Method {
	case http.MethodPost:
		t.processStore(w)
	case http.MethodGet:
		t.showTask(w, r)
	}

}

func (t *TaskHandler) processStore(w http.ResponseWriter) {
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

func GetTask(task string) string {
	if task == "" {
		return "sleep"
	}
	return ""
}

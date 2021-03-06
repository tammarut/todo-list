package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubTaskStore struct {
	title      map[string]string
	titleCalls []string
}

func (s *StubTaskStore) GetTask(taskname string) string {
	task := s.title[taskname]
	return task
}
func (s *StubTaskStore) RecordTask(name string) {
	s.titleCalls = append(s.titleCalls, name)
}

func TestGETTasks(t *testing.T) {
	store := StubTaskStore{
		map[string]string{
			"": "sleep",
		},
		nil,
	}
	handler := &TaskHandler{&store}

	t.Run("returns all tasks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := response.Body.String()
		want := "sleep"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns 404 No tasks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/work", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreTasks(t *testing.T) {
	store := StubTaskStore{
		map[string]string{},
		nil,
	}
	handler := &TaskHandler{&store}

	t.Run("Accepted store", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}

		if len(store.titleCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.titleCalls), 1)
		}
	})
}

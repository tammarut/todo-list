package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubTaskStore struct {
	title map[string]string
}

func (s *StubTaskStore) GetTask(taskname string) string {
	task := s.title[taskname]
	return task
}

func TestGETTasks(t *testing.T) {
	store := StubTaskStore{
		map[string]string{
			"": "sleep",
		},
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
}

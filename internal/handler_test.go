package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTasks(t *testing.T) {
	t.Run("returns all tasks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
		response := httptest.NewRecorder()

		TaskHandler(response, request)

		got := response.Body.String()
		want := "sleep"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

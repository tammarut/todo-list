package handler

import (
	"fmt"
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "sleep")
}

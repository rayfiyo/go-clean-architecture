package httpserver

import (
	"net/http"

	httpi "github.com/you/todoapp/internal/interface/http"
)

func NewMux(h *httpi.TodoHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", h.PostTodo)
	return mux
}

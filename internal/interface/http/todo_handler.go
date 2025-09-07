package httpi

import (
	"net/http"

	"github.com/you/todoapp/internal/usecase"
)

type TodoHandler struct {
	Add *usecase.AddTodoUsecase
}

func (h *TodoHandler) PostTodo(w http.ResponseWriter, r *http.Request) {
	// ここは薄く：HTTPの入出力をユースケースのI/Oに詰め替えるだけ
}

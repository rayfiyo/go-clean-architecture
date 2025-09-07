package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/you/todoapp/internal/domain"
	"github.com/you/todoapp/internal/infrastructure/sqlite"
	httpi "github.com/you/todoapp/internal/interface/http"
	"github.com/you/todoapp/internal/presentation/httpserver"
	"github.com/you/todoapp/internal/usecase"
)

func main() {
	db, _ := sql.Open("sqlite", "file:todo.db")
	var repo domain.TodoRepository = &sqlite.TodoRepositorySQLite{DB: db}

	addUC := &usecase.AddTodoUsecase{Repo: repo}
	handler := &httpi.TodoHandler{Add: addUC}
	srv := &http.Server{Addr: ":8080", Handler: httpserver.NewMux(handler)}
	log.Fatal(srv.ListenAndServe())
}

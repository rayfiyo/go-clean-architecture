package sqlite

import (
	"database/sql"

	"github.com/you/todoapp/internal/domain"
)

type TodoRepositorySQLite struct{ DB *sql.DB }

func (r *TodoRepositorySQLite) Save(t domain.Todo) error        { /* ... */ return nil }
func (r *TodoRepositorySQLite) FindAll() ([]domain.Todo, error) { /* ... */ return nil }

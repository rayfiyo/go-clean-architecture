package domain

type TodoRepository interface {
	Save(t Todo) error
	FindAll() ([]Todo, error)
}

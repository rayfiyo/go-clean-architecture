package usecase

import "github.com/you/todoapp/internal/domain"

type (
	AddTodoInput  struct{ Title string }
	AddTodoOutput struct{ ID domain.TodoID }
)

type AddTodoUsecase struct {
	Repo domain.TodoRepository
}

func (u *AddTodoUsecase) Execute(in AddTodoInput) (AddTodoOutput, error) {
	t := domain.Todo{ID: domain.TodoID(in.Title), Title: in.Title}
	if err := u.Repo.Save(t); err != nil {
		return AddTodoOutput{}, err
	}
	return AddTodoOutput{ID: t.ID}, nil
}

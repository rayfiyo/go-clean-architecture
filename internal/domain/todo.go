package domain

import "time"

type TodoID string

type Todo struct {
	ID        TodoID
	Title     string
	Done      bool
	CreatedAt time.Time
}

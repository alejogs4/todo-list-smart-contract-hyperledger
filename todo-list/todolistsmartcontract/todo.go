package todolistsmartcontract

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTodoInformation = errors.New("Invalid title or owner information")
	ErrGettingTodo            = errors.New("Error Getting todo")
	ErrNotFoundTodo           = errors.New("Todo was not found")
	ErrInvalidOwner           = errors.New("Introduced owner is not the owner for the task")
)

type Timestamp interface {
	GetSeconds() int64
}

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
}

func (t *Todo) Clone() Todo {
	return Todo{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		Owner:     t.Owner,
	}
}

func CreateTodo(title, owner string, timestamp Timestamp) Todo {
	return Todo{
		ID:        fmt.Sprint(timestamp.GetSeconds()),
		Title:     title,
		Completed: false,
		Owner:     owner,
	}
}

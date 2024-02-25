package repository

import (
	"errors"
	"go-practice/domain/object"
	"go-practice/domain/repository"
)

type todoRepository struct{}

func NewTodoReposiory() repository.ITodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) GetAll() ([]*object.Todo, error) {
	return todos, nil
}

func (r *todoRepository) Get(id int) (*object.Todo, error) {
	for _, todo := range todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return nil, errors.New("")
}

func (r *todoRepository) Store(t *object.Todo) error {
	todos = append(todos, t)
	return nil
}

func (r *todoRepository) Update(t *object.Todo) error {
	for _, todo := range todos {
		if todo.Id == t.Id {
			todo.Title = t.Title
			return nil
		}
	}
	return errors.New("")
}

func (r *todoRepository) Delete(id int) error {
	tmp := []*object.Todo{}
	for _, todo := range todos {
		if todo.Id != id {
			tmp = append(tmp, todo)
		}
	}
	todos = tmp
	return nil
}

var todos = []*object.Todo{
	{
		Id:    1,
		Title: "study",
	},
	{
		Id:    2,
		Title: "training",
	},
	{
		Id:    3,
		Title: "work",
	},
}

package repository

import (
	"go-practice/domain/object"
)

type ITodoRepository interface {
	GetAll() ([]*object.Todo, error)
	Get(int) (*object.Todo, error)
	Store(*object.Todo) error
	Update(*object.Todo) error
	Delete(int) error
}

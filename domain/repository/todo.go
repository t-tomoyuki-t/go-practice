package repository

import (
	"go-practice/domain/entity"
)

type ITodoRepository interface {
	GetAll() (*[]entity.Todo, error)
	Get(int) (*entity.Todo, error)
	Store(*entity.Todo) error
	Update(*entity.Todo) error
	Delete(int) error
}

package repository

import (
	"go-practice/domain/object"
	"go-practice/domain/repository"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoReposiory(db *gorm.DB) repository.ITodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAll() (*[]object.Todo, error) {
	todos := []object.Todo{}
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}
	return &todos, nil
}

func (r *todoRepository) Get(id int) (*object.Todo, error) {
	t := object.Todo{}
	res := r.db.First(&t, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &t, nil
}

func (r *todoRepository) Store(t *object.Todo) error {
	res := r.db.Create(t)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *todoRepository) Update(t *object.Todo) error {
	res := r.db.Save(t)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *todoRepository) Delete(id int) error {
	res := r.db.Delete(&object.Todo{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

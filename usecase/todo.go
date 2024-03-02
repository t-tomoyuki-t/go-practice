package usecase

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"
)

type TodoUseCase struct {
	r repository.ITodoRepository
}

func NewTodoUseCase(r repository.ITodoRepository) *TodoUseCase {
	return &TodoUseCase{r}
}

func (u *TodoUseCase) GetAll() (*[]entity.Todo, error) {
	l, err := u.r.GetAll()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (u *TodoUseCase) Get(id int) (*entity.Todo, error) {
	t, err := u.r.Get(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (u *TodoUseCase) Store(t *entity.Todo) error {
	err := u.r.Store(t)
	if err != nil {
		return err
	}
	return nil
}

func (u *TodoUseCase) Update(t *entity.Todo) error {
	err := u.r.Update(t)
	if err != nil {
		return err
	}
	return nil
}

func (u *TodoUseCase) Delete(id int) error {
	err := u.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"
)

type UserUseCase struct {
	r repository.IUserRepository
}

func NewUserUseCase(r repository.IUserRepository) *UserUseCase {
	return &UserUseCase{r}
}

func (u *UserUseCase) Get(id int) (*entity.User, error) {
	user, err := u.r.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

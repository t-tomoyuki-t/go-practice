package repository

import "go-practice/domain/entity"

type IUserRepository interface {
	Get(int) (*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	Save(*entity.User) error
}

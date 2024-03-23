package repository

import (
	"go-practice/domain/entity"
	"go-practice/domain/vo"
)

type IAuthRepository interface {
	GetSession(string) (string, error)
	SetSession(vo.Token, *entity.User) error
	DeleteSession(string) error
}

package repository

import (
	"go-practice/domain/entity"
	"go-practice/domain/vo"
)

type IAuthRepository interface {
	GetSession() error
	SetSession(vo.Token, *entity.User) error
	DeleteSession(vo.Token) error
}

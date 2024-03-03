package repository

import (
	"github.com/redis/go-redis/v9"
	"go-practice/domain/repository"
)

type authRepository struct {
	rdb *redis.Client
}

func NewAuthRepository(rdb *redis.Client) repository.IAuthRepository {
	return &authRepository{rdb}
}

func (r *authRepository) GetSession() error {
	return nil
}

func (r *authRepository) SetSession() error {
	return nil
}

func (r *authRepository) DeleteSession() error {
	return nil
}

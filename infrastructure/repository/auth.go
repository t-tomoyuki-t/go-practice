package repository

import (
	"context"
	"go-practice/domain/entity"
	"go-practice/domain/repository"
	"go-practice/domain/vo"
	"math"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
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

func (r *authRepository) SetSession(token vo.Token, user *entity.User) error {
	ctx := context.Background()
	err := r.rdb.Set(ctx, token.String(), strconv.Itoa(user.Id), time.Duration(vo.TTL_SECOND * int(math.Pow10(9)))).Err()
	return err
}

func (r *authRepository) DeleteSession(token vo.Token) error {
	ctx := context.Background()
	err := r.rdb.Del(ctx, token.String()).Err()
	return err
}

package usecase

import (
	"go-practice/domain/entity"
	"go-practice/domain/repository"
	"go-practice/domain/vo"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	ar repository.IAuthRepository
	ur repository.IUserRepository
}

func NewAuthUseCase(ar repository.IAuthRepository, ur repository.IUserRepository) *AuthUseCase {
	return &AuthUseCase{
		ar,
		ur,
	}
}

func (au *AuthUseCase) Login(a entity.Auth) (vo.Token, error) {
	u, err := au.ur.GetByEmail(a.Email)
	if err != nil {
		return vo.Token{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.GetPassword()), []byte(a.Password))
	if err != nil {
		return vo.Token{}, err
	}

	token, err := vo.NewToken()
	if err != nil {
		return token, err
	}

	err = au.ar.SetSession(token, u)
	if err != nil {
		return vo.Token{}, err
	}

	return token, nil
}
